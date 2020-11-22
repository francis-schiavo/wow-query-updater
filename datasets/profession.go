package datasets

type ProfessionType Enum

type Profession struct {
	Identifiable
	Name        LocalizedField   `json:"name"`
	Description LocalizedField   `json:"description"`
	TypeID      string           ``
	Type        *ProfessionType  `json:"type" pg:"rel:has-one"`
	SkillTiers  Identifiables    `json:"skill_tiers" pg:"-"`
	Media       *ProfessionMedia `json:"media" pg:"-"`
}

type ProfessionTier struct {
	Identifiable
	ProfessionID      int                  ``
	Profession        *Profession          `pg:"rel:has-one"`
	Name              LocalizedField       `json:"name"`
	MinimumSkillLevel int                  `json:"minimum_skill_level"`
	MaximumSkillLevel int                  `json:"maximum_skill_level"`
	Categories        []ProfessionCategory `json:"categories" pg:"-"`
}

type ProfessionCategory struct {
	Identifiable
	TierID  int             ``
	Tier    *ProfessionTier `pg:"rel:has-one"`
	Name    LocalizedField  `json:"name"`
	Recipes []Recipe        `json:"recipes" pg:"-"`
}

type ProfessionTierRecipes struct {
	ProfessionTierID int                 `pg:",pk"`
	ProfessionTier   *ProfessionTier     `pg:"rel:has-one"`
	RecipeID         int                 `pg:",pk"`
	Recipe           *Recipe             `pg:"rel:has-one"`
	CategoryID       int                 ``
	Category         *ProfessionCategory `pg:"rel:has-one"`
}

type ProfessionMedia struct {
	Identifiable
	ProfessionID int                ``
	Profession   *Profession        `pg:"rel:has-one"`
	Assets       []ProfessionAssets `pg:"-"`
}

type ProfessionAssets struct {
	ProfessionMediaID int              `pg:",pk"`
	ProfessionMedia   *ProfessionMedia `pg:"rel:has-one"`
	Asset
}

type Recipe struct {
	Identifiable
	Name          LocalizedField   `json:"name"`
	CraftedItemID int              ``
	CraftedItem   *Item            `json:"crafted_item" pg:"rel:has-one"`
	Reagents      []RecipeReagents `json:"reagents" pg:"-"`
	Media         *RecipeMedia     `json:"media" pg:"-"`
}

type RecipeReagents struct {
	RecipeID  int     `pg:",pk"`
	Recipe    *Recipe `pg:"rel:has-one"`
	ReagentID int     `pg:",pk"`
	Reagent   *Item   `json:"reagent" pg:"rel:has-one"`
	Quantity  int     `json:"quantity"`
}

type RecipeMedia struct {
	Identifiable
	RecipeID int            ``
	Recipe   *Recipe        `pg:"rel:has-one"`
	Assets   []RecipeAssets `pg:"-"`
}

type RecipeAssets struct {
	RecipeMediaID int          `pg:",pk"`
	RecipeMedia   *RecipeMedia `pg:"rel:has-one"`
	Asset
}
