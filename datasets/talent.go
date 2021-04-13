package datasets

type Talent struct {
	Identifiable
	PlayableClassID  int            ``
	PlayableClass    *PlayableClass `json:"playable_class" pg:"rel:has-one"`
	SpellID          int            ``
	Spell            *Spell         `json:"spell" pg:"rel:has-one"`
	OverridesSpellID int            ``
	OverridesSpell   *Spell         `json:"overrides_spell" pg:"rel:has-one"`
	TierIndex        int            `json:"tier_index"`
	ColumnIndex      int            `json:"column_index"`
	Level            int            `json:"level"`
	Description      LocalizedField `json:"description"`
}

type PvpTalent struct {
	Identifiable
	PlayableSpecializationID int                     ``
	PlayableSpecialization   *PlayableSpecialization `json:"playable_specialization" pg:"rel:has-one"`
	SpellID                  int                     ``
	Spell                    *Spell                  `json:"spell" pg:"rel:has-one"`
	OverridesSpellID         int                     ``
	OverridesSpell           *Spell                  `json:"overrides_spell" pg:"rel:has-one"`
	Description              LocalizedField          `json:"description"`
	UnlockPlayerLevel        int                     `json:"unlock_player_level"`
	CompatibleSlots          []int                   `json:"compatible_slots"`
}
