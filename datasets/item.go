package datasets

type ItemClass struct {
	ID             int            `json:"class_id" pg:",pk,use_zero"`
	Name           LocalizedField `json:"name"`
	ItemSubclasses Identifiables  `json:"item_subclasses" pg:"-"`
}

type ItemSubclass struct {
	ID                     int            `json:"subclass_id" pg:",pk,use_zero"`
	ClassID                int            `json:"class_id" pg:",pk,use_zero"`
	Class                  *ItemClass     `pg:"rel:has-one"`
	DisplayName            LocalizedField `json:"display_name"`
	HideSubclassInTooltips bool           `json:"hide_subclass_in_tooltips"`
}

type ItemQuality Enum
type InventoryType Enum
type Binding Enum
type Stat Enum

type Item struct {
	Identifiable
	Name                    LocalizedField `json:"name"`
	NameDescription         LocalizedDisplayString
	Description             LocalizedField `json:"description"`
	QualityID               string         ``
	Quality                 *ItemQuality   `json:"quality" pg:"rel:has-one"`
	BindingID               string         ``
	Binding                 *Binding       `pg:"rel:has-one"`
	Level                   int            `json:"level"`
	LevelDisplayString      LocalizedField
	Armor                   int
	ArmorDisplayString      LocalizedDisplayString
	Durability              int
	DurabilityDisplayString LocalizedField
	Charges                 int
	ChargesDisplayString    LocalizedField
	UniqueEquipped          LocalizedField
	Context                 int
	RequiredLevel           int            `json:"required_level"`
	ItemClassID             int            ``
	ItemClass               *ItemClass     `pg:"rel:has-one"`
	Class                   Identifiable   `json:"item_class" pg:"-"`
	ItemSubclassID          int            ``
	ItemSubclassClassID     int            ``
	ItemSubclass            *ItemSubclass  `pg:"rel:has-one"`
	Subclass                Identifiable   `json:"item_subclass" pg:"-"`
	InventoryTypeID         string         ``
	InventoryType           *InventoryType `json:"inventory_type" pg:"rel:has-one"`
	PurchasePrice           int            `json:"purchase_price"`
	SellPrice               int            `json:"sell_price"`
	MaxCount                int            `json:"max_count"`
	IsEquippable            bool           `json:"is_equippable"`
	IsStackable             bool           `json:"is_stackable"`
	PreviewItem             *PreviewItem   `json:"preview_item"`
	Media                   *ItemMedia     `json:"media" pg:"-"`
}

type ItemMetadata struct {
	ItemID      int   `pg:",pk"`
	Item        *Item `pg:"rel:has-one"`
	IsToy       bool
	IsRecipe    bool
	IsReagent   bool
	HasCharges  bool
	HasDuration bool
	HasSpells   bool
}

type ItemSpell struct {
	ItemID       int            `pg:",pk"`
	Item         *Item          `pg:"rel:has-one"`
	SpellID      int            `pg:",pk"`
	Spell        *Spell         `pg:"rel:has-one"`
	Description  LocalizedField `json:"description"`
	DisplayColor Color          `json:"display_color"`
}

type ItemLevelRequirement struct {
	ItemID        int            `pg:",pk"`
	Item          *Item          `pg:"rel:has-one"`
	DisplayString LocalizedField `json:"display_string"`
	Level         int            `json:"value"`
}

type ItemRaceRequirement struct {
	ItemID         int            `pg:",pk"`
	Item           *Item          `pg:"rel:has-one"`
	DisplayString  LocalizedField `json:"display_string"`
	PlayableRaceID int            ``
	PlayableRace   *PlayableRace  `pg:"rel:has-one"`
}

type ItemClassRequirement struct {
	ItemID          int            `pg:",pk"`
	Item            *Item          `pg:"rel:has-one"`
	DisplayString   LocalizedField `json:"display_string"`
	PlayableClassID int            ``
	PlayableClass   *PlayableClass `pg:"rel:has-one"`
}

type ItemSpecializationRequirement struct {
	ItemID                   int                     `pg:",pk"`
	Item                     *Item                   `pg:"rel:has-one"`
	DisplayString            LocalizedField          `json:"display_string"`
	PlayableSpecializationID int                     `pg:",pk"`
	PlayableSpecialization   *PlayableSpecialization `pg:"rel:has-one"`
}

