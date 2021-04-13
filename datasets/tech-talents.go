package datasets

type TechTalent struct {
	Identifiable
	TechTalentTreeID     int
	TechTalentTree       *TechTalentTree `json:"talent_tree" pg:"rel:has-one"`
	Name                 *LocalizedField  `json:"name"`
	Description          *LocalizedField  `json:"description"`
	SpellTooltipID       int
	SpellTooltip         *SpellTooltip `json:"spell_tooltip" pg:"rel:has-one"`
	Tier                 int           `json:"tier" pg:",use_zero"`
	DisplayOrder         int           `json:"display_order" pg:",use_zero"`
	PrerequisiteTalentID int
	PrerequisiteTalent   *TechTalent `json:"prerequisite_talent" pg:"rel:has-one"`
}

type TechTalentTree struct {
	Identifiable
	MaxTiers int           `json:"max_tiers"`
	Talents  Identifiables `json:"talents" pg:"-"`
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
