package datasets

type CreatureFamily struct {
	Identifiable
	Name                     LocalizedField          `json:"name"`
	PlayableSpecializationID int                     ``
	PlayableSpecialization   *PlayableSpecialization `json:"specialization" pg:"rel:has-one"`
	Media                    *CreatureFamilyMedia    `pg:"-"`
}

type CreatureFamilyMedia struct {
	Identifiable
	CreatureFamilyID int                    ``
	CreatureFamily   *CreatureFamily        `pg:"rel:has-one"`
	Assets           []CreatureFamilyAssets `pg:"-"`
}

type CreatureFamilyAssets struct {
	CreatureFamilyMediaID int                  ``
	CreatureFamilyMedia   *CreatureFamilyMedia `pg:"rel:has-one"`
	Asset
}

type CreatureType NamedItem

type Creature struct {
	Identifiable
	Name       LocalizedField  `json:"name"`
	TypeID     int             ``
	Type       *CreatureType   `json:"type" pg:"rel:has-one"`
	FamilyID   int             ``
	Family     *CreatureFamily `json:"family" pg:"rel:has-one"`
	Media      Identifiables   `json:"creature_displays" pg:"-"`
	IsTameable bool            `json:"is_tameable"`
}

type CreatureDisplayMedia struct {
	Identifiable
	CreatureID int                     ``
	Creature   *Creature               `pg:"rel:has-one"`
	Assets     []CreatureDisplayAssets `pg:"-"`
}

type CreatureDisplayAssets struct {
	CreatureDisplayMediaID int                   ``
	CreatureDisplayMedia   *CreatureDisplayMedia `pg:"rel:has-one"`
	Asset
}
