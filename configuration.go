package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`

	RedisHost string `json:"redis_host"`
	RedisDB   int    `json:"redis_db"`
}

func (c *Config) LoadFromFile(filename string) {
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	parseError := json.Unmarshal(byteValue, c)
	if parseError != nil {
		log.Fatal(err)
	}
}