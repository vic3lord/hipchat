package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Room  string `json:"room,omitempty"`
	From  string `json:"from,omitempty"`
	Token string `json:"token,omitempty"`
}

func NewConfig(f string) (Config, error) {
	var config = Config{
		Room:  "404",
		From:  "root",
		Token: "this-is-not-a-real-token",
	}
	file, err := ioutil.ReadFile(f)
	switch {
	case os.IsNotExist(err):
		return config, nil
	case err == nil:
		err = json.Unmarshal(file, &config)
		if err != nil {
			return config, err
		}
	default:
		return config, errors.New("failed")
	}

	return config, nil
}

func (config Config) String() string {
	return fmt.Sprintf("From: %s\nRoom: %s\nToken: %s", config.From, config.Room, config.Token)
}
