package datasets

type BattlePetType struct {
	Identifiable
	Type string         `json:"type"`
	Name LocalizedField `json:"name"`
}

type BattlePetAbility NamedItem

type PetAbility struct {
	PetID         int               `pg:",pk"`
	Pet           *Pet              `pg:"rel:has-one"`
	AbilityID     int               `pg:",pk"`
	Ability       *BattlePetAbility `json:"ability" pg:"rel:has-one"`
	Slot          int               `json:"slot"`
	RequiredLevel int               `json:"required_level"`
}

type Pet struct {
	Identifiable
	Name                    LocalizedField `json:"name"`
	Description             LocalizedField `json:"description"`
	IsCapturable            bool           `json:"is_capturable"`
	IsTradable              bool           `json:"is_tradable"`
	IsBattlepet             bool           `json:"is_battlepet"`
	BattlePetTypeID         int            `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	BattlePetType           *BattlePetType `json:"battle_pet_type" pg:"rel:has-one"`
	IsAllianceOnly          bool           `json:"is_alliance_only"`
	IsHordeOnly             bool           `json:"is_horde_only"`
	Abilities               []PetAbility   `json:"abilities" pg:"-"`
	SourceID                string         ``
	Source                  *Source        `json:"source" pg:"rel:has-one"`
	Icon                    string         `json:"icon"`
	CreatureID              int            ``
	Creature                *Creature      `json:"creature" pg:"rel:has-one"`
	IsRandomCreatureDisplay bool           `json:"is_random_creature_display"`
}
