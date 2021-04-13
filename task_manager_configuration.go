package main

import (
	"wow-query-updater/datasets"
	"wow-query-updater/updater"
)

func SetupTaskManager(concurrency int, logLevel updater.LogType, classic bool) *updater.TaskManager {
	taskManager := updater.NewTaskManager(concurrency, logLevel)

	// Common
	taskManager.AddIndexTask("playable_race", "PlayableRaceIndex", "races", "PlayableRace", updater.UpdatePlayableRace)

	taskManager.AddIndexTask("power_type", "PowerTypeIndex", "power_types", "PowerType", updater.UpdatePowerType)
	taskManager.AddIndexTask("playable_class", "PlayableClassIndex", "classes", "PlayableClass", updater.UpdatePlayableClass)
	taskManager.AddMediaTask("playable_class_assets", &datasets.PlayableClassMedia{}, "PlayableClassMedia", updater.UpdatePlayableClassMedia)
	taskManager.AddIndexTask("playable_specialization", "PlayableSpecializationIndex", "character_specializations", "PlayableSpecialization", updater.UpdatePlayableSpecialization)
	taskManager.AddIndexTask("playable_pet_specialization", "PlayableSpecializationIndex", "pet_specializations", "PlayableSpecialization", updater.UpdatePlayableSpecialization)

	// Creature
	taskManager.AddIndexTask("creature_family", "CreatureFamilyIndex", "creature_families", "CreatureFamily", updater.UpdateCreatureFamily)
	taskManager.AddIndexTask("creature_type", "CreatureTypeIndex", "creature_types", "CreatureType", updater.UpdateCreatureType)
	taskManager.AddSearchTask("creature", "CreatureSearch", "Creature", updater.UpdateCreature)

	taskManager.AddIndexTaskLimited("item_class", "ItemClassIndex", "item_classes", "ItemClass", updater.UpdateItemClass, 50)
	taskManager.AddSimpleTask("workaround_add_missing_classes", updater.InsertMissingItemClasses)
	taskManager.AddSimpleTask("workaround_add_missing_stats", updater.InsertMissingStats)

	// Item
	if classic {
		taskManager.AddSearchTask("item", "ItemSearch", "Item", updater.UpdateItem)
	} else{
		//Preload profession
		taskManager.AddIndexTask("profession", "ProfessionIndex", "professions", "Profession", updater.UpdateProfession)

		// Reputation
		taskManager.AddIndexTask("reputation_tier", "ReputationTierIndex", "reputation_tiers", "ReputationTier", updater.UpdateReputationTier)
		taskManager.AddIndexTask("reputation_faction", "ReputationFactionIndex", "root_factions", "ReputationFaction", updater.UpdateReputationFaction)
		taskManager.AddIndexTask("reputation_faction", "ReputationFactionIndex", "factions", "ReputationFaction", updater.UpdateParentReputation)
		taskManager.AddSimpleTask("workaround_add_missing_reputation_tiers", updater.InsertMissingReputationTiers)

		// Spell
		taskManager.AddSearchTask("spell", "SpellSearch", "Spell", updater.UpdateSpell)
		taskManager.AddSimpleTask("workaround_add_missing_spells", updater.InsertMissingSpells)

		// Items
		taskManager.AddSearchTask("item", "ItemSearch", "Item", updater.UpdateItem)

		// Common
		taskManager.AddIndexTaskLimited("talents", "TalentIndex", "talents", "Talent", updater.UpdateTalent, 20)
		taskManager.AddMediaTask("playable_specialization_media", &datasets.PlayableSpecializationMedia{}, "PlayableSpecializationMedia", updater.UpdatePlayableSpecializationMedia)
		taskManager.AddIndexTaskLimited("pvp_talents", "PvPTalentIndex", "pvp_talents", "PvPTalent", updater.UpdatePvpTalent, 20)

		taskManager.AddIndexTask("title", "TitleIndex", "titles", "Title", updater.UpdateTitle)

		// Azerite
		taskManager.AddIndexTaskLimited("azerite_essence", "AzeriteEssenceIndex", "azerite_essences", "AzeriteEssence", updater.UpdateAzeriteEssence, 20)
		taskManager.AddMediaTask("azerite_essence_media", &datasets.AzeriteEssenceMedia{}, "AzeriteEssenceMedia", updater.UpdateAzeriteEssenceMedia)

		// Achievement
		taskManager.AddIndexTask("root_achievement_category", "AchievementCategoryIndex", "root_categories", "AchievementCategory", updater.UpdateAchievementCategory)
		taskManager.AddIndexTask("guild_achievement_category", "AchievementCategoryIndex", "guild_categories", "AchievementCategory", updater.UpdateAchievementCategory)
		taskManager.AddIndexTask("achievement_category", "AchievementCategoryIndex", "categories", "AchievementCategory", updater.UpdateAchievementCategory)

		taskManager.AddIndexTask("root_achievement_category_update_parent", "AchievementCategoryIndex", "root_categories", "AchievementCategory", updater.UpdateParentCategory)
		taskManager.AddIndexTask("guild_achievement_category_update_parent", "AchievementCategoryIndex", "guild_categories", "AchievementCategory", updater.UpdateParentCategory)
		taskManager.AddIndexTask("achievement_category_update_parent", "AchievementCategoryIndex", "categories", "AchievementCategory", updater.UpdateParentCategory)

		taskManager.AddIndexTask("achievement", "AchievementIndex", "achievements", "Achievement", updater.UpdateAchievement)
		taskManager.AddMediaTask("achievement_assets", &datasets.AchievementMedia{}, "AchievementMedia", updater.UpdateAchievementMedia)

		// Quest
		taskManager.AddIndexTaskLimited("quest_category", "QuestCategoryIndex", "categories", "QuestCategory", updater.UpdateQuestCategory, 50)
		taskManager.AddIndexTaskLimited("quest_type", "QuestTypeIndex", "types", "QuestType", updater.UpdateQuestType, 50)
		taskManager.AddIndexTaskLimited("quest_area", "QuestAreaIndex", "areas", "QuestArea", updater.UpdateQuestArea, 50)

		// Collections
		taskManager.AddIndexTask("mount", "MountIndex", "mounts", "Mount", updater.UpdateMount)
		taskManager.AddMediaTask("mount_media", &datasets.MountDisplayMedia{}, "CreatureDisplayMedia", updater.UpdateMountDisplayMedia)
		taskManager.AddIndexTask("pet", "PetIndex", "pets", "Pet", updater.UpdatePet)

		// Profession
		taskManager.AddIndexTaskLimited("profession", "ProfessionIndex", "professions", "Profession", updater.UpdateProfessionTiers, 30)
		taskManager.AddMediaTask("profession_media", &datasets.ProfessionMedia{}, "ProfessionMedia", updater.UpdateProfessionMedia)
		taskManager.AddMediaTask("recipe_media", &datasets.RecipeMedia{}, "RecipeMedia", updater.UpdateRecipeMedia)

		//Journal
		taskManager.AddIndexTask("journal_expansion", "JournalExpansionIndex", "tiers", "JournalExpansion", updater.UpdateJournalExpansion)
		taskManager.AddIndexTask("journal_instance", "JournalInstanceIndex", "instances", "JournalInstance", updater.UpdateJournalInstance)
		taskManager.AddIndexTask("journal_encounter", "JournalEncounterIndex", "encounters", "JournalEncounter", updater.UpdateJournalEncounter)
		taskManager.AddMediaTask("journal_instance_media", &datasets.JournalInstanceMedia{}, "JournalInstanceMedia", updater.UpdateInstanceMedia)
		//
		// Tech talent
		taskManager.AddIndexTask("tech_talent_tree", "TechTalentTreeIndex", "talent_trees", "TechTalentTree", updater.UpdateTechTalentTree)
		//taskManager.AddIndexTask("tech talent", "TechTalentIndex", "talents", "TechTalent", updater.UpdateTechTalent)
		taskManager.AddIndexTask("workaround_tech_talent", "TechTalentTreeIndex", "talent_trees", "TechTalentTree", updater.UpdateTechTalentUsingTree)

		// Covenant
		taskManager.AddIndexTask("conduit", "ConduitIndex", "conduits", "Conduit", updater.UpdateConduit)
		taskManager.AddIndexTask("covenant", "CovenantIndex", "covenants", "Covenant", updater.UpdateCovenant)
		taskManager.AddIndexTask("soulbind", "SoulbindIndex", "soulbinds", "Soulbind", updater.UpdateSoulbind)
		taskManager.AddMediaTask("covenant_media", &datasets.CovenantMedia{}, "CovenantMedia", updater.UpdateCovenantMedia)
	}

	// Shared Media
	taskManager.AddMediaTask("creature_media", &datasets.CreatureDisplayMedia{}, "CreatureDisplayMedia", updater.UpdateCreatureDisplayMedia)
	taskManager.AddMediaTask("item_media", &datasets.ItemMedia{}, "ItemMedia", updater.UpdateItemMedia)
	taskManager.AddMediaTask("creature_family_media", &datasets.CreatureFamilyMedia{}, "CreatureFamilyMedia", updater.UpdateCreatureFamilyMedia)

	if !classic {
		taskManager.AddMediaTask("spell_media", &datasets.SpellMedia{}, "SpellMedia", updater.UpdateSpellMedia)
	}

	return taskManager
}