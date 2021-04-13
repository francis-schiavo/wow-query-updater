package connections

import (
	"crypto/md5"
	"encoding/hex"
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"log"
	"sync/atomic"
	"time"
	"wow-query-updater/datasets"
)

type PostgresCache struct {
	Key              string
	CachedRequests   int32
	UncachedRequests int32
}

func NewPostgresCache(key string) *PostgresCache {
	return &PostgresCache{
		Key:         key,
	}
}

func (cache *PostgresCache) SaveToCache(identifier string, data *blizzard_api.ApiResponse) {
	hash := md5.Sum(data.Body)
	cacheData := &datasets.Cache{
		Key:       cache.Key,
		Url:       identifier,
		Status:    data.Status,
		Payload:   string(data.Body),
		Namespace: data.BattleNetHeaders.Namespace,
		Schema:    data.BattleNetHeaders.Schema,
		Revision:  data.BattleNetHeaders.Revision,
		Hash:      hex.EncodeToString(hash[:]),
		CachedAt:  time.Now(),
	}

	_, err := dbConn.Model(cacheData).OnConflict("(key,url) DO UPDATE").Insert()
	if err != nil {
		log.Fatalf("Failed to save cache for %s with error:\n %v\n", identifier, err)
	}
}

func (cache *PostgresCache) LoadFromCache(identifier string) (bool, *blizzard_api.ApiResponse) {
	cacheObj := &datasets.Cache{}
	count, err := dbConn.Model(cacheObj).Where("key = ? AND url = ? AND status = 200", cache.Key, identifier).SelectAndCount()
	if err != nil || count != 1 {
		atomic.AddInt32(&cache.UncachedRequests, 1)
		return false, nil
	}

	atomic.AddInt32(&cache.CachedRequests, 1)
	return true, &blizzard_api.ApiResponse{
		Status: cacheObj.Status,
		BattleNetHeaders: &blizzard_api.BattleNetHeaders{
			Namespace: cacheObj.Namespace,
			Schema:    cacheObj.Schema,
			Revision:  cacheObj.Revision,
		},
		Cached: true,
		Body:   []byte(cacheObj.Payload),
		Error:  nil,
	}
}
