package datasets

type PlayableRace struct {
	Identifiable
	Name         LocalizedField       `json:"name"`
	GenderName   GenderLocalizedField `json:"gender_name"`
	FactionID    string               ``
	Faction      *Faction             `json:"faction" pg:"rel:has-one"`
	IsSelectable bool                 `json:"is_selectable"`
	IsAlliedRace bool                 `json:"is_allied_race"`
}
