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

		// TODO: Add update, delete
		switch msg.Subject {
		case "command.booking.create":
			var cmd booking.CreateBookingCommand

			err = json.Unmarshal(msg.Data, &cmd)
			if err != nil {
				log.Println("Error:", err)
			}

			log.Println("Received command:", cmd)

			err := h.HandleCreateBookingCommand(cmd)
			if err != nil {
				msg.Respond([]byte(err.Error()))
				log.Println("Error:", err)
			}

			msg.Respond([]byte("OK"))

		case "command.booking.delete":
			id := string(msg.Data)

			log.Println("Received delete command for booking id:", id)

			cmd := booking.DeleteBookingCommand{
				Id: id,
			}

			err := h.HandleDeleteBookingCommand(cmd)
			if err != nil {
				msg.Respond([]byte(err.Error()))
				log.Println("Error:", err)
			}

			msg.Respond([]byte("OK"))
		}
	}

}
