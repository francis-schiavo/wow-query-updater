package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func UpdateReputationTier(data *blizzard_api.ApiResponse) {
	var reputationTierGroup datasets.ReputationTierGroup
	data.Parse(&reputationTierGroup)

	for _, tier := range reputationTierGroup.Tiers {
		tier.ReputationTierID = reputationTierGroup.ID
		insertOnceExpr(&tier, "(id,reputation_tier_id) DO UPDATE", "name", "min_value", "max_value")
	}
}

func UpdateReputationFaction(data *blizzard_api.ApiResponse) {
	var reputationFaction datasets.ReputationFaction
	data.Parse(&reputationFaction)

	insertOnceUpdate(&reputationFaction, "reputation_tier_id", "parent_faction_id", "name", "description")
}

func UpdateParentReputation(data *blizzard_api.ApiResponse) {
	var reputationFaction datasets.ReputationFaction
	data.Parse(&reputationFaction)

	insertOnceUpdate(&reputationFaction, "reputation_tier_id", "parent_faction_id", "name", "description")

	for _, faction := range reputationFaction.ReputationFaction {
		connections.GetDBConn().
			Model(&datasets.ReputationFaction{}).
			Set("parent_faction_id = ?", reputationFaction.ID).
			Where("id = ?", faction.ID).
			Update()
	}
}