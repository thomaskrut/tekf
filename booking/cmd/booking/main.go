package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/thomaskrut/tekf/booking"
	"github.com/thomaskrut/tekf/booking/client"
)

var (
	natsUrl              = "nats://nats:4222"
	bookingSubject       = "command.booking.*"
	createBookingSubject = "command.booking.create"
	deleteBookingSubject = "command.booking.delete"
	updateBookingSubject = "command.booking.update"
)

func main() {

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sub, err := nc.SubscribeSync(bookingSubject)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	c := client.New()

	h := booking.NewBookingCommandHandler(nc, c)

	log.Printf("Listening on subject %s\n", bookingSubject)
	for {
		msg, err := sub.NextMsg(1 * 1000 * 1000 * 1000)
		if err != nil {
			if err == nats.ErrTimeout {
				continue
			}
			log.Fatal(err)
		}

		if h.LoadState(); err != nil {
			log.Fatal(err)
		}

		switch msg.Subject {
		case createBookingSubject:
			var cmd booking.CreateBookingCommand

			err = json.Unmarshal(msg.Data, &cmd)
			if err != nil {
				log.Println("Error:", err)
			}

			log.Printf("Received %s", createBookingSubject)

			err := h.HandleCreateBookingCommand(cmd)
			if err != nil {
				msg.Respond([]byte(err.Error()))
				log.Println("Error:", err)
			}

			msg.Respond([]byte("OK"))

		case deleteBookingSubject:
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

		case updateBookingSubject:
			var cmd booking.UpdateBookingCommand

			err = json.Unmarshal(msg.Data, &cmd)
			if err != nil {
				log.Println("Error:", err)
			}

			log.Printf("Received %s", updateBookingSubject)

			err := h.HandleUpdateBookingCommand(cmd)
			if err != nil {
				msg.Respond([]byte(err.Error()))
				log.Println("Error:", err)
			}

			msg.Respond([]byte("OK"))
		}
	}

}
