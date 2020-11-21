package datasets

type ReputationTierItem struct {
	Identifiable
	ReputationTierID int             `pg:",pk,use_zero"`
	ReputationTier   *ReputationTier `pg:"rel:has-one"`
	Name             LocalizedField  `json:"name"`
	MinValue         int             `json:"min_value"`
	MaxValue         int             `json:"max_value"`
}

type ReputationTier struct {
	Identifiable
	Name  LocalizedField       `json:"name"`
	Tiers []ReputationTierItem `json:"tiers" pg:"-"`
}

type ReputationFaction struct {
	Identifiable
	ParentFactionID   int                 ``
	ParentFaction     *ReputationFaction  `pg:"rel:has-one"`
	ReputationTierID  int                 `pg:",use_zero"`
	ReputationTier    *ReputationTier     `json:"reputation_tiers" pg:"rel:has-one"`
	Name              LocalizedField      `json:"name"`
	Description       LocalizedField      `json:"description"`
	IsHeader          bool                `json:"is_header"`
	ReputationFaction []ReputationFaction `json:"factions" pg:"-"`
}
