package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func UpdateTechTalentTree(data *blizzard_api.ApiResponse) {
	var techTalentTree datasets.TechTalentTree
	data.Parse(&techTalentTree)

	// Preload the tech talents
	for _, talent := range techTalentTree.Talents {
		insertOnce(&datasets.TechTalent{
			Identifiable: talent,
			Name: nil,
			Description: nil,
		})
	}

	insertOnceUpdate(&techTalentTree, "max_tiers")
}

func UpdateTechTalent(data *blizzard_api.ApiResponse) {
	var techTalent datasets.TechTalent
	data.Parse(&techTalent)

	techTalent.TechTalentTreeID = techTalent.TechTalentTree.ID
	techTalent.SpellTooltipID = techTalent.SpellTooltip.ID

	if techTalent.PrerequisiteTalent != nil {
		techTalent.PrerequisiteTalentID = techTalent.PrerequisiteTalent.ID
	}

	if techTalent.SpellTooltip != nil {
		updateSpellTooltip(techTalent.SpellTooltip)
	}

	insertOnceUpdate(&techTalent, "name", "description", "spell_tooltip_id", "tier", "display_order", "prerequisite_talent_id", "tech_talent_tree_id")
}