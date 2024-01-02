package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
)

func main() {

	r := chi.NewRouter()

	r.Post("/booking", CreateBookingHandler)

	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received payload: %s", payload)

	if err = nc.Publish("command.booking.create", payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	// Wait for reply?

	w.WriteHeader(http.StatusOK)
}
