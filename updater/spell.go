package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func updateSpellTooltip(tooltip *datasets.SpellTooltip) {
	insertOnceUpdate(tooltip, "description", "cast_time", "cooldown", "range", "power_cost")
}

func UpdateSpell(data *blizzard_api.ApiResponse) {
	var spell datasets.Spell
	data.Parse(&spell)

	if spell.ID == 0 {
		return
	}

	insertOnceUpdate(&spell, "name", "description")
	spell.Media.SpellID = spell.ID
	insertOnceUpdate(spell.Media, "spell_id")
}

func UpdateSpellMedia(data *blizzard_api.ApiResponse, id int) {
	var spell datasets.SpellMedia
	data.Parse(&spell)
	for _, asset := range spell.Assets {
		asset.SpellMediaID = spell.ID
		insertOnceExpr(&asset, "(spell_media_id,key) DO UPDATE", "value")
	}
}
