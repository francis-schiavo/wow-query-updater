package datasets

type PlayableClass struct {
	Identifiable
	Name        LocalizedField       `json:"name"`
	GenderName  GenderLocalizedField `json:"gender_name"`
	PowerTypeID int                  `pg:",use_zero"`
	PowerType   *PowerType           `json:"power_type" pg:"rel:has-one"`
	Media       *PlayableClassMedia  `pg:"-"`
}

type PlayableClassMedia struct {
	Identifiable
	PlayableClassID int                   ``
	PlayableClass   *PlayableClass        `pg:"rel:has-one"`
	Assets          []PlayableClassAssets `pg:"-"`
}

type PlayableClassAssets struct {
	PlayableClassMediaID int            `pg:",pk"`
	PlayableClassMedia   *PlayableClass `pg:"rel:has-one"`
	Asset
}
