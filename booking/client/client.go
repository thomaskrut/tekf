package client

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/thomaskrut/tekf/booking/pb/protos/v1"
)

type Client struct {
	connection *grpc.ClientConn
	pb.BookingEventServiceClient
}

func New() *Client {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &Client{
		connection:                conn,
		BookingEventServiceClient: pb.NewBookingEventServiceClient(conn),
	}
}

func (c *Client) Write(ctx context.Context, event *pb.BookingEvent) error {
	req := &pb.WriteBookingEventRequest{
		BookingEvent: event,
	}
	_, err := c.BookingEventServiceClient.WriteBookingEvent(ctx, req)
	return err
}

func (c *Client) ReadAll(ctx context.Context) ([]*pb.BookingEvent, error) {
	req := &pb.ReadBookingEventsRequest{
		LastKnownEventId: 0,
	}

	stream, err := c.BookingEventServiceClient.ReadBookingEvents(ctx, req)
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
		events = append(events, response.BookingEvent)
	}

	return events, nil
}

func (c *Client) Close() {
	c.connection.Close()
}
