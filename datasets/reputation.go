package datasets

type ReputationTierItem struct {
	ReputationTierID int `pg:",pk,use_zero,on_delete:RESTRICT, on_update: CASCADE"`
	ReputationTier   *ReputationTier
	Identifiable
	Name     LocalizedField `json:"name"`
	MinValue int            `json:"min_value"`
	MaxValue int            `json:"max_value"`
}

type ReputationTier struct {
	Identifiable
	Tiers               []ReputationTierItem    `json:"tiers" pg:"-"`
}

type ReputationFaction struct {
	Identifiable
	ParentFactionID   int                 `pg:",on_delete:RESTRICT, on_update: CASCADE"`
	ParentFaction     *ReputationFaction  `json:""`
	Name              LocalizedField      `json:"name"`
	Description       LocalizedField      `json:"description"`
	ReputationTierID  int                 `pg:",use_zero,on_delete:RESTRICT, on_update: CASCADE"`
	ReputationTier    *ReputationTier     `json:"reputation_tiers"`
	IsHeader          bool                `json:"is_header"`
	ReputationFaction []ReputationFaction `json:"factions" pg:"-"`
}