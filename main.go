package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config is the configuration JSON file
type Config struct {
	Room  string `json:"room,omitempty"`
	From  string `json:"from,omitempty"`
	Token string `json:"token,omitempty"`
}

var defaultConfig = Config{
	Room:  "Roomie room",
	From:  "root",
	Token: "this-is-a-token",
}

var (
	message     = flag.String("msg", "", "Write the message you want to send")
	room        = flag.String("room", defaultConfig.Room, "Which room to send the message to")
	from        = flag.String("from", defaultConfig.From, "From field")
	token       = flag.String("token", defaultConfig.Token, "AuthToken API")
	confFile    = flag.String("config", "$HOME/.hipchatrc", "Config file path")
	printConfig = flag.Bool("print", false, "Print config")
)

func (config Config) String() string {
	return fmt.Sprintf("From: %s\nRoom: %s\nToken: %s", config.From, config.Room, config.Token)
}

func main() {
	flag.Parse()

	// Read configuration file
	var config Config
	conf, err := ioutil.ReadFile(*confFile)

	switch {
	case os.IsNotExist(err):
		log.Printf("Configuration file %s does not exist, using defaults\n", *confFile)
		config = defaultConfig
	case err == nil:
		if err = json.Unmarshal(conf, &config); err != nil {
			log.Printf("[ERROR] cannot parse config file %s, %q\n", *confFile, err)
		}
	default:
		log.Printf("[ERROR] %q\n", err)
	}

	if *printConfig {
		fmt.Println(config)
		return
	}

	SendMessage(*token, *room, *from, *message)
}
