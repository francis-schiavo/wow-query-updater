package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func UpdateQuest(data *blizzard_api.ApiResponse) {
	var quest datasets.Quest
	data.Parse(&quest)

	if quest.Area != nil {
		quest.AreaID = quest.Area.ID
		insertOnce(quest.Area)
	}
	if quest.Category != nil {
		quest.CategoryID = quest.Category.ID
		insertOnce(quest.Category)
	}
	if quest.Type != nil {
		quest.TypeID = quest.Type.ID
		insertOnce(quest.Type)
	}

	if quest.Requirements.Faction != nil {
		quest.Requirements.FactionID = quest.Requirements.Faction.ID
	}

 	insertOnceUpdate(&quest, "area_id", "category_id", "type_id", "title", "description", "recommended_minimum_level", "recommended_maximum_level", "is_daily", "is_weekly", "is_repeatable")

	if quest.Rewards != nil {
		questReward := &datasets.QuestReward{
			QuestID:       quest.ID,
			Experience:    quest.Rewards.Experience,
			Money:         quest.Rewards.MoneyObj.Value,
			Honor:         quest.Rewards.Honor,
			ArtifactPower: quest.Rewards.ArtifactPower,
		}
		if quest.Rewards.Title != nil {
			questReward.TitleID = quest.Rewards.Title.ID
		}
		insertOnceExpr(questReward, "(quest_id) DO UPDATE", "experience", "money")
	}

	for _, reputation := range quest.Rewards.Reputations {
		reputation.RewardID = reputation.Reward.ID
		reputation.QuestID = quest.ID
		insertOnceExpr(reputation, "(quest_id, reward_id) DO UPDATE", "value")
	}

	for _, class := range quest.Requirements.Classes {
		class.QuestID = quest.ID
		insertOnceExpr(class, "(quest_id, playable_class_id) DO NOTHING")
	}

	for _, race := range quest.Requirements.Races {
		race.QuestID = quest.ID
		insertOnceExpr(race, "(quest_id, playable_race_id) DO NOTHING")
	}
}

func UpdateQuestCategory(data *blizzard_api.ApiResponse) {
	var questCategory datasets.QuestCategory
	data.Parse(&questCategory)
	insertOnceUpdate(&questCategory, "category")

	for _, quest := range questCategory.Quests {
		questResponse := connections.WowClient.Quest(quest.ID, nil)
		UpdateQuest(questResponse)
	}
}

func UpdateQuestType(data *blizzard_api.ApiResponse) {
	var questType datasets.QuestType
	data.Parse(&questType)
	insertOnceUpdate(&questType, "type")
	for _, quest := range questType.Quests {
		questResponse := connections.WowClient.Quest(quest.ID, nil)
		UpdateQuest(questResponse)
	}
}

func UpdateQuestArea(data *blizzard_api.ApiResponse) {
	var questArea datasets.QuestArea
	data.Parse(&questArea)
	insertOnceUpdate(&questArea, "area")
	for _, quest := range questArea.Quests {
		questResponse := connections.WowClient.Quest(quest.ID, nil)
		UpdateQuest(questResponse)
	}
}