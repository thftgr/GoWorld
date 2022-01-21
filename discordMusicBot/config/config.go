package config

import (
	"encoding/json"
	"os"
)

type ConfigJson struct {
	Disocrd struct {
		Token string `json:"token"`
	} `json:"disocrd"`
	Youtube struct {
		Token string `json:"token"`
	} `json:"youtube"`
}

var Config ConfigJson

func init() {
	b, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &Config)
	if err != nil {
		panic(err)
	}
}
