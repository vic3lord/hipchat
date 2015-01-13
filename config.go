package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Room  string `json:"room",omitempty`
	From  string `json:"from",omitempty`
	Token string `json:"token",omitempty`
}

func ReadConfig(file string) Config {
	conf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("[ERROR] Cannot open config file %s, %q", file, err)
		return Config{}
	}

	var confStruct Config
	if err = json.Unmarshal(conf, &confStruct); err != nil {
		log.Printf("[ERROR] parsing config file %s, %q", file, err)
	}

	return confStruct
}
