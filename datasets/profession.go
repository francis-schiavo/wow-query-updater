package datasets

type ProfessionType struct {
	ID   string         `json:"type" pg:",pk"`
	Name LocalizedField `json:"name"`
}

type Profession struct {
	Identifiable
	Name        LocalizedField   `json:"name"`
	Description LocalizedField   `json:"description"`
	TypeID      string           `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Type        *ProfessionType  `json:"type"`
	SkillTiers  Identifiables    `json:"skill_tiers"`
	Media       *ProfessionMedia `json:"media" pg:"-"`
}

type ProfessionTier struct {
	Identifiable
	Name              LocalizedField       `json:"name"`
	MinimumSkillLevel int                  `json:"minimum_skill_level"`
	MaximumSkillLevel int                  `json:"maximum_skill_level"`
	Categories        []ProfessionCategory `json:"categories"`
}

type ProfessionCategory struct {
	Identifiable
	TierID         int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Tier *ProfessionTier
	Name           LocalizedField `json:"name"`
	Recipes        []Recipe       `json:"recipes" pg:"-"`
}

type ProfessionTierRecipes struct {
	ProfessionTierID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	ProfessionTier   *ProfessionTier
	RecipeID         int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Recipe           *Recipe
	CategoryID       int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Category         *ProfessionCategory
}

type ProfessionMedia struct {
	Identifiable
	ProfessionID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Profession   *Profession
	Assets       []ProfessionAssets `pg:"-"`
}

type ProfessionAssets struct {
	ProfessionMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	ProfessionMedia   *ProfessionMedia
	Asset
}

type Recipe struct {
	Identifiable
	Name          LocalizedField `json:"name"`
	CraftedItemID int
	CraftedItem   *Item            `json:"crafted_item"`
	Reagents      []RecipeReagents `json:"reagents" pg:"-"`
	Media         *RecipeMedia     `json:"media" pg:"-"`
}

type RecipeReagents struct {
	RecipeID  int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Recipe    *Recipe
	ReagentID int   `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Reagent   *Item `json:"reagent"`
	Quantity  int   `json:"quantity"`
}

type RecipeMedia struct {
	Identifiable
	RecipeID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Recipe   *Recipe
	Assets   []RecipeAssets `pg:"-"`
}

type RecipeAssets struct {
	RecipeMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	RecipeMedia   *RecipeMedia
	Asset
}
