package main

import (
	"log"
	"net"

	"github.com/thomaskrut/tekf/eventstore"
	"github.com/thomaskrut/tekf/eventstore/pb/protos/v1"
	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()
	v1.RegisterBookingEventServiceServer(s, eventstore.NewBookingEventServiceServer())
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting server on port :9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
