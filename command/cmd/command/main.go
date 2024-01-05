package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
)

var natsUrl = "nats://nats:4222"

func main() {

	r := chi.NewRouter()

	r.Post("/booking", CreateBookingHandler)
	r.Delete("/booking/{id}", DeleteBookingHandler)

	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	id := chi.URLParam(r, "id")

	log.Printf("Received id: %s", id)

	response, err := nc.Request("command.booking.delete", []byte(id), time.Second*3)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	log.Println(string(response.Data))

	if string(response.Data) != "OK" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.Data)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received payload: %s", payload)

	response, err := nc.Request("command.booking.create", payload, time.Second*3)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	log.Println(string(response.Data))

	if string(response.Data) != "OK" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.Data)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
