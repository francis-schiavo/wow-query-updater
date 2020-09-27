package datasets

type Spell struct {
	Identifiable
	Name        LocalizedField `json:"name"`
	Description LocalizedField `json:"description"`
	Media       *SpellMedia    `pg:"-"`
}

type SpellMedia struct {
	Identifiable
	SpellID int `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	Spell   *Spell
	Assets  []SpellAssets `pg:"-"`
}

type SpellAssets struct {
	SpellMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	SpellMedia   *SpellMedia
	Asset
}
