package eventstore

import (
	"context"
	"fmt"
	"log"

	pb "github.com/thomaskrut/tekf/eventstore/pb/protos/v1"
)

type BookingEventService struct {
	pb.BookingEventServiceServer
}

func NewBookingEventServiceServer() *BookingEventService {
	return &BookingEventService{}
}

func (b *BookingEventService) WriteBookingEvent(ctx context.Context, req *pb.WriteBookingEventRequest) (*pb.WriteBookingEventResponse, error) {
	log.Println("WriteBookingEvent RPC", req.BookingEvent)
	return &pb.WriteBookingEventResponse{}, fmt.Errorf("not implemented")
}

func (b *BookingEventService) ReadBookingEvent(ctx context.Context, req *pb.ReadBookingEventsRequest) (*pb.ReadBookingEventsResponse, error) {
	return &pb.ReadBookingEventsResponse{}, fmt.Errorf("not implemented")
}
