package datasets

type TechTalent struct {
	Identifiable
	TalentTree         *TechTalentTree `json:"talent_tree" pg:"rel:has-one"`
	Name               LocalizedField  `json:"name"`
	Description        LocalizedField  `json:"description"`
	SpellTooltip       *SpellTooltip   `json:"spell_tooltip"`
	Tier               int             `json:"tier"`
	DisplayOrder       int             `json:"display_order"`
	PrerequisiteTalent *TechTalent     `json:"prerequisite_talent" pg:"rel:has-one"`
}

type TechTalentTree struct {
	Identifiable
	MaxTiers int           `json:"max_tiers"`
	Talents  Identifiables `json:"talents"`
}

type TechTalentMedia struct {
	Identifiable
	ItemID int           `pg:",pk"`
	Item   *Item         `pg:"rel:has-one"`
	Assets Identifiables `pg:"-"`
}

type TechTalentAssets struct {
	TechTalentMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Asset
}
