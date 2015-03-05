package main

import (
	"log"
	"time"

	"github.com/vic3lord/hipchat-msg/Godeps/_workspace/src/github.com/andybons/hipchat"
)

const (
	timeout time.Duration = 2
)

func SendMessage() {
	c := hipchat.Client{AuthToken: *token}
	request := hipchat.MessageRequest{
		RoomId:        *room,
		From:          *from,
		Message:       *message,
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
		log.Println("[INFO] Message sent successfully \U0001f604")
	case err := <-errc:
		log.Printf("[ERROR] Cannot send message: %s %q", *message, err)
	case <-time.After(time.Second * timeout):
		log.Printf("[ERROR] Timout after %d seconds", timeout)
	}
}
