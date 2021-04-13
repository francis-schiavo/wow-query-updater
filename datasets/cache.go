package datasets

import "time"

type Cache struct {
	Key       string `pg:",notnull,pk"`
	Url       string `pg:",notnull,pk"`
	Status    int
	Payload   string `pg:"type:jsonb"`
	Namespace string
	Schema    string
	Revision  string
	Hash      string
	CachedAt  time.Time
}
