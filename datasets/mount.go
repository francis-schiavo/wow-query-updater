package datasets

type Mount struct {
	Identifiable
	Name                       LocalizedField      `json:"name"`
	Description                LocalizedField      `json:"description"`
	SourceID                   string              ``
	Source                     *Source             `json:"source" pg:"rel:has-one"`
	FactionID                  string              ``
	Faction                    *Faction            `json:"faction" pg:"rel:has-one"`
	MountDisplays              []*MountDisplayMedia `json:"creature_displays" pg:"-"`
	ShouldExcludeIfUncollected bool                `json:"should_exclude_if_uncollected"`
}

type MountDisplayMedia struct {
	Identifiable
	MountID int                  ``
	Mount   *Mount               `pg:"rel:has-one"`
	Assets  []MountDisplayAssets `pg:"-"`
}

type MountDisplayAssets struct {
	MountDisplayMediaID int                `pg:",pk"`
	MountDisplayMedia   *MountDisplayMedia `pg:"rel:has-one"`
	Asset
}
