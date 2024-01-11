package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/thomaskrut/tekf/checkinout"
	"github.com/thomaskrut/tekf/checkinout/client"
)

var (
	natsUrl         = "nats://nats:4222"
	bookingSubject  = "command.booking.*"
	checkinSubject  = "command.booking.checkin"
	checkoutSubject = "command.booking.checkout"
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

	h := checkinout.NewCommandHandler(nc, c)

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
		case checkinSubject:
			id := string(msg.Data)

			log.Printf("Received %s for id: %s", checkinSubject, id)

			err := h.HandleCheckinCommand(id)
			if err != nil {
				msg.Respond([]byte(err.Error()))
				log.Println("Error:", err)
			}

			msg.Respond([]byte("OK"))

		case checkoutSubject:
			id := string(msg.Data)

			log.Printf("Received %s for id: %s", checkoutSubject, id)

			err := h.HandleCheckoutCommand(id)
			if err != nil {
				msg.Respond([]byte(err.Error()))
				log.Println("Error:", err)
			}

			msg.Respond([]byte("OK"))
		}
	}

}
