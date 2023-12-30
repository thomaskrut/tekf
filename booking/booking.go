package booking

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
)

type CreateBookingCommand struct {
	UnitId int    `json:"unitId"`
	From   string `json:"from"`
	To     string `json:"to"`
	Guests int    `json:"guests"`
	Name   string `json:"name"`
}

type CreateBookingEvent struct {
	EventType string    `json:"eventType"`
	Id        string    `json:"id"`
	UnitId    int       `json:"unitId"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Guests    int       `json:"guests"`
	Name      string    `json:"name"`
}

type publisher interface {
	Publish(subject string, data []byte) error
}

type BookingService struct {
	Publisher publisher
}

func NewBookingService(publisher publisher) *BookingService {
	return &BookingService{
		Publisher: publisher,
	}
}

func (s *BookingService) HandleCreateBookingCommand(cmd CreateBookingCommand) error {

	fromTime, err := time.Parse(time.RFC3339, cmd.From)
	if err != nil {
		return err
	}

	toTime, err := time.Parse(time.RFC3339, cmd.To)
	if err != nil {
		return err
	}

	event := CreateBookingEvent{
		EventType: "BOOKING_CREATED",
		Id:        ulid.Make().String(),
		UnitId:    cmd.UnitId,
		From:      fromTime,
		To:        toTime,
		Guests:    cmd.Guests,
		Name:      cmd.Name,
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return s.Publisher.Publish("event.booking.create", bytes)
}
