package main

import (
	"flag"
)

var (
	message = flag.String("m", "", "Write the message you want to send")
	room    = flag.String("r", "", "Which room to send the message to")
	from    = flag.String("f", "", "From field")
	token   = flag.String("t", "", "AuthToken API")
)

func main() {
	flag.Parse()
	SendMessage(*message)
}
