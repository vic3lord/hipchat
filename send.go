package main

import (
	"github.com/andybons/hipchat"
	"log"
)

func SendMessage(msg string, config Config) {
	c := hipchat.Client{AuthToken: config.Token}
	request := hipchat.MessageRequest{
		RoomId:        config.Room,
		From:          config.From,
		Message:       msg,
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}

	if err := c.PostMessage(request); err != nil {
		log.Printf("Cannot send message: %s %q", msg, err)
	}
}
