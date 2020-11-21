package datasets

type JournalCategory struct {
	ID string `json:"type"`
}

type JournalMode Enum

type JournalExpansion struct {
	Identifiable
	Name     LocalizedField `json:"name"`
	Dungeons Identifiables  `json:"dungeons" pg:"-"`
	Raids    Identifiables  `json:"raids" pg:"-"`
}

type JournalEncounterSection struct {
	Identifiable
	ParentSectionID        int                       ``
	ParentSection          *JournalEncounterSection  `pg:"rel:has-one"`
	Title                  LocalizedField            `json:"title"`
	BodyText               LocalizedField            `json:"body_text"`
	Sections               []JournalEncounterSection `json:"sections" pg:"-"`
	CreatureDisplayMediaID int                       ``
	CreatureDisplayMedia   *CreatureDisplayMedia     `json:"creature_display" pg:"rel:has-one"`
}

type JournalEncounter struct {
	Identifiable
	Name        LocalizedField   `json:"name"`
	Description LocalizedField   `json:"description"`
	Creatures   Identifiables    `json:"creatures" pg:"-"`
	Items       Identifiables    `json:"items" pg:"-"`
	Sections    Identifiables    `json:"sections" pg:"-"`
	InstanceID  int              ``
	Instance    *JournalInstance `json:"instance" pg:"rel:has-one"`
	CategoryID  string           ``
	Category    *JournalCategory `json:"category" pg:"rel:has-one"`
	Modes       []JournalMode    `json:"modes" pg:"-"`
}

type JournalEncounterModes struct {
	EncounterID int               `pg:",pk"`
	Encounter   *JournalEncounter `pg:"rel:has-one"`
	ModeID      string            `pg:",pk"`
	Mode        *JournalMode      `pg:"rel:has-one"`
}

type JournalEncounterCreatures struct {
	EncounterID int               `pg:",pk"`
	Encounter   *JournalEncounter `pg:"rel:has-one"`
	CreatureID  int               `json:"id" pg:",pk"`
	Creature    *Creature         `pg:"rel:has-one"`
}

type JournalEncounterItems struct {
	Identifiable
	EncounterID int               ``
	Encounter   *JournalEncounter `pg:"rel:has-one"`
	ItemID      int               ``
	Item        *Item             `json:"item" pg:"rel:has-one"`
}

type JournalMap NamedItem
type JournalArea NamedItem
type JournalLocation NamedItem

type JournalInstance struct {
	Identifiable
	Name         LocalizedField        `json:"name"`
	Description  LocalizedField        `json:"description"`
	MinimumLevel int                   `json:"minimum_level"`
	CategoryID   string                ``
	Category     *JournalCategory      `json:"category" pg:"rel:has-one"`
	ExpansionID  int                   ``
	Expansion    *JournalExpansion     `json:"expansion" pg:"rel:has-one"`
	MapID        int                   ``
	Map          *JournalMap           `json:"map" pg:"rel:has-one"`
	AreaID       int                   ``
	Area         *JournalArea          `json:"area" pg:"rel:has-one"`
	LocationID   int                   ``
	Location     *JournalLocation      `json:"location" pg:"rel:has-one"`
	Encounters   Identifiables         `json:"encounters" pg:"-"`
	Modes        []JournalInstanceMode `json:"modes" pg:"-"`
	Media        *JournalInstanceMedia `json:"media" pg:"-"`
}

type JournalInstanceMode struct {
	InstanceID int              `pg:",pk"`
	Instance   *JournalInstance `pg:"rel:has-one"`
	ModeID     string           `pg:",pk"`
	Mode       *JournalMode     `json:"mode" pg:"rel:has-one"`
	Players    int              `json:"players"`
	IsTracked  bool             `json:"is_tracked"`
}

type JournalInstanceMedia struct {
	SelfReference
	Identifiable
	JournalID int                     ``
	Journal   *JournalInstance        `pg:"rel:has-one"`
	Assets    []JournalInstanceAssets `pg:"-"`
}

type JournalInstanceAssets struct {
	JournalInstanceMediaID int                   `pg:",pk"`
	JournalInstanceMedia   *JournalInstanceMedia `pg:"rel:has-one"`
	Asset
}
