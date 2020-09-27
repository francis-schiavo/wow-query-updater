package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func UpdateJournalExpansion(data *blizzard_api.ApiResponse) {
	var journalExpansion datasets.JournalExpansion
	data.Parse(&journalExpansion)

	insertOnceUpdate(&journalExpansion, "name")
}

func UpdateJournalInstance(data *blizzard_api.ApiResponse) {
	var journalInstance datasets.JournalInstance
	data.Parse(&journalInstance)

	journalInstance.ExpansionID = journalInstance.Expansion.ID

	if journalInstance.Category != nil {
		journalInstance.CategoryID = journalInstance.Category.ID
		insertOnce(journalInstance.Category)
	}

	if journalInstance.Map != nil {
		journalInstance.MapID = journalInstance.Map.ID
		insertOnceUpdate(journalInstance.Map, "name")
	}

	if journalInstance.Area != nil {
		journalInstance.AreaID = journalInstance.Area.ID
		insertOnceUpdate(journalInstance.Area, "name")
	}

	if journalInstance.Location != nil {
		journalInstance.LocationID = journalInstance.Location.ID
		insertOnceUpdate(journalInstance.Location, "name")
	}

	insertOnceUpdate(&journalInstance, "name", "description", "minimum_level", "category_id", "expansion_id", "map_id", "area_id", "location_id")

	for _, mode := range journalInstance.Modes {
		insertOnce(mode.Mode)

		mode.InstanceID = journalInstance.ID
		mode.ModeID = mode.Mode.ID
		insertOnceExpr(&mode, "(instance_id,mode_id) DO UPDATE", "players", "is_tracked")
	}

	if journalInstance.Media != nil {
		journalInstance.Media.JournalID = journalInstance.ID
		insertOnce(journalInstance.Media)
	}
}

func updateEncounterSection(sections []datasets.JournalEncounterSection, parentSection *int) {
	for _, section := range sections {
		if parentSection != nil {
			section.ParentSectionID = *parentSection
		}

		if section.CreatureDisplayMedia != nil {
			insertOnce(section.CreatureDisplayMedia)
			section.CreatureDisplayMediaID = section.CreatureDisplayMedia.ID
		}

		insertOnceUpdate(&section, "parent_section_id", "title", "body_text", "creature_display_media_id")

		if section.Sections != nil {
			updateEncounterSection(section.Sections, &section.ID)
		}
	}
}

func UpdateJournalEncounter(data *blizzard_api.ApiResponse) {
	var journalEncounter datasets.JournalEncounter
	data.Parse(&journalEncounter)

	journalEncounter.InstanceID = journalEncounter.Instance.ID
	journalEncounter.CategoryID = journalEncounter.Category.ID

	insertOnceUpdate(&journalEncounter, "name", "description", "instance_id", "category_id")

	for _, creature := range journalEncounter.Creatures {
		creature.EncounterID = journalEncounter.ID
		insertOnceExpr(&creature, "(encounter_id,creature_id) DO NOTHING")
	}

	for _, items := range journalEncounter.Items {
		items.EncounterID = journalEncounter.ID
		items.ItemID = items.Item.ID
		insertOnceUpdate(&items, "encounter_id", "item_id")
	}

	for _, mode := range journalEncounter.Modes {
		insertOnceExpr(&datasets.JournalEncounterModes{
			EncounterID: journalEncounter.ID,
			ModeID:      mode.ID,
		}, "(encounter_id,mode_id) DO NOTHING")
	}

	updateEncounterSection(journalEncounter.Sections, nil)
}

func UpdateInstanceMedia(data *blizzard_api.ApiResponse, id int) {
	var journalInstanceMedia datasets.JournalInstanceMedia
	data.Parse(&journalInstanceMedia)

	for _, asset := range journalInstanceMedia.Assets {
		asset.JournalInstanceMediaID = id
		insertOnceExpr(&asset, "(journal_instance_media_id,key) DO UPDATE", "value")
	}
}