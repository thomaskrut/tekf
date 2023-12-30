package booking

import (
	"context"
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
	pb "github.com/thomaskrut/tekf/booking/pb/protos/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateBookingCommand struct {
	UnitId int    `json:"unitId"`
	From   string `json:"from"`
	To     string `json:"to"`
	Guests int    `json:"guests"`
	Name   string `json:"name"`
}

type BookingEvent struct {
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

type eventStoreClient interface {
	Write(context.Context, *pb.BookingEvent) error
}

type BookingCommandHandler struct {
	Publisher        publisher
	EventStoreClient eventStoreClient
}

func NewBookingCommandHandler(p publisher, e eventStoreClient) *BookingCommandHandler {
	return &BookingCommandHandler{
		Publisher:        p,
		EventStoreClient: e,
	}
}

func (s *BookingCommandHandler) HandleCreateBookingCommand(cmd CreateBookingCommand) error {

	fromTime, err := time.Parse("2006-01-02", cmd.From)
	if err != nil {
		return err
	}
	protoTimeFrom := timestamppb.New(fromTime)

	toTime, err := time.Parse("2006-01-02", cmd.To)
	if err != nil {
		return err
	}
	protoTimeTo := timestamppb.New(toTime)

	// Check state to see if booking is possible

	event := pb.BookingEvent{
		EventType: pb.EventType_EVENT_TYPE_CREATE_BOOKING,
		Id:        ulid.Make().String(),
		UnitId:    int32(cmd.UnitId),
		From:      protoTimeFrom,
		To:        protoTimeTo,
		Guests:    int32(cmd.Guests),
		Name:      cmd.Name,
	}

	bytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	err = s.EventStoreClient.Write(context.Background(), &event)
	if err != nil {
		return err
	}

	return s.Publisher.Publish("event.booking.create", bytes)
}
