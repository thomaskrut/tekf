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
	ErrBookingNotFound   = fmt.Errorf("booking not found")
)

type CreateBookingCommand struct {
	UnitId int    `json:"unitId"`
	From   string `json:"from"`
	To     string `json:"to"`
	Guests int    `json:"guests"`
	Name   string `json:"name"`
}

type DeleteBookingCommand struct {
	Id string `json:"id"`
}

type UpdateBookingCommand struct {
	Id     string `json:"id"`
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
	ReadAll(context.Context) ([]*pb.BookingEvent, error)
}

type BookingCommandHandler struct {
	Publisher        publisher
	EventStoreClient eventStoreClient
	State            *State
}

func NewBookingCommandHandler(p publisher, e eventStoreClient) *BookingCommandHandler {

	handler := &BookingCommandHandler{
		Publisher:        p,
		EventStoreClient: e,
	}

	return handler
}

func (b *BookingCommandHandler) LoadState() error {

	if b.State != nil {
		return nil
	}

	b.State = &State{
		UnitBookings: make(UnitBookings),
	}

	events, err := b.EventStoreClient.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, event := range events {
		log.Println(event)
		b.State.Apply(event)
	}

	return nil
}

func (b *BookingCommandHandler) HandleDeleteBookingCommand(cmd DeleteBookingCommand) error {

	booking := b.State.getBooking(cmd.Id)
	if booking == nil {
		return ErrBookingNotFound
	}

	event := pb.BookingEvent{
		EventType: pb.EventType_EVENT_TYPE_BOOKING_DELETED,
		Metadata: &pb.Metadata{
			Timestamp: timestamppb.Now(),
		},
		Booking: &pb.Booking{
			Id: cmd.Id,
		},
	}

	err := b.EventStoreClient.Write(context.Background(), &event)
	if err != nil {
		return err
	}

	if err = b.State.Apply(&event); err != nil {
		return err
	}

	bytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	// No subscribers to this event yet
	return b.Publisher.Publish("event.booking.delete", bytes)
}

func (b *BookingCommandHandler) HandleCreateBookingCommand(cmd CreateBookingCommand) error {

	event := pb.BookingEvent{
		EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
		Metadata: &pb.Metadata{
			Timestamp: timestamppb.Now(),
		},
		Booking: &pb.Booking{
			Id:     ulid.Make().String(),
			UnitId: int32(cmd.UnitId),
			From:   cmd.From,
			To:     cmd.To,
			Guests: int32(cmd.Guests),
			Name:   cmd.Name,
		},
	}

	if err := b.validateEvent(&event); err != nil {
		return err
	}

	err := b.EventStoreClient.Write(context.Background(), &event)
	if err != nil {
		return err
	}

	if err = b.State.Apply(&event); err != nil {
		return err
	}

	bytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	return b.Publisher.Publish("event.booking.create", bytes)
}

func (b *BookingCommandHandler) HandleUpdateBookingCommand(cmd UpdateBookingCommand) error {

	booking := b.State.getBooking(cmd.Id)
	if booking == nil {
		return ErrBookingNotFound
	}

	event := pb.BookingEvent{
		EventType: pb.EventType_EVENT_TYPE_BOOKING_UPDATED,
		Metadata: &pb.Metadata{
			Timestamp: timestamppb.Now(),
		},
		Booking: &pb.Booking{
			Id:     cmd.Id,
			UnitId: int32(cmd.UnitId),
			From:   cmd.From,
			To:     cmd.To,
			Guests: int32(cmd.Guests),
			Name:   cmd.Name,
		},
	}

	if err := b.validateEvent(&event); err != nil {
		return err
	}

	err := b.EventStoreClient.Write(context.Background(), &event)
	if err != nil {
		return err
	}

	if err = b.State.Apply(&event); err != nil {
		return err
	}

	bytes, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	return b.Publisher.Publish("event.booking.update", bytes)
}

func (b *BookingCommandHandler) validateEvent(event *pb.BookingEvent) error {
	fromTime, err := time.Parse("2006-01-02", event.Booking.From)
	if err != nil {
		return ErrUnableToParseDate
	}

	toTime, err := time.Parse("2006-01-02", event.Booking.To)
	if err != nil {
		return ErrUnableToParseDate
	}

	if event.Booking.UnitId < 0 || event.Booking.UnitId > 20 {
		return ErrInvalidUnitId
	}

	if event.Booking.Guests < 1 || event.Booking.Guests > 5 {
		return ErrInvalidGuestCount
	}

	if event.Booking.Name == "" {
		return ErrInvalidGuestName
	}

	//today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	if fromTime.After(toTime) || fromTime.Equal(toTime) /*|| fromTime.Before(today)*/ {
		return ErrInvalidDateRange
	}

	if !b.State.checkAvailability(int(event.Booking.UnitId), fromTime, toTime, event.Booking.Id) {
		return ErrUnitNotAvailable
	}

	return nil
}
