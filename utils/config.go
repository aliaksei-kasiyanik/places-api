package utils

import (
	"encoding/json"
	"os"
	"gopkg.in/mgo.v2"
	"time"
)

type MongoConfiguration struct {
	Addrs   []string `json:"addrs"`
	Timeout int64 `json:"timeout"` // in sec
}

type Configuration struct {
	AppAddr     string      `json:"appAddr"`
	MongoConfig *MongoConfiguration `json:"mongo"`
}

func NewConfiguration(configPath string) *Configuration {
	if file, err := os.Open(configPath); err == nil {
		decoder := json.NewDecoder(file)
		configuration := &Configuration{}
		err := decoder.Decode(configuration)
		if err != nil {
			panic("Invalid configuration file.")
		}
		return configuration
	}
	panic("Can't find config file.")
}

func (c *Configuration) GetMongoDialInfo() *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs: c.MongoConfig.Addrs,
		Timeout: time.Duration(c.MongoConfig.Timeout) * time.Second,
	}
}
