package updater

import (
	"fmt"
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func updateItemSubclass(data *blizzard_api.ApiResponse) {
	var itemSubclass datasets.ItemSubclass
	data.Parse(&itemSubclass)
	insertOnceExpr(&itemSubclass, "(id,class_id) DO UPDATE", "display_name", "hide_subclass_in_tooltips")
}

func UpdateItemClass(data *blizzard_api.ApiResponse) {
	var itemClass datasets.ItemClass
	data.Parse(&itemClass)
	insertOnceUpdate(&itemClass, "name")

	for _, subclass := range itemClass.ItemSubclasses {
		subclassResponse := connections.WowClient.ItemSubclass(itemClass.ID, subclass.ID, nil)
		updateItemSubclass(subclassResponse)
	}
}

func updateItemProfessionRequirement(item *datasets.Item) {
	p_id := 0
	p_t_id := 0

	profession := &datasets.Profession{
		Identifiable: datasets.Identifiable{
			ID: item.PreviewItem.Requirements.Skill.Profession.GetID(),
		},
	}
	err := connections.GetDBConn().Model(profession).WherePK().Select()
	if err != nil {
		connections.GetDBConn().Model(&datasets.UpdateError{
			Endpoint: "Item profession requirement",
			RecordID: item.ID,
			Error:    fmt.Sprintf("Item %d: references unknown profession %d.", item.ID, profession.ID),
		}).Insert()
	} else {
		p_id = profession.ID
	}

	if p_id == 0 {
		professionTier := &datasets.ProfessionTier{
			Identifiable: datasets.Identifiable{
				ID: item.PreviewItem.Requirements.Skill.Profession.GetID(),
			},
		}
		err2 := connections.GetDBConn().Model(professionTier).WherePK().Select()
		if err2 != nil {
			connections.GetDBConn().Model(&datasets.UpdateError{
				Endpoint: "Item profession requirement",
				RecordID: item.ID,
				Error:    fmt.Sprintf("Item %d: references skill-tier %d not found.", item.ID, professionTier.ID),
			}).Insert()
		} else {
			p_id = professionTier.ProfessionID
			p_t_id = professionTier.ID
		}
	}

	if profession.ID > 0 {
		skill := datasets.ItemSkillRequirement{
			ItemID:           item.ID,
			ProfessionID:     p_id,
			ProfessionTierID: p_t_id,
			DisplayString:    item.PreviewItem.Requirements.Skill.DisplayString,
			Level:            item.PreviewItem.Requirements.Skill.Level,
		}
		insertOnceExpr(&skill, "(item_id) DO UPDATE", "profession_id", "profession_tier_id", "display_string", "level")
	} else {
		connections.GetDBConn().Model(&datasets.UpdateError{
			Endpoint: "Item profession requirement",
			RecordID: item.ID,
			Error:    fmt.Sprintf("Item %d: references invalid profession or skill-tier %d.", item.ID, item.PreviewItem.Requirements.Skill.Profession.GetID()),
		}).Insert()
	}
}

func UpdateItem(data *blizzard_api.ApiResponse) {
	var item datasets.Item
	data.Parse(&item)

	if item.InventoryType != nil {
		item.InventoryTypeID = item.InventoryType.ID
		insertOnce(item.InventoryType)
	}

	if item.Quality != nil {
		item.QualityID = item.Quality.ID
		insertOnce(item.Quality)
	}

	if item.Class.ID == 10 {
		insertOnce(&datasets.ItemClass{
			ID:             10,
			Name:           item.Class.Name,
		})
		insertOnceExpr(&datasets.ItemSubclass{
			ID:          item.Subclass.ID,
			ClassID:     10,
			DisplayName: item.Subclass.Name,
		}, "(id,class_id) DO NOTHING")
	}
	item.ItemClassID = item.Class.ID
	item.ItemSubclassClassID = item.Class.ID
	item.ItemSubclassID = item.Subclass.ID

	if item.PreviewItem.Binding != nil {
		insertOnce(item.PreviewItem.Binding)
		item.BindingID = item.PreviewItem.Binding.ID
	}

	if item.PreviewItem.Description != nil {
		item.Description = item.PreviewItem.Description
	}

	if item.PreviewItem.UniqueEquipped != nil {
		item.UniqueEquipped = item.PreviewItem.UniqueEquipped
	}

	if item.PreviewItem.NameDescription != nil {
		item.NameDescription = item.PreviewItem.NameDescription
	}

	if item.PreviewItem.Armor != nil {
		item.Armor = item.PreviewItem.Armor.Value
		item.ArmorDisplayString = &item.PreviewItem.Armor.Display
	}

	if item.PreviewItem.Durability != nil {
		item.Durability = item.PreviewItem.Durability.Value
		item.DurabilityDisplayString = &item.PreviewItem.Durability.DisplayString
	}

	if item.PreviewItem.Level != nil {
		item.Level = item.PreviewItem.Level.Value
		item.LevelDisplayString = &item.PreviewItem.Level.DisplayString
	}

	if item.PreviewItem.Charges != nil {
		item.Charges = item.PreviewItem.Charges.Value
		item.ChargesDisplayString = &item.PreviewItem.Charges.DisplayString
	}

	if item.PreviewItem.Context != nil {
		item.Context = *item.PreviewItem.Context
	}

	insertOnceUpdate(&item, "name", "description", "quality_id", "level", "required_level", "item_class_id", "item_subclass_id", "inventory_type_id", "purchase_price", "sell_price", "max_count", "is_equippable", "is_stackable", "preview_item")

	if item.Media != nil {
		item.Media.ItemID = item.ID
		insertOnceExpr(item.Media, "(id,item_id) DO NOTHING")
	}

	if item.PreviewItem.Requirements.Faction != nil {
		insertOnceExpr(&datasets.ItemFactionRequirement{
			ItemID:        item.ID,
			FactionID:     item.PreviewItem.Requirements.Faction.Faction.ID,
			DisplayString: item.PreviewItem.Requirements.Faction.DisplayString,
		},
			"(item_id) DO UPDATE",
			"display_string")
	}

	if item.PreviewItem.Requirements.PlayableRaces != nil {
		for _, race := range item.PreviewItem.Requirements.PlayableRaces.Links {
			insertOnceExpr(&datasets.ItemRaceRequirement{
				ItemID:         item.ID,
				PlayableRaceID: race.ID,
				DisplayString:  item.PreviewItem.Requirements.PlayableRaces.DisplayString,
			},
				"(item_id,playable_race_id) DO UPDATE",
				"display_string")
		}
	}

	if item.PreviewItem.Requirements.PlayableClasses != nil {
		for _, class := range item.PreviewItem.Requirements.PlayableClasses.Links {
			insertOnceExpr(&datasets.ItemClassRequirement{
				ItemID:          item.ID,
				PlayableClassID: class.ID,
				DisplayString:   item.PreviewItem.Requirements.PlayableClasses.DisplayString,
			},
				"(item_id,playable_class_id) DO UPDATE",
				"display_string")
		}
	}

	if item.PreviewItem.Requirements.PlayableSpecializations != nil {
		for _, race := range item.PreviewItem.Requirements.PlayableSpecializations.Links {
			insertOnceExpr(&datasets.ItemSpecializationRequirement{
				ItemID:                   item.ID,
				PlayableSpecializationID: race.ID,
				DisplayString:            item.PreviewItem.Requirements.PlayableSpecializations.DisplayString,
			},
				"(item_id,playable_specialization_id) DO UPDATE",
				"display_string")
		}
	}

	if item.PreviewItem.Requirements.Skill != nil {
		updateItemProfessionRequirement(&item)
	}

	if item.PreviewItem.Requirements.Ability != nil {
		insertOnceExpr(&datasets.ItemAbilityRequirement{
			ItemID:        item.ID,
			SpellID:       item.PreviewItem.Requirements.Ability.Spell.ID,
			DisplayString: item.PreviewItem.Requirements.Ability.DisplayString,
		},
		"(item_id) DO UPDATE",
		"spell_id", "display_string")
	}

	if item.PreviewItem.Requirements.Level != nil {
		insertOnceExpr(&datasets.ItemLevelRequirement{
			ItemID:        item.ID,
			Level:         item.PreviewItem.Requirements.Level.Level,
			DisplayString: item.PreviewItem.Requirements.Level.DisplayString,
		},
			"(item_id) DO UPDATE",
			"level", "display_string")
	}

	if item.PreviewItem.Requirements.Reputation != nil {
		insertOnceExpr(&datasets.ItemReputationRequirement{
			ItemID:        item.ID,
			FactionID:     item.PreviewItem.Requirements.Reputation.Faction.ID,
			DisplayString: item.PreviewItem.Requirements.Reputation.DisplayString,
		},
			"(item_id) DO UPDATE",
			"faction_id", "display_string")
	}

	if item.PreviewItem.Stats != nil {
		for _, stat := range item.PreviewItem.Stats {
			if stat.Stat == nil {
				db := connections.GetDBConn()
				db.Model(&datasets.UpdateError{
					Endpoint: "Item Stat",
					RecordID: item.ID,
					Error:    fmt.Sprintf("Stat type '%s' not found.", stat.Display.DisplayString.EnUS),
				}).Insert()
				continue
			}
			insertOnce(stat.Stat)

			stat.ItemID = item.ID
			stat.StatID = stat.Stat.ID
			insertOnceExpr(stat, "(item_id,stat_id) DO UPDATE", "value", "display", "is_negated", "is_equip_bonus")
		}
	}

	metadata := &datasets.ItemMetadata{
		ItemID:      item.ID,
		IsToy:       item.PreviewItem.Toy != nil,
		IsRecipe:    item.PreviewItem.Recipe != nil,
		IsReagent:   item.PreviewItem.CraftingReagent != nil,
		HasCharges:  item.PreviewItem.Charges != nil,
		HasDuration: item.PreviewItem.ExpirationTimeLeft != nil,
		HasSpells:   item.PreviewItem.Spells != nil,
	}
	insertOnceExpr(metadata,  "(item_id) DO UPDATE", "is_toy", "is_recipe", "is_reagent", "has_charges", "has_duration", "has_spells")
}

func UpdateItemMedia(data *blizzard_api.ApiResponse, id int) {
	var itemMedia datasets.ItemMedia
	data.Parse(&itemMedia)

	for _, asset := range itemMedia.Assets {
		asset.ItemMediaID = itemMedia.ID
		insertOnceExpr(&asset, "(item_media_id,key) DO UPDATE", "value")
	}
}
