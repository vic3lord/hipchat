package main

import (
	"flag"
	"os"
)

var (
	message   = flag.String("m", "", "Write the message you want to send")
	room      = flag.String("r", "", "Which room to send the message to")
	from      = flag.String("f", "", "From field")
	token     = flag.String("t", "", "AuthToken API")
	conf_file = flag.String("c", "$HOME/.hipchat.json", "Config file path")
)

func main() {
	flag.Parse()

	config := ReadConfig(os.ExpandEnv(*conf_file))
	if config.From == "" {
		config.From = *from
	}
	if config.Room == "" {
		config.Room = *room
	}
	if config.Token == "" {
		config.Token = *token
	}

	SendMessage(*message, config)
}
