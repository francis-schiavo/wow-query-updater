package datasets

type Spell struct {
	Identifiable
	Name        LocalizedField `json:"name"`
	Description LocalizedField `json:"description"`
	Media       *SpellMedia    `pg:"-"`
}

type SpellTooltip struct {
	Identifiable
	Spell       *Spell         `json:"spell" pg:"-"`
	Description LocalizedField `json:"description"`
	CastTime    LocalizedField `json:"cast_time"`
	Cooldown    LocalizedField `json:"cooldown"`
	Range       LocalizedField `json:"range"`
	PowerCost   LocalizedField `json:"power_cost"`
}

type SpellMedia struct {
	Identifiable
	SpellID int           ``
	Spell   *Spell        `pg:"rel:has-one"`
	Assets  []SpellAssets `pg:"-"`
}

type SpellAssets struct {
	SpellMediaID int         `pg:",pk"`
	SpellMedia   *SpellMedia `pg:"rel:has-one"`
	Asset
}
