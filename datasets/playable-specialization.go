package datasets

type PlayableSpecialization struct {
	Identifiable
	PlayableClassID   int                          ``
	PlayableClass     *PlayableClass               `json:"playable_class" pg:"rel:has-one"`
	Name              LocalizedField               `json:"name"`
	GenderDescription GenderLocalizedField         `json:"gender_description"`
	RoleID            string                       ``
	Role              Role                         `json:"role" pg:"rel:has-one"`
	Media             *PlayableSpecializationMedia `pg:"-"`
}

type PlayableSpecializationMedia struct {
	Identifiable
	PlayableSpecializationID int                           ``
	PlayableSpecialization   *PlayableSpecialization       `pg:"rel:has-one"`
	Assets                   []PlayableSpecializationAsset `pg:"-"`
}

type PlayableSpecializationAsset struct {
	PlayableSpecializationMediaID int                          `pg:",pk"`
	PlayableSpecializationMedia   *PlayableSpecializationMedia `pg:"rel:has-one"`
	Asset
}
