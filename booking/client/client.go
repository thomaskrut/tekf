package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/thomaskrut/tekf/pb/protos/v1"
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

func (c *Client) Close() {
	c.connection.Close()
}
