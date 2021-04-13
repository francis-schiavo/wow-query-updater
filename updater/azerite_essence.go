package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func updateAzeritePower(essence *datasets.AzeriteEssence) {
	for _, power := range essence.Powers {
		spell := connections.WowClient.Spell(power.MainPowerSpell.ID, nil)
		UpdateSpell(spell)
		power.MainPowerSpellID = power.MainPowerSpell.ID

		spell = connections.WowClient.Spell(power.PassivePowerSpell.ID, nil)
		UpdateSpell(spell)
		power.PassivePowerSpellID = power.PassivePowerSpell.ID

		insertOnceUpdate(&power, "rank", "main_power_spell_id", "passive_power_spell_id")
		insertOnceExpr(&datasets.AzeriteEssencePower{
			AzeriteEssenceID: essence.ID,
			AzeritePowerID: power.ID,
		}, "(azerite_essence_id,azerite_power_id) DO NOTHING")
	}
}

func updateAzeriteSpecializations(essence *datasets.AzeriteEssence) {
	for _, spec := range essence.AllowedSpecializations {
		insertOnceExpr(&datasets.AzeriteEssenceSpecializations{
			AzeriteEssenceID:         essence.ID,
			PlayableSpecializationID: spec.ID,
		}, "(azerite_essence_id,playable_specialization_id) DO NOTHING")
	}
}

func UpdateAzeriteEssence(data *blizzard_api.ApiResponse) {
	var azeriteEssence datasets.AzeriteEssence
	data.Parse(&azeriteEssence)

	insertOnceUpdate(&azeriteEssence, "name")

	azeriteEssence.Media.AzeriteEssenceID = azeriteEssence.ID
	insertOnceUpdate(azeriteEssence.Media, "azerite_essence_id")

	updateAzeritePower(&azeriteEssence)
	updateAzeriteSpecializations(&azeriteEssence)
}

func UpdateAzeriteEssenceMedia(data *blizzard_api.ApiResponse, id int) {
	var azeriteEssenceMedia datasets.AzeriteEssenceMedia
	data.Parse(&azeriteEssenceMedia)

	for _, asset := range azeriteEssenceMedia.Assets {
		asset.AzeriteEssenceMediaID = id
		insert(&asset)
	}
}