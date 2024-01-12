package client

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/thomaskrut/tekf/checkinout/pb/protos/v1"
)

var eventStoreServiceUrl = "eventstore-service:9090"

type Client struct {
	lastKnownRevision int32
}

func New() *Client {
	return &Client{}
}

func (c *Client) Write(ctx context.Context, event *pb.BookingEvent) error {

	conn, err := grpc.Dial(eventStoreServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewBookingEventServiceClient(conn)

	req := &pb.WriteBookingEventRequest{
		BookingEvent: event,
	}
	_, err = client.WriteBookingEvent(ctx, req)
	return err
}

func (c *Client) ReadAll(ctx context.Context) ([]*pb.BookingEvent, error) {

	conn, err := grpc.Dial(eventStoreServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewBookingEventServiceClient(conn)

	req := &pb.ReadBookingEventsRequest{
		LastKnownEventId: 0,
	}

	stream, err := client.ReadBookingEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error getting stream: %w", err)
	}

	var events []*pb.BookingEvent

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading stream: %w", err)
		}
		c.lastKnownRevision = response.Revision
		events = append(events, response.BookingEvent)
	}

	return events, nil
}

func (c *Client) ReadLatest(ctx context.Context) ([]*pb.BookingEvent, error) {

	conn, err := grpc.Dial(eventStoreServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewBookingEventServiceClient(conn)

	req := &pb.ReadBookingEventsRequest{
		LastKnownEventId: c.lastKnownRevision + 1,
	}

	stream, err := client.ReadBookingEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error getting stream: %w", err)
	}

	var events []*pb.BookingEvent

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading stream: %w", err)
		}
		c.lastKnownRevision = response.Revision
		events = append(events, response.BookingEvent)
	}

	return events, nil
}
