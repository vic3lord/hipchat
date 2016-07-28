package main

import (
	"fmt"
	"time"

	"github.com/andybons/hipchat"
)

const timeout time.Duration = 2

// SendMessage gets config and message params then sends the message to a room
func SendMessage(config Config, msg string) {
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
		fmt.Println("[INFO] Message sent successfully \U0001f604")
	case err := <-errc:
		fmt.Printf("[ERROR] Cannot send message: %s %q", msg, err)
	case <-time.After(time.Second * timeout):
		fmt.Printf("[ERROR] Timout after %d seconds", timeout)
	}
}
