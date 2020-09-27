package datasets

type Source struct {
	ID   string         `json:"type" pg:",pk"`
	Name LocalizedField `json:"name"`
}

type Mount struct {
	Identifiable
	Name                       LocalizedField      `json:"name"`
	Description                LocalizedField      `json:"description"`
	SourceID                   string              `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Source                     *Source             `json:"source"`
	FactionID                  string              `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Faction                    *Faction            `json:"faction"`
	MountDisplays              []MountDisplayMedia `json:"creature_displays"`
	ShouldExcludeIfUncollected bool                `json:"should_exclude_if_uncollected"`
}

type MountDisplayMedia struct {
	Identifiable
	MountID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Mount   *Mount
	Assets  []MountDisplayAssets `pg:"-"`
}

type MountDisplayAssets struct {
	MountDisplayMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	MountDisplayMedia   *MountDisplayMedia
	Asset
}

type BattlePetType struct {
	Identifiable
	Type string         `json:"type"`
	Name LocalizedField `json:"name"`
}

type BattlePetAbility struct {
	Identifiable
	Name LocalizedField `json:"name"`
}

type PetAbility struct {
	PetID         int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Pet           *Pet
	AbilityID     int               `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Ability       *BattlePetAbility `json:"ability"`
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
	BattlePetType           *BattlePetType `json:"battle_pet_type"`
	IsAllianceOnly          bool           `json:"is_alliance_only"`
	IsHordeOnly             bool           `json:"is_horde_only"`
	Abilities               []PetAbility   `json:"abilities" pg:"-"`
	SourceID                string         `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Source                  *Source        `json:"source"`
	Icon                    string         `json:"icon"`
	CreatureID              int            `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Creature                *Creature      `json:"creature"`
	IsRandomCreatureDisplay bool           `json:"is_random_creature_display"`
}