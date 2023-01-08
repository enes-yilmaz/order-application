package config

import (
	"fmt"
	"time"
)

var ac AppConfig

type AppConfig struct {
	Env                    string
	MongoClientDuration    time.Duration
	MongoClientURI         string
	CustomerDBName         string
	CustomerCollectionName string
}

var cfgs = map[string]AppConfig{
	"qa": {
		Env:                    "qa",
		MongoClientDuration:    30 * time.Second,
		MongoClientURI:         "mongodb://localhost:27017",
		CustomerDBName:         "customers",
		CustomerCollectionName: "customer",
	},
	"prod": {
		Env:                    "qa",
		MongoClientDuration:    30 * time.Second,
		MongoClientURI:         "mongodb://localhost:27017",
		CustomerDBName:         "customers",
		CustomerCollectionName: "customer",
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