type ItemFactionRequirement struct {
	ItemID        int            `pg:",pk"`
	Item          *Item          `pg:"rel:has-one"`
	DisplayString LocalizedField `json:"display_string"`
	FactionID     string         ``
	Faction       *Faction       `json:"value" pg:"rel:has-one"`
}

type ItemReputationRequirement struct {
	ItemID             int                `pg:",pk"`
	Item               *Item              `pg:"rel:has-one"`
	DisplayString      LocalizedField     `json:"display_string"`
	FactionID          int                ``
	Faction            *ReputationFaction `json:"faction" pg:"rel:has-one"`
	MinReputationLevel int                `json:"min_reputation_level"`
}

type ItemSkillRequirement struct {
	ItemID           int             `pg:",pk"`
	Item             *Item           `pg:"rel:has-one"`
	ProfessionID     int             ``
	Profession       *Profession     `json:"profession" pg:"rel:has-one"`
	ProfessionTierID int             ``
	ProfessionTier   *ProfessionTier `pg:"rel:has-one"`
	DisplayString    LocalizedField  `json:"display_string"`
	Level            int             `json:"level"`
}

type ItemAbilityRequirement struct {
	ItemID        int            `pg:",pk"`
	Item          *Item          `pg:"rel:has-one"`
	DisplayString LocalizedField `json:"display_string"`
	SpellID       int            ``
	Spell         *Spell         `json:"spell" pg:"rel:has-one"`
}

type ItemStat struct {
	ItemID       int                    `pg:",pk"`
	Item         *Item                  `pg:"rel:has-one"`
	StatID       string                 `pg:",pk"`
	Stat         *Stat                  `json:"type" pg:"rel:has-one"`
	Value        int                    `json:"value"`
	Display      LocalizedDisplayString `json:"display"`
	IsNegated    bool                   `json:"is_negated"`
	IsEquipBonus bool                   `json:"is_equip_bonus"`
}

