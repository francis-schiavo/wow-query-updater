package datasets

type AchievementCategory struct {
	Identifiable
	Name                LocalizedField       `json:"name" pg:",notnull"`
	IsGuildCategory     bool                 `json:"is_guild_category"`
	DisplayOrder        int                  `json:"display_order"`
	ParentCategoryID    int                  `pg:"on_delete:RESTRICT,on_update:CASCADE"`
	ParentCategory      *AchievementCategory `json:"parent_category"`
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

type CriteriaOperator struct {
	ID   string         `json:"type" pg:",pk"`
	Name LocalizedField `json:"name"`
}

type Achievement struct {
	Identifiable
	CategoryID                int                  `pg:",on_delete:RESTRICT, on_update: CASCADE"`
	Category                  *AchievementCategory `json:"category"`
	Name                      LocalizedField       `json:"name"`
	Description               LocalizedField       `json:"description"`
	Points                    int                  `json:"points"`
	IsAccountWide             bool                 `json:"is_account_wide"`
	FactionID                 string               `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	Faction                   *Faction             `json:"faction"`
	Criteria                  *Criteria            `json:"criteria" pg:",many2many:achievement_criterias"`
	PrerequisiteAchievementID int                  `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	PrerequisiteAchievement   *Achievement         `json:"prerequisite_achievement"`
	NextAchievementID         int                  `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	NextAchievement           *Achievement         `json:"next_achievement"`
	Media                     *AchievementMedia    `json:"media" pg:"-"`
	DisplayOrder              int                  `json:"display_order"`
	RewardDescription         LocalizedField       `json:"reward_description"`
	RewardItemID              int                  `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	RewardItem                *Item                `json:"reward_item"`
	GuildRewardItems          []*Item              `json:"guild_reward_items" pg:"-"`
}

type GuildRewardItems struct {
	AchievementID int `pg:",pk,on_delete:RESTRICT, on_update: CASCADE"`
	Achievement   *Achievement
	ItemID        int `pg:",pk,on_delete:RESTRICT, on_update: CASCADE"`
	Item          *Item
}

type Criteria struct {
	Identifiable
	Description      LocalizedField    `json:"description"`
	Amount           int               `json:"amount"`
	ShowProgressBar  bool              `json:"show_progress_bar"`
	IsGold           bool              `json:"is_gold"`
	OperatorID       string            `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	Operator         *CriteriaOperator `json:"operator"`
	ChildCriteria    []*Criteria       `json:"child_criteria" pg:"-"`
	FactionID        string            `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	Faction          *Faction          `json:"faction"`
	AchievementID    int               `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	Achievement      *Achievement      `json:"achievement"`
	ParentCriteriaID int               `pg:"on_delete:RESTRICT, on_update: CASCADE"`
	ParentCriteria   *Criteria
}

type AchievementCriteria struct {
	AchievementID int `pg:",pk"`
	Achievement   *Achievement
	CriteriaID    int `pg:",pk"`
	Criteria      *Criteria
}

type AchievementAssets struct {
	AchievementMediaID int `pg:",pk,on_delete:RESTRICT,on_update:CASCADE"`
	AchievementMedia   *AchievementMedia
	Asset
}

type AchievementMedia struct {
	Identifiable
	AchievementID int `pg:",on_delete:RESTRICT,on_update: CASCADE"`
	Achievement   *Achievement
	Assets        []AchievementAssets `json:"assets" pg:"-"`
}