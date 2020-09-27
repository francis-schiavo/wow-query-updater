package datasets

type UpdateError struct {
	ID       int
	Endpoint string
	RecordID int
	Error    string
}

type PowerType struct {
	Identifiable
	Name LocalizedField `json:"name"`
}

type PlayableRace struct {
	Identifiable
	Name         LocalizedField       `json:"name"`
	GenderName   GenderLocalizedField `json:"gender_name"`
	FactionID    string               `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	Faction      *Faction             `json:"faction"`
	IsSelectable bool                 `json:"is_selectable"`
	IsAlliedRace bool                 `json:"is_allied_race"`
}

type PlayableClass struct {
	Identifiable
	Name        LocalizedField       `json:"name"`
	GenderName  GenderLocalizedField `json:"gender_name"`
	PowerTypeID int                  `pg:",use_zero,on_delete:RESTRICT,on_update:CASCADE"`
	PowerType   *PowerType           `json:"power_type"`
	Media       *PlayableClassMedia  `pg:"-"`
}

type PlayableClassMedia struct {
	Identifiable
	PlayableClassID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	PlayableClass   *PlayableClass
	Assets          []PlayableClassAssets `pg:"-"`
}

type PlayableClassAssets struct {
	PlayableClassMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	PlayableClassMedia   *PlayableClass
	Asset
}

type Role struct {
	ID   string         `json:"type" pg:",pk"`
	Name LocalizedField `json:"name"`
}

type PlayableSpecialization struct {
	Identifiable
	PlayableClassID   int                          `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	PlayableClass     *PlayableClass               `json:"playable_class"`
	Name              LocalizedField               `json:"name"`
	GenderDescription GenderLocalizedField         `json:"gender_description"`
	RoleID            string                       `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	Role              Role                         `json:"role"`
	Media             *PlayableSpecializationMedia `pg:"-"`
}

type PlayableSpecializationMedia struct {
	Identifiable
	PlayableSpecializationID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	PlayableSpecialization   *PlayableSpecialization
	Assets                   []PlayableSpecializationAsset `pg:"-"`
}

type PlayableSpecializationAsset struct {
	PlayableSpecializationMediaID int `pg:",pk"`
	PlayableSpecializationMedia   *PlayableSpecializationMedia
	Asset
}

type Talent struct {
	Identifiable
	TierIndex        int            `json:"tier_index"`
	ColumnIndex      int            `json:"column_index"`
	Level            int            `json:"level"`
	Description      LocalizedField `json:"description"`
	SpellID          int            `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	Spell            *Spell         `json:"spell"`
	OverridesSpellID int            `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	OverridesSpell   *Spell         `json:"overrides_spell"`
	PlayableClassID  int            `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	PlayableClass    *PlayableClass `json:"playable_class"`
}

type PvpTalent struct {
	Identifiable
	SpellID                  int                     `pg:"on_delete:RESTRICT,on_update: CASCADE"`
	Spell                    *Spell                  `json:"spell"`
	PlayableSpecializationID int                     `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	PlayableSpecialization   *PlayableSpecialization `json:"playable_specialization"`
	OverridesSpellID         int                     `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	OverridesSpell           *Spell                  `json:"overrides_spell"`
	Description              LocalizedField          `json:"description"`
	UnlockPlayerLevel        int                     `json:"unlock_player_level"`
	CompatibleSlots          []int                   `json:"compatible_slots"`
}

type Title struct {
	Identifiable
	Name       LocalizedField       `json:"name"`
	GenderName GenderLocalizedField `json:"gender_name"`
}
