package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func UpdateMount(data *blizzard_api.ApiResponse) {
	var mount datasets.Mount
	data.Parse(&mount)

	if mount.Faction != nil {
		mount.FactionID = mount.Faction.ID
	}
	if mount.Source != nil {
		mount.SourceID = mount.Source.ID
		insertOnce(mount.Source)
	}

	insertOnceUpdate(&mount, "name", "description", "faction_id", "source_id")

	if mount.MountDisplays != nil{
		for _, media := range mount.MountDisplays {
			media.MountID = mount.ID
			insertOnceUpdate(&media, "mount_id")
		}
	}
}

func UpdateMountDisplayMedia(data *blizzard_api.ApiResponse, id int) {
	var mountDisplayMedia datasets.MountDisplayMedia
	data.Parse(&mountDisplayMedia)

	for _, asset := range mountDisplayMedia.Assets {
		asset.MountDisplayMediaID = mountDisplayMedia.ID
		insertOnceExpr(&asset, "(mount_display_media_id,key) DO UPDATE", "value")
	}
}

func UpdatePet(data *blizzard_api.ApiResponse) {
	var pet datasets.Pet
	data.Parse(&pet)

	if pet.Creature != nil {
		pet.CreatureID = pet.Creature.ID
	}
	if pet.Source != nil {
		pet.SourceID = pet.Source.ID
		insertOnce(pet.Source)
	}

	if pet.Source != nil {
		pet.SourceID = pet.Source.ID
		insertOnce(pet.Source)
	}

	if pet.BattlePetType != nil {
		pet.BattlePetTypeID = pet.BattlePetType.ID
		insertOnceUpdate(pet.BattlePetType,  "type", "name")
	}

	insertOnceUpdate(&pet, "name", "description", "battle_pet_type_id", "is_capturable", "is_tradable", "is_battlepet", "is_alliance_only", "is_horde_only", "source_id", "icon", "creature_id", "is_random_creature_display")

	for _, ability := range pet.Abilities {
		ability.AbilityID = ability.Ability.ID
		insertOnceUpdate(ability.Ability, "name")

		ability.PetID = pet.ID
		insertOnceExpr(&ability, "(pet_id,ability_id) DO UPDATE", "slot", "required_level")
	}
}