package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	message     = flag.String("msg", "", "Write the message you want to send")
	room        = flag.String("room", "", "Which room to send the message to")
	from        = flag.String("from", "", "From field")
	token       = flag.String("token", "", "AuthToken API")
	confFile    = flag.String("config", os.ExpandEnv("$HOME/.hipchatrc"), "Config file path")
	printConfig = flag.Bool("print", false, "Print config")
)

func main() {
	flag.Parse()

	config, err := NewConfig(*confFile)
	if err != nil {
		fmt.Println(err)
	}

	if *token != "" {
		config.Token = *token
	}

	if *from != "" {
		config.From = *from
	}

	if *room != "" {
		config.Room = *room
	}

	if *printConfig {
		fmt.Println(config)
		return
	}

	if *message != "" {
		SendMessage(config, *message)
	}
}
