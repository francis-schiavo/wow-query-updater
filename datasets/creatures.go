package datasets

type CreatureFamily struct {
	Identifiable
	Name                     LocalizedField          `json:"name"`
	PlayableSpecializationID int                     `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	PlayableSpecialization   *PlayableSpecialization `json:"specialization"`
	Media                    *CreatureFamilyMedia    `pg:"-"`
}

type CreatureFamilyMedia struct {
	Identifiable
	CreatureFamilyID int                    `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	CreatureFamily   *CreatureFamily
	Assets           []CreatureFamilyAssets `pg:"-"`
}

type CreatureFamilyAssets struct {
	CreatureFamilyMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	CreatureFamilyMedia   *CreatureFamilyMedia
	Asset
}

type CreatureType struct {
	Identifiable
	Name LocalizedField `json:"name"`
}

type Creature struct {
	Identifiable
	Name       LocalizedField        `json:"name"`
	TypeID     int                   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Type       *CreatureType         `json:"type"`
	FamilyID   int                   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Family     *CreatureFamily       `json:"family"`
	Media      []CreatureDisplayMedia `json:"creature_displays" pg:"-"`
	IsTameable bool                  `json:"is_tameable"`
}

type CreatureDisplayMedia struct {
	Identifiable
	CreatureID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Creature   *Creature
	Assets     []CreatureDisplayAssets `pg:"-"`
}

type CreatureDisplayAssets struct {
	CreatureDisplayMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	CreatureDisplayMedia   *CreatureDisplayMedia
	Asset
}