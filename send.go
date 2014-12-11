package main

import (
	"log"
	"time"

	"github.com/andybons/hipchat"
)

const (
	timeout time.Duration = 2
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

	errc := make(chan error, 1)
	okc := make(chan bool, 1)
	go func() {
		err := c.PostMessage(request)
		if err != nil {
			errc <- err
		} else {
			okc <- true
		}
	}()

	select {
	case <-okc:
	case err := <-errc:
		log.Printf("Cannot send message: %s %q", msg, err)
	case <-time.After(time.Second * timeout):
		log.Printf("Timout after %d seconds", timeout)
	}
}
