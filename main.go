package main

import (
	"flag"
	"os"
)

var (
	config       = ReadConfig(os.ExpandEnv(*conf_file))
	message      = flag.String("msg", "", "Write the message you want to send")
	room         = flag.String("room", config.Room, "Which room to send the message to")
	from         = flag.String("from", config.From, "From field")
	token        = flag.String("token", config.Token, "AuthToken API")
	conf_file    = flag.String("config", "$HOME/.hipchatrc", "Config file path")
	print_config = flag.Bool("print", false, "Print config")
)

func main() {
	flag.Parse()

	// pretty print config if -print is on
	if *print_config {
		PrintConfig(config)
		return
	}

	// finally send the message
	SendMessage()
}
