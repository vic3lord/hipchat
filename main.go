package main

import (
	"flag"
	"os"
)

var (
	message      = flag.String("msg", "", "Write the message you want to send")
	room         = flag.String("room", "", "Which room to send the message to")
	from         = flag.String("from", "", "From field")
	token        = flag.String("token", "", "AuthToken API")
	conf_file    = flag.String("config", "$HOME/.hipchatrc", "Config file path")
	print_config = flag.Bool("print", false, "Print config")
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
	if *print_config {
		PrintConfig(config)
		return
	}

	SendMessage(*message, config)
}
