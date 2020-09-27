package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func UpdateAchievementCategory(data *blizzard_api.ApiResponse) {
	var achievementCategory datasets.AchievementCategory
	data.Parse(&achievementCategory)

	if achievementCategory.AggregatesByFaction != nil {
		achievementCategory.AllianceQuantity = achievementCategory.AggregatesByFaction.Alliance.Quantity
		achievementCategory.AlliancePoints = achievementCategory.AggregatesByFaction.Alliance.Points
		achievementCategory.HordeQuantity = achievementCategory.AggregatesByFaction.Horde.Quantity
		achievementCategory.HordePoints = achievementCategory.AggregatesByFaction.Horde.Points
	}

	insertOnceUpdate(&achievementCategory, "name", "is_guild_category", "display_order", "parent_category_id")
}

func UpdateParentCategory(data *blizzard_api.ApiResponse) {
	var achievementCategory datasets.AchievementCategory
	data.Parse(&achievementCategory)

	for _, category := range achievementCategory.Subcategories {
		connections.GetDBConn().
			Model(&datasets.AchievementCategory{}).
			Set("parent_category_id = ?", achievementCategory.ID).
			Where("id = ?", category.ID).
			Update()
	}
}

func updateAchievementCriteria(criteria *datasets.Criteria, parentCriteria int) {
	if criteria.Achievement != nil {
		insertOnce(criteria.Achievement)
		criteria.AchievementID = criteria.Achievement.ID
	}
	if criteria.Operator != nil {
		insertOnce(criteria.Operator)
		criteria.OperatorID = criteria.Operator.ID
	}
	if criteria.Faction != nil {
		criteria.FactionID = criteria.Faction.ID
	}
	if parentCriteria != 0 {
		criteria.ParentCriteriaID = parentCriteria
	}
	insertOnce(criteria)

	if criteria.ChildCriteria != nil {
		for _, cCriteria := range criteria.ChildCriteria {
			updateAchievementCriteria(cCriteria, criteria.ID)
		}
	}
}

func UpdateAchievement(data *blizzard_api.ApiResponse) {
	var achievement datasets.Achievement
	data.Parse(&achievement)

	if achievement.PrerequisiteAchievement != nil {
		insertOnce(achievement.PrerequisiteAchievement)
		achievement.PrerequisiteAchievementID = achievement.PrerequisiteAchievement.ID
	}
	if achievement.NextAchievement != nil {
		insertOnce(achievement.NextAchievement)
		achievement.NextAchievementID = achievement.NextAchievement.ID
	}
	if achievement.RewardItem != nil {
		achievement.RewardItemID = achievement.RewardItem.ID
	}

	achievement.CategoryID = achievement.Category.ID
	insertOnceUpdate(&achievement, "category_id", "name", "points", "is_account_wide", "prerequisite_achievement_id", "next_achievement_id", "display_order")
	achievement.Media.AchievementID = achievement.ID
	insertOnce(achievement.Media)

	if achievement.Criteria != nil {
		updateAchievementCriteria(achievement.Criteria, 0)
		insert(&datasets.AchievementCriteria{
			AchievementID: achievement.ID,
			CriteriaID:    achievement.Criteria.ID,
		})
	}
	if achievement.GuildRewardItems != nil {
		for _, item := range achievement.GuildRewardItems {
			insertOnceExpr(&datasets.GuildRewardItems{
				AchievementID: achievement.ID,
				ItemID:        item.ID,
			},
			"(achievement_id,item_id) DO NOTHING")
		}
	}
}

func UpdateAchievementMedia(data *blizzard_api.ApiResponse, id int) {
	var achievementMedia datasets.AchievementMedia
	data.Parse(&achievementMedia)

	for _, asset := range achievementMedia.Assets {
		asset.AchievementMediaID = achievementMedia.ID
		insertOnceExpr(&asset, "(achievement_media_id,key) DO UPDATE", "value")
	}
}