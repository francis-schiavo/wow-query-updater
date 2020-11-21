package datasets

type UpdateError struct {
	ID       int
	Endpoint string
	RecordID int
	Error    string
}

type Faction Enum
type Role Enum
type Source Enum
type Operator Enum
type Currency NamedItem