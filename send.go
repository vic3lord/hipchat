package main

import (
	"github.com/andybons/hipchat"
	"log"
)

func SendMessage(msg string) {
	c := hipchat.Client{AuthToken: *token}
	request := hipchat.MessageRequest{
		RoomId:        *room,
		From:          *from,
		Message:       msg,
		Color:         hipchat.ColorYellow,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}

	if err := c.PostMessage(request); err != nil {
		log.Printf("Cannot send message: %s %q", msg, err)
	}
}
