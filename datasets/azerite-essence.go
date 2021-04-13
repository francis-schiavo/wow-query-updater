package datasets

type AzeritePower struct {
	Identifiable
	Rank                int    `json:"rank"`
	MainPowerSpellID    int    ``
	MainPowerSpell      *Spell `json:"main_power_spell" pg:"rel:has-one"`
	PassivePowerSpellID int    ``
	PassivePowerSpell   *Spell `json:"passive_power_spell" pg:"rel:has-one"`
}
type AzeritePowers []AzeritePower

type AzeriteEssence struct {
	Identifiable
	Name                   LocalizedField       `json:"name"`
	AllowedSpecializations Identifiables        `json:"allowed_specializations" pg:"-"`
	Powers                 AzeritePowers        `json:"powers" pg:"-"`
	Media                  *AzeriteEssenceMedia `pg:"-"`
}

type AzeriteEssencePower struct {
	AzeriteEssenceID int             `pg:",pk"`
	AzeriteEssence   *AzeriteEssence `pg:"rel:has-one"`
	AzeritePowerID   int             `pg:",pk"`
	AzeritePower     *AzeritePower   `pg:"rel:has-one"`
}

type AzeriteEssenceSpecializations struct {
	AzeriteEssenceID         int                     `pg:",pk"`
	AzeriteEssence           *AzeriteEssence         `pg:"rel:has-one"`
	PlayableSpecializationID int                     `pg:",pk"`
	PlayableSpecialization   *PlayableSpecialization `pg:"rel:has-one"`
}

type AzeriteEssenceMedia struct {
	SelfReference
	Identifiable
	AzeriteEssenceID int                   ``
	AzeriteEssence   *AzeriteEssence       `pg:"rel:has-one"`
	Assets           []AzeriteEssenceAsset `pg:"-"`
}

type AzeriteEssenceAsset struct {
	AzeriteEssenceMediaID int                  ``
	AzeriteEssenceMedia   *AzeriteEssenceMedia `pg:"rel:has-one"`
	Asset
}