type PreviewItem struct {
	Spells    []*ItemSpell `json:"spells,omitempty"`
	SellPrice struct {
		Value          int `json:"value"`
		DisplayStrings struct {
			Header LocalizedField `json:"header"`
			Gold   LocalizedField `json:"gold"`
			Silver LocalizedField `json:"silver"`
			Copper LocalizedField `json:"copper"`
		} `json:"display_strings"`
	} `json:"sell_price,omitempty"`
	Requirements struct {
		Faction       *ItemFactionRequirement `json:"faction,omitempty"`
		PlayableRaces *struct {
			Links         Identifiables  `json:"links"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"playable_races,omitempty"`
		PlayableClasses *struct {
			Links         Identifiables  `json:"links"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"playable_classes,omitempty"`
		PlayableSpecializations *struct {
			Links         Identifiables  `json:"links"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"playable_specializations,omitempty"`
		Level      *ItemLevelRequirement      `json:"level,omitempty"`
		Reputation *ItemReputationRequirement `json:"reputation,omitempty"`
		Skill      *struct {
			Profession    *Identifiable  `json:"profession"`
			Level         int            `json:"level"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"skill,omitempty"`
		Ability *ItemAbilityRequirement `json:"ability,omitempty"`

		Map struct {
			Name LocalizedField `json:"name"`
			ID   int            `json:"id"`
		} `json:"map,omitempty"`
		Holiday struct {
			DisplayString LocalizedField `json:"display_string"`
		} `json:"holiday,omitempty"`
		Areas []struct {
			Name LocalizedField `json:"name"`
			ID   int            `json:"id"`
		} `json:"areas,omitempty"`
	} `json:"requirements,omitempty"`
	IsSubclassHidden bool            `json:"is_subclass_hidden"`
	Description      *LocalizedField `json:"description"`
	Binding          *Binding        `json:"binding"`
	Armor            *struct {
		Value   int                    `json:"value"`
		Display LocalizedDisplayString `json:"display"`
	} `json:"armor,omitempty"`
	Level *struct {
		Value         int            `json:"value"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"level,omitempty"`
	Durability *struct {
		Value         int            `json:"value"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"durability,omitempty"`
	NameDescription *LocalizedDisplayString `json:"name_description,omitempty"`
	UniqueEquipped  *LocalizedField         `json:"unique_equipped,omitempty"`
	Charges         *struct {
		Value         int            `json:"value"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"charges,omitempty"`
	Context *int        `json:"context,omitempty"`
	Stats   []*ItemStat `json:"stats,omitempty"`

	BonusList []int `json:"bonus_list,omitempty"`
	Weapon    struct {
		Damage struct {
			MinValue      int            `json:"min_value"`
			MaxValue      int            `json:"max_value"`
			DisplayString LocalizedField `json:"display_string"`
			DamageClass   struct {
				Type string         `json:"type"`
				Name LocalizedField `json:"name"`
			} `json:"damage_class"`
		} `json:"damage"`
		AttackSpeed struct {
			Value         int            `json:"value"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"attack_speed"`
		Dps struct {
			Value         int            `json:"value"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"dps"`
	} `json:"weapon"`
	Sockets []struct {
		SocketType struct {
			Type string         `json:"type"`
			Name LocalizedField `json:"name"`
		} `json:"socket_type"`
	} `json:"sockets"`
	SocketBonus LocalizedField `json:"socket_bonus"`
	Set         struct {
		ItemSet struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name LocalizedField `json:"name"`
			ID   int            `json:"id"`
		} `json:"item_set"`
		Items []struct {
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name LocalizedField `json:"name"`
				ID   int            `json:"id"`
			} `json:"item"`
		} `json:"items"`
		Effects []struct {
			DisplayString LocalizedField `json:"display_string"`
			RequiredCount int            `json:"required_count"`
		} `json:"effects"`
		Legacy        LocalizedField `json:"legacy"`
		DisplayString LocalizedField `json:"display_string"`
		Requirements  struct {
			Skill struct {
				Profession struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name LocalizedField `json:"name"`
					ID   int            `json:"id"`
				} `json:"profession"`
				Level         int            `json:"level"`
				DisplayString LocalizedField `json:"display_string"`
			} `json:"skill"`
		} `json:"requirements"`
	} `json:"set"`
	GemProperties struct {
		Effect       LocalizedField `json:"effect"`
		RelicType    LocalizedField `json:"relic_type"`
		MinItemLevel struct {
			Value         int            `json:"value"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"min_item_level"`
	} `json:"gem_properties"`
	ModifiedAppearanceID int `json:"modified_appearance_id"`
	AzeriteDetails       struct {
		SelectedPowersString LocalizedField `json:"selected_powers_string"`
		Level                struct {
			Value         int            `json:"value"`
			DisplayString LocalizedField `json:"display_string"`
		} `json:"level"`
	} `json:"azerite_details"`
	Recipe *struct {
		Item struct {
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"item"`
		} `json:"item"`
		Reagents []struct {
			Reagent struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name LocalizedField `json:"name"`
				ID   int            `json:"id"`
			} `json:"reagent"`
			Quantity int `json:"quantity"`
		} `json:"reagents"`
		ReagentsDisplayString LocalizedField `json:"reagents_display_string"`
	} `json:"recipe,omitempty"`
	ExpirationTimeLeft *struct {
		Value         int            `json:"value"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"expiration_time_left,omitempty"`
	LimitCategory LocalizedField `json:"limit_category"`
	ShieldBlock   struct {
		Value   int `json:"value"`
		Display struct {
			DisplayString LocalizedField `json:"display_string"`
			Color         Color          `json:"color"`
		} `json:"display"`
	} `json:"shield_block"`
	ItemStartsQuest struct {
		Quest struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name LocalizedField `json:"name"`
			ID   int            `json:"id"`
		} `json:"quest"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"item_starts_quest"`
	CraftingReagent *LocalizedField `json:"crafting_reagent,omitempty"`
	Toy             *LocalizedField `json:"toy,omitempty"`
	ContainerSlots  struct {
		Value         int            `json:"value"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"container_slots"`
	Upgrades struct {
		Value         int            `json:"value"`
		MaxValue      int            `json:"max_value"`
		DisplayString LocalizedField `json:"display_string"`
	} `json:"upgrades"`
	Conjured           LocalizedField `json:"conjured"`
	Legacy             LocalizedField `json:"legacy"`
	CreateLootSpecItem LocalizedField `json:"create_loot_spec_item"`
	IsCorrupted        bool           `json:"is_corrupted"`
}

type ItemMedia struct {
	Identifiable
	ItemID int          `pg:",pk"`
	Item   *Item        `pg:"rel:has-one"`
	Assets []ItemAssets `pg:"-"`
}

type ItemAssets struct {
	ItemMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Asset
}
