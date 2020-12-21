package datasets

type TechTalent struct {
	Identifiable
	TalentTreeID         int             ``
	TalentTree           *TechTalentTree `json:"talent_tree" pg:"rel:has-one"`
	Name                 LocalizedField  `json:"name"`
	Description          LocalizedField  `json:"description"`
	SpellTooltip         *SpellTooltip   `json:"spell_tooltip"`
	Tier                 int             `json:"tier" pg:",use_zero"`
	DisplayOrder         int             `json:"display_order" pg:",use_zero"`
	PrerequisiteTalent   *TechTalent     `json:"prerequisite_talent"`
}

type TechTalentTree struct {
	Identifiable
	MaxTiers int           `json:"max_tiers"`
	Talents  Identifiables `json:"talents"`
}

type TechTalentMedia struct {
	Identifiable
	TechTalentID int          ``
	TechTalent *TechTalent    `pg:"rel:has-one"`
	Assets []TechTalentAssets `pg:"-"`
}

type TechTalentAssets struct {
	TechTalentMediaID int              `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	TechTalentMedia   *TechTalentMedia `pg:"rel:has-one"`
	Asset
}
