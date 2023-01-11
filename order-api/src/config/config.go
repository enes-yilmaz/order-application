package config

import (
	"fmt"
	"time"
)

var ac AppConfig

type AppConfig struct {
	Env                 string
	MongoClientDuration time.Duration
	MongoClientURI      string
	OrderDBName         string
	OrderCollectionName string
	CustomerClient      string
}

var cfgs = map[string]AppConfig{
	"qa": {
		Env:                 "qa",
		MongoClientDuration: 30 * time.Second,
		MongoClientURI:      "mongodb://localhost:27017",
		OrderDBName:         "orders",
		OrderCollectionName: "order",
		CustomerClient:      "http://localhost:4000/customer-api/customers",
	},
	"prod": {
		Env:                 "qa",
		MongoClientDuration: 30 * time.Second,
		MongoClientURI:      "mongodb://localhost:27017",
		OrderDBName:         "customers",
		OrderCollectionName: "order",
		CustomerClient:      "http://localhost:4000/customer-api/customers",
	},
}

func Config(env string) AppConfig {
	cfg, exist := cfgs[env]
	if !exist {
		panic(fmt.Sprintf("config '%s' not found", env))
	}
	return cfg
}

func SetConfig(cfg AppConfig) {
	ac = cfg
}

func GetConfigs() *AppConfig {
	return &ac
}
