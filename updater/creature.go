package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func UpdateCreatureFamily(data *blizzard_api.ApiResponse) {
	var creatureFamily datasets.CreatureFamily
	data.Parse(&creatureFamily)
	if creatureFamily.PlayableSpecialization != nil {
		creatureFamily.PlayableSpecializationID = creatureFamily.PlayableSpecialization.ID
	}
	insertOnceUpdate(&creatureFamily, "name", "playable_specialization_id")

	if creatureFamily.Media != nil {
		creatureFamily.Media.CreatureFamilyID = creatureFamily.ID
		insertOnceUpdate(creatureFamily.Media, "creature_family_id")
	}
}

func UpdateCreatureFamilyMedia(data *blizzard_api.ApiResponse, id int) {
	var creatureFamilyMedia datasets.CreatureFamilyMedia
	data.Parse(&creatureFamilyMedia)

	for _, asset := range creatureFamilyMedia.Assets {
		asset.CreatureFamilyMediaID = id
		insertOnceExpr(&asset, "(creature_family_media_id,key) DO UPDATE", "value")
	}
}

func UpdateCreatureType(data *blizzard_api.ApiResponse) {
	var creatureType datasets.CreatureType
	data.Parse(&creatureType)
	insertOnceUpdate(&creatureType, "name")
}

func UpdateCreature(data *blizzard_api.ApiResponse) {
	var creature datasets.Creature
	data.Parse(&creature)

	if creature.Family != nil && creature.Family.ID != 57 {
		creature.FamilyID = creature.Family.ID
	}
	creature.TypeID = creature.Type.ID
	insertOnceUpdate(&creature, "name", "type_id", "family_id", "is_tameable")

	if creature.Media != nil {
		for _, media := range creature.Media {
			media.CreatureID = creature.ID
			insertOnceUpdate(&media, "creature_id")
		}
	}
}

func UpdateCreatureDisplayMedia(data *blizzard_api.ApiResponse, id int) {
	var creatureDisplayMedia datasets.CreatureDisplayMedia
	data.Parse(&creatureDisplayMedia)

	for _, asset := range creatureDisplayMedia.Assets {
		asset.CreatureDisplayMediaID = creatureDisplayMedia.ID
		insertOnceExpr(&asset, "(creature_display_media_id,key) DO UPDATE", "value")
	}
}