package datasets

type JournalCategory struct {
	ID string `json:"type"`
}

type JournalMode struct {
	ID   string         `json:"type"`
	Name LocalizedField `json:"name"`
}

type JournalExpansion struct {
	Identifiable
	Name     LocalizedField  `json:"name"`
	Dungeons []Identifiables `json:"dungeons" pg:"-"`
	Raids    []Identifiables `json:"raids" pg:"-"`
}

type JournalEncounterSection struct {
	Identifiable
	ParentSectionID        int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	ParentSection          *JournalEncounterSection
	Title                  LocalizedField            `json:"title"`
	BodyText               LocalizedField            `json:"body_text"`
	Sections               []JournalEncounterSection `json:"sections" pg:"-"`
	CreatureDisplayMediaID int                       `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	CreatureDisplayMedia   *CreatureDisplayMedia     `json:"creature_display"`
}

type JournalEncounter struct {
	Identifiable
	Name        LocalizedField              `json:"name"`
	Description LocalizedField              `json:"description"`
	Creatures   []JournalEncounterCreatures `json:"creatures" pg:"-"`
	Items       []JournalEncounterItems     `json:"items" pg:"-"`
	Sections    []JournalEncounterSection   `json:"sections" pg:"-"`
	InstanceID  int                         `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Instance    JournalInstance             `json:"instance"`
	CategoryID  string                      `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Category    JournalCategory             `json:"category"`
	Modes       []JournalMode               `json:"modes" pg:"-"`
}

type JournalEncounterModes struct {
	EncounterID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Encounter   *JournalEncounter
	ModeID      string `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Mode        *JournalMode
}

type JournalEncounterCreatures struct {
	EncounterID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Encounter   *JournalEncounter
	CreatureID  int `json:"id" pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
}

type JournalEncounterItems struct {
	Identifiable
	EncounterID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Encounter   *JournalEncounter
	ItemID             int   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Item               *Item `json:"item"`
}

type JournalMap struct {
	Identifiable
	Name LocalizedField `json:"name"`
}

type JournalArea struct {
	Identifiable
	Name LocalizedField `json:"name"`
}

type JournalLocation struct {
	Identifiable
	Name LocalizedField `json:"name"`
}

type JournalInstance struct {
	Identifiable
	Name         LocalizedField        `json:"name"`
	Description  LocalizedField        `json:"description"`
	MinimumLevel int                   `json:"minimum_level"`
	CategoryID   string                `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Category     *JournalCategory       `json:"category"`
	ExpansionID  int                   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Expansion    *JournalExpansion     `json:"expansion"`
	MapID        int                   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Map          *JournalMap            `json:"map"`
	AreaID       int                   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Area         *JournalArea           `json:"area"`
	LocationID   int                   `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Location     *JournalLocation       `json:"location"`
	Encounters   []Identifiables       `json:"encounters" pg:"-"`
	Modes        []JournalInstanceMode `json:"modes" pg:"-"`
	Media        *JournalInstanceMedia  `json:"media" pg:"-"`
}

type JournalInstanceMode struct {
	InstanceID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Instance   *JournalInstance
	ModeID     string       `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	Mode       *JournalMode `json:"mode"`
	Players           int          `json:"players"`
	IsTracked         bool         `json:"is_tracked"`
}

type JournalInstanceMedia struct {
	SelfReference
	Identifiable
	JournalID int `pg:",on_delete:RESTRICT,on_update:CASCADE"`
	Journal   *JournalInstance
	Assets    []JournalInstanceAssets `pg:"-"`
}

type JournalInstanceAssets struct {
	JournalInstanceMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	JournalInstanceMedia   *JournalInstanceMedia
	Asset
}
