package datasets

type QuestType struct {
	Identifiable
	Type   LocalizedField `json:"type"`
	Quests Identifiables  `json:"quests"`
}

type QuestArea struct {
	Identifiable
	Area   LocalizedField `json:"area"`
	Quests Identifiables  `json:"quests"`
}

type QuestCategory struct {
	Identifiable
	Category LocalizedField `json:"category"`
	Quests   Identifiables  `json:"quests"`
}

type Quest struct {
	Identifiable
	Title                   LocalizedField    `json:"title"`
	CategoryID              int               ``
	Category                *QuestCategory    `json:"category" pg:"rel:has-one"`
	AreaID                  int               ``
	Area                    *QuestArea        `json:"area" pg:"rel:has-one"`
	TypeID                  int               ``
	Type                    *QuestType        `json:"type" pg:"rel:has-one"`
	Description             LocalizedField    `json:"description"`
	RecommendedMinimumLevel int               `json:"recommended_minimum_level"`
	RecommendedMaximumLevel int               `json:"recommended_maximum_level"`
	RecommendedLevel        int               `json:"recommended_level"`
	IsDaily                 bool              `json:"is_daily"`
	IsWeekly                bool              `json:"is_weekly"`
	IsRepeatable            bool              `json:"is_repeatable"`
	SpellRewardID           int               ``
	SpellReward             *Spell            `pg:"rel:has-one"`
	TradeSkillSpellID       int               ``
	TradeSkillSpell         *Spell            `pg:"rel:has-one"`
	Requirements            *QuestRequirement `json:"requirements" pg:"-"`
	Rewards                 *QuestReward      `json:"rewards" pg:"-"`
}

type QuestReputationReward struct {
	QuestID  int                `pg:",pk"`
	Quest    *Quest             `pg:"rel:has-one"`
	RewardID int                `pg:",pk"`
	Reward   *ReputationFaction `json:"reward" pg:"rel:has-one"`
	Value    int                `json:"value"`
}

type QuestFollowerReward struct {
	QuestID    int       `pg:",pk"`
	Quest      *Quest    `pg:"rel:has-one"`
	FactionID  string    `pg:",pk"`
	Faction    *Faction  `pg:"rel:has-one"`
	FollowerID int       ``
	Follower   *Follower `pg:"rel:has-one"`
}

type QuestCurrencyReward struct {
	QuestID  int       `pg:",pk"`
	Quest    *Quest    `pg:"rel:has-one"`
	RewardID int       `pg:",pk"`
	Reward   *Currency `json:"reward" pg:"rel:has-one"`
	Value    int       `json:"value"`
}

type QuestItemReward struct {
	QuestID int    `pg:",pk"`
	Quest   *Quest `pg:"rel:has-one"`
	ItemID  int    `pg:",pk"`
	Item    *Item  `json:"reward" pg:"rel:has-one"`
}

type QuestItemChoiceReward struct {
	QuestID                  int                     `pg:",pk"`
	Quest                    *Quest                  `pg:"rel:has-one"`
	PlayableSpecializationID int                     `pg:",pk"`
	PlayableSpecialization   *PlayableSpecialization `pg:"rel:has-one"`
	ItemID                   int                     `pg:",pk"`
	Item                     *Item                   `json:"reward" pg:"rel:has-one"`
}

type QuestRewardSkill NamedItem

type QuestReward struct {
	QuestID           int                      `pg:",pk"`
	Quest             *Quest                   `pg:"rel:has-one"`
	Experience        int                      `json:"experience"`
	Honor             int                      `json:"honor"`
	ArtifactPower     int                      `json:"artifact_power"`
	SpellID           int                      ``
	Spell             *Spell                   `json:"spell" pg:"rel:has-one"`
	TradeSkillSpellID int                      ``
	TradeSkillSpell   *Spell                   `json:"trade_skill_spell" pg:"rel:has-one"`
	TitleID           int                      ``
	Title             *Title                   `json:"title" pg:"rel:has-one"`
	Reputations       []*QuestReputationReward `json:"reputations" pg:"-"`
	Money             int
	MoneyObj          struct {
		Value int `json:"value"`
		Units struct {
			Gold   int `json:"gold"`
			Silver int `json:"silver"`
			Copper int `json:"copper"`
		} `json:"units"`
	} `json:"money" pg:"-"`
	Items struct {
		Items []struct {
			Item *QuestItemReward `json:"item"`
		} `json:"items"`
		ChoiceOf []struct {
			Item         QuestItemChoiceReward `json:"item"`
			Requirements struct {
				PlayableSpecializations Identifiables `json:"playable_specializations"`
			} `json:"requirements"`
		} `json:"choice_of"`
	} `json:"items" pg:"-"`
	Follower struct {
		Alliance Identifiable `json:"alliance"`
		Horde    Identifiable `json:"horde"`
	} `json:"follower" pg:"-"`
	Currency []*QuestCurrencyReward `json:"currency" pg:"-"`
	Skill    struct {
		Reward struct {
			Name LocalizedField `json:"name"`
			ID   int            `json:"id"`
		} `json:"reward"`
		Value int `json:"value"`
	} `json:"skill"`
}

type QuestClassRequirement struct {
	QuestID         int            `pg:",pk"`
	Quest           *Quest         `pg:"rel:has-one"`
	PlayableClassID int            `json:"id" pg:",pk"`
	PlayableClass   *PlayableClass `pg:"rel:has-one"`
}

type QuestRaceRequirement struct {
	QuestID        int           `pg:",pk"`
	Quest          *Quest        `pg:"rel:has-one"`
	PlayableRaceID int           `json:"id" pg:",pk"`
	PlayableRace   *PlayableRace `pg:"rel:has-one"`
}

type QuestReputationRequirement struct {
	QuestID             int                `pg:",pk"`
	Quest               *Quest             `pg:"rel:has-one"`
	ReputationFactionID int                `pg:",pk"`
	ReputationFaction   *ReputationFaction `json:"faction" pg:"rel:has-one"`
	MinReputation       int                `json:"min_reputation"`
	MaxReputation       int                `json:"max_reputation"`
}

type QuestRequirement struct {
	Identifiable
	MinCharacterLevel int                           `json:"min_character_level"`
	MaxCharacterLevel int                           `json:"max_character_level"`
	FactionID         string                        ``
	Faction           *Faction                      `json:"faction" pg:"rel:has-one"`
	Classes           []*QuestClassRequirement      `json:"classes" pg:"-"`
	Reputations       []*QuestReputationRequirement `json:"reputations" pg:"-"`
	Races             []*QuestRaceRequirement       `json:"races"  pg:"-"`
}

// TODO: Previous quest requirements
