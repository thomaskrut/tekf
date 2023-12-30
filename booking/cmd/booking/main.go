package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/thomaskrut/tekf/booking"
	"github.com/thomaskrut/tekf/booking/client"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sub, err := nc.SubscribeSync("command.booking.*")
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	c := client.New()
	defer c.Close()

	h := booking.NewBookingCommandHandler(nc, c)

	log.Printf("Listening on subject [command.booking.*]\n")
	for {
		msg, err := sub.NextMsg(1 * 1000 * 1000 * 1000)
		if err != nil {
			if err == nats.ErrTimeout {
				continue
			}
			log.Fatal(err)
		}

		switch msg.Subject {
		case "command.booking.create":
			var cmd booking.CreateBookingCommand
			err = json.Unmarshal(msg.Data, &cmd)
			log.Println("Received command:", cmd)
			if err != nil {
				log.Println("Error:", err)
				return
			}
			err := h.HandleCreateBookingCommand(cmd)
			if err != nil {
				log.Println("Error:", err)
				return
			}
		}
	}

}
