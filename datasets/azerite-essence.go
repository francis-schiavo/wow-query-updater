package datasets

type AzeriteEssence struct {
	Identifiable
	Name                   LocalizedField       `json:"name"`
	AllowedSpecializations Identifiables        `json:"allowed_specializations" pg:"-"`
	Powers                 AzeritePowers        `json:"powers" pg:"-"`
	Media                  *AzeriteEssenceMedia `pg:"-"`
}

type AzeritePower struct {
	Identifiable
	Rank                int    `json:"rank"`
	MainPowerSpellID    int    `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	MainPowerSpell      *Spell `json:"main_power_spell"`
	PassivePowerSpellID int    `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	PassivePowerSpell   *Spell `json:"passive_power_spell"`
}
type AzeritePowers []AzeritePower

type AzeriteEssencePower struct {
	AzeriteEssenceID int `pg:",pk"`
	AzeriteEssence   *AzeriteEssence
	AzeritePowerID   int `pg:",pk"`
	AzeritePower     *AzeritePower
}

type AzeriteEssenceSpecializations struct {
	AzeriteEssenceID         int `pg:",pk"`
	AzeriteEssence           *AzeriteEssence
	PlayableSpecializationID int `pg:",pk"`
	PlayableSpecialization   *PlayableSpecialization
}

type AzeriteEssenceMedia struct {
	SelfReference
	Identifiable
	AzeriteEssenceID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	AzeriteEssence   *AzeriteEssence
	Assets           []AzeriteEssenceAsset `pg:"-"`
}

type AzeriteEssenceAsset struct {
	AzeriteEssenceMediaID int `pg:",pk"`
	AzeriteEssenceMedia   *AzeriteEssenceMedia
	Asset
}
