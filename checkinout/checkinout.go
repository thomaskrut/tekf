package checkinout

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/thomaskrut/tekf/checkinout/pb/protos/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrBookingNotFound          = fmt.Errorf("booking not found")
	ErrBookingAlreadyCheckedIn  = fmt.Errorf("booking already checked in")
	ErrBookingAlreadyCheckedOut = fmt.Errorf("booking already checked out")
)

type publisher interface {
	Publish(subject string, data []byte) error
}

type eventStoreClient interface {
	Write(context.Context, *pb.BookingEvent) error
	ReadAll(context.Context) ([]*pb.BookingEvent, error)
	ReadLatest(context.Context) ([]*pb.BookingEvent, error)
}

type CommandHandler struct {
	Publisher        publisher
	EventStoreClient eventStoreClient
	State            *State
}

func NewCommandHandler(publisher publisher, eventStoreClient eventStoreClient) *CommandHandler {
	return &CommandHandler{
		Publisher:        publisher,
		EventStoreClient: eventStoreClient,
	}
}

func (b *CommandHandler) HandleCheckinCommand(id string) error {
	event := pb.BookingEvent{
		EventType: pb.EventType_EVENT_TYPE_CHECKIN_BOOKING,
		Metadata: &pb.Metadata{
			Timestamp: timestamppb.Now(),
		},
		Booking: &pb.Booking{
			Id: id,
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
	return b.Publisher.Publish("event.booking.checkin", bytes)
}

func (b *CommandHandler) HandleCheckoutCommand(id string) error {
	event := pb.BookingEvent{
		EventType: pb.EventType_EVENT_TYPE_CHECKOUT_BOOKING,
		Metadata: &pb.Metadata{
			Timestamp: timestamppb.Now(),
		},
		Booking: &pb.Booking{
			Id: id,
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
	return b.Publisher.Publish("event.booking.checkout", bytes)
}

func (b *CommandHandler) LoadState() error {

	var events []*pb.BookingEvent
	var err error
	ctx := context.Background()

	if b.State != nil {

		events, err = b.EventStoreClient.ReadLatest(ctx)
		if err != nil {
			return err
		}

	} else {

		b.State = &State{
			Checkins:  make(Checkins),
			Checkouts: make(Checkouts),
			today:     time.Now().Format("2006-01-02"),
		}

		events, err = b.EventStoreClient.ReadAll(ctx)
		if err != nil {
			return err
		}

	}

	for _, event := range events {
		log.Println(event)
		b.State.Apply(event)
	}

	return nil
}
