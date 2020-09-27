package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func UpdatePowerType(data *blizzard_api.ApiResponse) {
	var power datasets.PowerType
	data.Parse(&power)
	insert(&power)
}

func UpdatePlayableRace(data *blizzard_api.ApiResponse) {
	var race datasets.PlayableRace
	data.Parse(&race)

	insertOnce(race.Faction)
	race.FactionID = race.Faction.ID
	insertOnceUpdate(&race, "name", "gender_name", "faction_id", "is_selectable", "is_allied_race")
}

func UpdatePlayableClass(data *blizzard_api.ApiResponse) {
	var class datasets.PlayableClass
	data.Parse(&class)

	class.PowerTypeID = class.PowerType.ID
	insertOnceUpdate(&class, "name", "gender_name", "power_type_id")
	class.Media.PlayableClassID = class.ID
	insertOnceUpdate(class.Media, "playable_class_id")
}

func UpdatePlayableClassMedia(data *blizzard_api.ApiResponse, id int) {
	var class datasets.PlayableClassMedia
	data.Parse(&class)

	for _, asset := range class.Assets {
		asset.PlayableClassMediaID = class.ID
		insertOnceExpr(&asset, "(playable_class_media_id,key) DO UPDATE", "value")
	}
}

func UpdatePlayableSpecialization(data *blizzard_api.ApiResponse) {
	var playableSpecialization datasets.PlayableSpecialization
	data.Parse(&playableSpecialization)

	insertOnceUpdate(&playableSpecialization.Role, "name")

	playableSpecialization.RoleID = playableSpecialization.Role.ID
	playableSpecialization.PlayableClassID = playableSpecialization.PlayableClass.ID
	insertOnceUpdate(&playableSpecialization, "playable_class_id", "name", "gender_description", "role_id")
	playableSpecialization.Media.PlayableSpecializationID = playableSpecialization.ID
	insertOnceUpdate(playableSpecialization.Media, "playable_specialization_id")
}

func UpdatePlayableSpecializationMedia(data *blizzard_api.ApiResponse, id int) {
	var playableSpecializationMedia datasets.PlayableSpecializationMedia
	data.Parse(&playableSpecializationMedia)

	for _, asset := range playableSpecializationMedia.Assets {
		asset.PlayableSpecializationMediaID = playableSpecializationMedia.ID
		insertOnceExpr(&asset, "(playable_specialization_media_id,key) DO UPDATE", "value")
	}
}

func UpdateSpell(data *blizzard_api.ApiResponse) bool {
	var spell datasets.Spell
	data.Parse(&spell)

	if spell.ID == 0 {
		return false
	}

	insertOnceUpdate(&spell, "name", "description")
	spell.Media.SpellID = spell.ID
	insertOnceUpdate(spell.Media, "spell_id")
	return true
}

func UpdateSpellMedia(data *blizzard_api.ApiResponse, id int) {
	var spell datasets.SpellMedia
	data.Parse(&spell)
	for _, asset := range spell.Assets {
		asset.SpellMediaID = spell.ID
		insertOnceExpr(&asset, "(spell_media_id,key) DO UPDATE", "value")
	}
}

func UpdateTalent(data *blizzard_api.ApiResponse) {
	var talent datasets.Talent
	data.Parse(&talent)

	talent.PlayableClassID = talent.PlayableClass.ID
	talent.SpellID = talent.Spell.ID

	spellResponse := connections.WowClient.Spell(talent.SpellID, nil)
	UpdateSpell(spellResponse)

	insertOnceUpdate(&talent, "tier_index", "column_index", "level", "spell_id", "playable_class_id")
}

func UpdatePvpTalent(data *blizzard_api.ApiResponse) {
	var talent datasets.PvpTalent
	data.Parse(&talent)

	talent.PlayableSpecializationID = talent.PlayableSpecialization.ID
	talent.SpellID = talent.Spell.ID
	spellResponse := connections.WowClient.Spell(talent.SpellID, nil)
	UpdateSpell(spellResponse)

	if talent.OverridesSpell != nil {
		talent.OverridesSpellID = talent.OverridesSpell.ID
		spellResponse := connections.WowClient.Spell(talent.OverridesSpellID, nil)
		UpdateSpell(spellResponse)
	}

	insertOnceUpdate(&talent, "spell_id", "overrides_spell_id", "description", "unlock_player_level", "compatible_slots")
}

func UpdateTitle(data *blizzard_api.ApiResponse) {
	var title datasets.Title
	data.Parse(&title)
	insertOnceUpdate(&title, "name", "gender_name")
}