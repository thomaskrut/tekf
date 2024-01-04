package eventstore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/gofrs/uuid"
	pb "github.com/thomaskrut/tekf/eventstore/pb/protos/v1"
)

type BookingEventService struct {
	db *esdb.Client
	pb.BookingEventServiceServer
}

func NewBookingEventServiceServer(db *esdb.Client) *BookingEventService {
	return &BookingEventService{
		db: db,
	}
}

func (b *BookingEventService) WriteBookingEvent(ctx context.Context, req *pb.WriteBookingEventRequest) (*pb.WriteBookingEventResponse, error) {
	log.Println("WriteBookingEvent RPC", req.BookingEvent)

	booking, err := json.Marshal(req.BookingEvent.Booking)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal booking: %w", err)
	}

	metadata, err := json.Marshal(req.BookingEvent.Metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	eventData := esdb.EventData{
		EventID:     uuid.Must(uuid.NewV4()),
		EventType:   req.BookingEvent.EventType.String(),
		ContentType: esdb.ContentTypeJson,
		Data:        booking,
		Metadata:    metadata,
	}

	_, err = b.db.AppendToStream(ctx, "bookings-stream", esdb.AppendToStreamOptions{}, eventData)
	if err != nil {
		return nil, fmt.Errorf("failed to append to stream: %w", err)
	}

	return &pb.WriteBookingEventResponse{
		Success: true,
	}, nil
}

func (b *BookingEventService) ReadBookingEvents(req *pb.ReadBookingEventsRequest, stream pb.BookingEventService_ReadBookingEventsServer) error {
	log.Println("ReadBookingEvents RPC", req)

	options := esdb.ReadStreamOptions{
		From:      esdb.Revision(uint64(req.LastKnownEventId)),
		Direction: esdb.Forwards,
	}

	eventStream, err := b.db.ReadStream(stream.Context(), "bookings-stream", options, 99999)
	if err != nil {
		return fmt.Errorf("failed to read stream: %w", err)
	}
	defer eventStream.Close()

	for {

		event, err := eventStream.Recv()

		if err, ok := esdb.FromError(err); !ok {

			if err.Code() == esdb.ErrorCodeResourceNotFound {
				return nil
			}

			if errors.Is(err, io.EOF) {
				break
			} else {
				return fmt.Errorf("failed to read event: %w", err)
			}
		}

		var bookingEvent pb.BookingEvent

		err = json.Unmarshal(event.Event.Data, &bookingEvent.Booking)
		if err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}

		bookingEvent.EventType = pb.EventType(pb.EventType_value[event.Event.EventType])

		if err = stream.Send(&pb.ReadBookingEventsResponse{
			BookingEvent: &bookingEvent,
		}); err != nil {
			return fmt.Errorf("failed to send event: %w", err)
		}
	}

	return nil

}
