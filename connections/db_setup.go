package connections

import (
	"github.com/go-pg/pg/v10/orm"
	"wow-query-updater/datasets"
)

func DatabaseSetup(classic bool) {
	createTable(&datasets.UpdateError{})
	createTable(&datasets.Cache{})
	// Enums / simple references
	createTable(&datasets.Faction{})
	createTable(&datasets.Operator{})
	createTable(&datasets.Role{})
	createTable(&datasets.Source{})
	createTable(&datasets.Binding{})
	createTable(&datasets.InventoryType{})
	createTable(&datasets.ItemQuality{})
	createTable(&datasets.Stat{})
	createTable(&datasets.Currency{})
	createTable(&datasets.Follower{})

	// Character related
	createTable(&datasets.PowerType{})
	createTable(&datasets.PlayableRace{})
	createTable(&datasets.PlayableClass{})
	createTable(&datasets.PlayableClassMedia{})
	createTable(&datasets.PlayableClassAssets{})
	createTable(&datasets.PlayableSpecialization{})

	// Creatures
	createTable(&datasets.CreatureFamily{})
	createTable(&datasets.CreatureFamilyMedia{})
	createTable(&datasets.CreatureFamilyAssets{})
	createTable(&datasets.CreatureType{})
	createTable(&datasets.Creature{})
	createTable(&datasets.CreatureDisplayMedia{})
	createTable(&datasets.CreatureDisplayAssets{})

	createTable(&datasets.ItemClass{})
	createTable(&datasets.ItemSubclass{})

	createTable(&datasets.Item{})
	createTable(&datasets.ItemMedia{})
	createTable(&datasets.ItemAssets{})

	if !classic {
		createTable(&datasets.ProfessionType{})
		createTable(&datasets.Profession{})
		createTable(&datasets.ProfessionTier{})
		createTable(&datasets.ProfessionCategory{})

		createTable(&datasets.Spell{})
		createTable(&datasets.SpellTooltip{})
		createTable(&datasets.SpellMedia{})
		createTable(&datasets.SpellAssets{})

		createTable(&datasets.ReputationTier{})
		createTable(&datasets.ReputationFaction{})


		createTable(&datasets.ItemStat{})
		createTable(&datasets.ItemFactionRequirement{})
		createTable(&datasets.ItemRaceRequirement{})
		createTable(&datasets.ItemClassRequirement{})
		createTable(&datasets.ItemSpecializationRequirement{})
		createTable(&datasets.ItemLevelRequirement{})
		createTable(&datasets.ItemReputationRequirement{})
		createTable(&datasets.ItemAbilityRequirement{})
		createTable(&datasets.ItemSkillRequirement{})
		createTable(&datasets.ItemMetadata{})

		createTable(&datasets.PlayableSpecializationMedia{})
		createTable(&datasets.PlayableSpecializationAsset{})

		createTable(&datasets.AzeriteEssence{})
		createTable(&datasets.AzeritePower{})
		createTable(&datasets.AzeriteEssencePower{})
		createTable(&datasets.AzeriteEssenceSpecializations{})
		createTable(&datasets.AzeriteEssenceMedia{})
		createTable(&datasets.AzeriteEssenceAsset{})

		createTable(&datasets.Talent{})
		createTable(&datasets.PvpTalent{})

		createTable(&datasets.AchievementCategory{})
		createTable(&datasets.Achievement{})
		createTable(&datasets.AchievementMedia{})
		createTable(&datasets.AchievementAssets{})

		createTable(&datasets.Criteria{})
	    createTable(&datasets.AchievementCriteria{})
		createTable(&datasets.GuildRewardItems{})

		createTable(&datasets.Title{})

		createTable(&datasets.Source{})
		createTable(&datasets.Mount{})
		createTable(&datasets.MountDisplayMedia{})
		createTable(&datasets.MountDisplayAssets{})

		createTable(&datasets.QuestCategory{})
		createTable(&datasets.QuestArea{})
		createTable(&datasets.QuestType{})
		createTable(&datasets.Quest{})
		createTable(&datasets.QuestReward{})
		createTable(&datasets.QuestReputationReward{})
		createTable(&datasets.QuestItemReward{})
		createTable(&datasets.QuestItemChoiceReward{})
		createTable(&datasets.QuestFollowerReward{})
		createTable(&datasets.QuestCurrencyReward{})
		createTable(&datasets.QuestRequirement{})
		createTable(&datasets.QuestClassRequirement{})
		createTable(&datasets.QuestRaceRequirement{})
		createTable(&datasets.QuestReputationRequirement{})

		createTable(&datasets.Recipe{})
		createTable(&datasets.RecipeReagents{})
		createTable(&datasets.RecipeMedia{})
		createTable(&datasets.RecipeAssets{})

		createTable(&datasets.ProfessionTierRecipes{})
		createTable(&datasets.ProfessionMedia{})
		createTable(&datasets.ProfessionAssets{})

		createTable(&datasets.JournalExpansion{})
		createTable(&datasets.JournalMode{})
		createTable(&datasets.JournalMap{})
		createTable(&datasets.JournalArea{})
		createTable(&datasets.JournalLocation{})
		createTable(&datasets.JournalCategory{})
		createTable(&datasets.JournalInstance{})
		createTable(&datasets.JournalInstanceMode{})
		createTable(&datasets.JournalInstanceMedia{})
		createTable(&datasets.JournalInstanceAssets{})
		createTable(&datasets.JournalEncounter{})
		createTable(&datasets.JournalEncounterItems{})
		createTable(&datasets.JournalEncounterCreatures{})
		createTable(&datasets.JournalEncounterSection{})
		createTable(&datasets.JournalEncounterModes{})

		createTable(&datasets.BattlePetType{})
		createTable(&datasets.BattlePetAbility{})
		createTable(&datasets.Pet{})
		createTable(&datasets.PetAbility{})

		createTable(&datasets.TechTalentTree{})
		createTable(&datasets.TechTalent{})
		createTable(&datasets.TechTalentMedia{})
		createTable(&datasets.TechTalentAssets{})

		createTable(&datasets.SocketType{})
		createTable(&datasets.Conduit{})
		createTable(&datasets.ConduitRank{})

		createTable(&datasets.Covenant{})
		createTable(&datasets.CovenantAbility{})
		createTable(&datasets.RenownReward{})
		createTable(&datasets.CovenantRenownReward{})
		createTable(&datasets.CovenantMedia{})
		createTable(&datasets.CovenantAssets{})
		createTable(&datasets.Soulbind{})
	}
}

func createTable(model interface{}) {
	err := GetDBConn().Model(model).CreateTable(&orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	if err != nil {
		panic(err)
	}
}