package datasets

type Faction struct {
	ID string `json:"type" pg:",pk"`
	Name LocalizedField `json:"name"`
}
