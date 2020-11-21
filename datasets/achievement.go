package datasets

type AchievementCategory struct {
	Identifiable
	Name                LocalizedField `json:"name" pg:",notnull"`
	IsGuildCategory     bool           `json:"is_guild_category"`
	DisplayOrder        int            `json:"display_order"`
	ParentCategoryID    int
	ParentCategory      *AchievementCategory `json:"parent_category" pg:"rel:has-one"`
	HordeQuantity       int
	HordePoints         int
	AllianceQuantity    int
	AlliancePoints      int
	AggregatesByFaction *struct {
		Alliance struct {
			Quantity int `json:"quantity"`
			Points   int `json:"points"`
		} `json:"alliance"`
		Horde struct {
			Quantity int `json:"quantity"`
			Points   int `json:"points"`
		} `json:"horde"`
	} `json:"aggregates_by_faction" pg:"-"`
	Subcategories Identifiables `json:"subcategories" pg:"-"`
}

type Achievement struct {
	Identifiable
	CategoryID                int
	Category                  *AchievementCategory `json:"category" pg:"rel:has-one"`
	Name                      LocalizedField       `json:"name"`
	Description               LocalizedField       `json:"description"`
	Points                    int                  `json:"points"`
	IsAccountWide             bool                 `json:"is_account_wide"`
	FactionID                 string               ``
	Faction                   *Faction             `json:"faction" pg:"rel:has-one"`
	Criteria                  Identifiable         `json:"criteria"`
	PrerequisiteAchievementID int                  ``
	PrerequisiteAchievement   *Achievement         `json:"prerequisite_achievement" pg:"rel:has-one"`
	NextAchievementID         int                  ``
	NextAchievement           *Achievement         `json:"next_achievement" pg:"rel:has-one"`
	Media                     *AchievementMedia    `json:"media" pg:"-"`
	DisplayOrder              int                  `json:"display_order"`
	RewardDescription         LocalizedField       `json:"reward_description"`
	RewardItemID              int                  ``
	RewardItem                *Item                `json:"reward_item" pg:"rel:has-one"`
	GuildRewardItems          []*Item              `json:"guild_reward_items" pg:"-"`
}

type GuildRewardItems struct {
	AchievementID int          `pg:",pk"`
	Achievement   *Achievement `pg:"rel:has-one"`
	ItemID        int          `pg:",pk"`
	Item          *Item        `pg:"rel:has-one"`
}

type Criteria struct {
	Identifiable
	Description      LocalizedField `json:"description"`
	Amount           int            `json:"amount"`
	ShowProgressBar  bool           `json:"show_progress_bar"`
	IsGold           bool           `json:"is_gold"`
	OperatorID       string         ``
	Operator         *Operator      `json:"operator" pg:"rel:has-one"`
	ChildCriteria    []*Criteria    `json:"child_criteria" pg:"-"`
	FactionID        string         ``
	Faction          *Faction       `json:"faction" pg:"rel:has-one"`
	AchievementID    int            ``
	Achievement      *Achievement   `json:"achievement" pg:"rel:has-one"`
	ParentCriteriaID int            ``
	ParentCriteria   *Criteria      `pg:"rel:has-one"`
}

type AchievementCriteria struct {
	AchievementID int          `pg:",pk"`
	Achievement   *Achievement `pg:"rel:has-one"`
	CriteriaID    int          `pg:",pk"`
	Criteria      *Criteria    `pg:"rel:has-one"`
}

type AchievementAssets struct {
	AchievementMediaID int               `pg:",pk"`
	AchievementMedia   *AchievementMedia `pg:"rel:has-one"`
	Asset
}

type AchievementMedia struct {
	Identifiable
	AchievementID int                 ``
	Achievement   *Achievement        `pg:"rel:has-one"`
	Assets        []AchievementAssets `json:"assets" pg:"-"`
}
