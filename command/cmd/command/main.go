package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/nats-io/nats.go"
)

var (
	natsUrl              = "nats://nats:4222"
	createBookingSubject = "command.booking.create"
	deleteBookingSubject = "command.booking.delete"
	updateBookingSubject = "command.booking.update"
)

func main() {

	r := chi.NewRouter()

	corsHandler := cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"*"},
	})

	r.Use(corsHandler)
	r.Post("/booking", CreateBookingHandler)
	r.Delete("/booking/{id}", DeleteBookingHandler)
	r.Put("/booking", UpdateBookingHandler)

	log.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func UpdateBookingHandler(w http.ResponseWriter, r *http.Request) {
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

	response, err := nc.Request(updateBookingSubject, payload, time.Second*3)
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

func DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	id := chi.URLParam(r, "id")

	log.Printf("Received id: %s", id)

	response, err := nc.Request(deleteBookingSubject, []byte(id), time.Second*3)
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

	response, err := nc.Request(createBookingSubject, payload, time.Second*3)
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
