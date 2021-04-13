package datasets

type ModifiedCraftingCategory NamedItem

type ModifiedCraftingReagentSlotType struct {
	Identifiable
	Description          LocalizedField `json:"description"`
	CompatibleCategories Identifiables  `json:"compatible_categories"`
}

type ModifiedCraftingReagentSlotTypeCompatibleCategories struct {
	ModifiedCraftingCategoryID        int                              ``
	ModifiedCraftingCategory          *ModifiedCraftingCategory        `pg:"rel:has-one"`
	ModifiedCraftingReagentSlotTypeID int                              ``
	ModifiedCraftingReagentSlotType   *ModifiedCraftingReagentSlotType `pg:"rel:has-one"`
}
