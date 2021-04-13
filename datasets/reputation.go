package datasets

type ReputationTier struct {
	ReputationTierID int             `pg:",pk,use_zero"`
	Identifiable
	Name             LocalizedField  `json:"name"`
	MinValue         int             `json:"min_value"`
	MaxValue         int             `json:"max_value"`
}

type ReputationTierGroup struct {
	Identifiable
	Tiers []ReputationTier `json:"tiers" pg:"-"`
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
