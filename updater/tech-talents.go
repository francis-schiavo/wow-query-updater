package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func UpdateTechTalentsTrees(data *blizzard_api.ApiResponse) {
	var techTalentTree datasets.TechTalentTree
	data.Parse(&techTalentTree)
	insertOnceUpdate(&techTalentTree, "max_tiers", "talents")
}

func UpdateTechTalents(data *blizzard_api.ApiResponse) {
	var techTalent datasets.TechTalent
	data.Parse(&techTalent)
	techTalent.TalentTreeID = techTalent.TalentTree.ID
	insertOnceUpdate(&techTalent, "talent_tree_id", "name", "description", "spell_tooltip", "tier", "display_order", "prerequisite_talent")
}

func UpdateTechTalentMedia(data *blizzard_api.ApiResponse, id int) {
	var techTalentMedia datasets.TechTalentMedia
	data.Parse(&techTalentMedia)

	for _, asset := range techTalentMedia.Assets {
		asset.TechTalentMediaID = techTalentMedia.ID
		insertOnceExpr(&asset, "(tech_talent_media_id,key) DO UPDATE", "value")
	}
}
