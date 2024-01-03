package eventstore

import (
	"context"
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("failed to marshal booking: %v", err)
	}

	metadata, err := json.Marshal(req.BookingEvent.Metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %v", err)
	}

	eventData := esdb.EventData{
		EventID:     uuid.Must(uuid.NewV4()),
		EventType:   req.BookingEvent.EventType.String(),
		ContentType: esdb.ContentTypeJson,
		Data:        booking,
		Metadata:    metadata,
	}

	_, err = b.db.AppendToStream(ctx, "bookings-stream", esdb.AppendToStreamOptions{}, eventData)

	return &pb.WriteBookingEventResponse{}, nil
}

func (b *BookingEventService) ReadBookingEvent(ctx context.Context, req *pb.ReadBookingEventsRequest) (*pb.ReadBookingEventsResponse, error) {
	return &pb.ReadBookingEventsResponse{}, fmt.Errorf("not implemented")
}
