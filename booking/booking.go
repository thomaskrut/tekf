package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
	pb "github.com/thomaskrut/tekf/booking/pb/protos/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrInvalidUnitId     = fmt.Errorf("invalid unit id")
	ErrInvalidGuestName  = fmt.Errorf("invalid guest name")
	ErrInvalidGuestCount = fmt.Errorf("invalid guest count")
	ErrInvalidDateRange  = fmt.Errorf("invalid date range")
	ErrUnitNotAvailable  = fmt.Errorf("unit not available")
	ErrUnableToParseDate = fmt.Errorf("error parsing date value")
)

type CreateBookingCommand struct {
	UnitId int    `json:"unitId"`
	From   string `json:"from"`
	To     string `json:"to"`
	Guests int    `json:"guests"`
	Name   string `json:"name"`
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
	State            State
}

func NewBookingCommandHandler(p publisher, e eventStoreClient) *BookingCommandHandler {
	return &BookingCommandHandler{
		Publisher:        p,
		EventStoreClient: e,
		State:            State{},
	}
}

func (s *BookingCommandHandler) HandleCreateBookingCommand(cmd CreateBookingCommand) error {

	fromTime, err := time.Parse("2006-01-02", cmd.From)
	if err != nil {
		return ErrUnableToParseDate
	}
	protoTimeFrom := timestamppb.New(fromTime)

	toTime, err := time.Parse("2006-01-02", cmd.To)
	if err != nil {
		return ErrUnableToParseDate
	}
	protoTimeTo := timestamppb.New(toTime)

	if cmd.UnitId < 0 || cmd.UnitId > 10 {
		return ErrInvalidUnitId
	}

	if cmd.Guests < 1 || cmd.Guests > 5 {
		return ErrInvalidGuestCount
	}

	if cmd.Name == "" {
		return ErrInvalidGuestName
	}

	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	if fromTime.After(toTime) || fromTime.Equal(toTime) || fromTime.Before(today) {
		return ErrInvalidDateRange
	}

	if !s.State.checkAvailability(cmd.UnitId, fromTime, toTime) {
		return ErrUnitNotAvailable
	}

	log.Println(s.State.UnitBookings[cmd.UnitId])

	event := pb.BookingEvent{
		Id:        ulid.Make().String(),
		EventType: pb.EventType_EVENT_TYPE_CREATE_BOOKING,
		Metadata: &pb.Metadata{
			Timestamp: timestamppb.Now(),
		},
		Booking: &pb.Booking{
			UnitId: int32(cmd.UnitId),
			From:   protoTimeFrom,
			To:     protoTimeTo,
			Guests: int32(cmd.Guests),
			Name:   cmd.Name,
		},
	}

	err = s.EventStoreClient.Write(context.Background(), &event)
	if err != nil {
		return err
	}

	s.State.Apply(&event)

	bytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	return s.Publisher.Publish("event.booking.create", bytes)
}
