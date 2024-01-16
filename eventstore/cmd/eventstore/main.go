package main

import (
	"log"
	"net"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/thomaskrut/tekf/eventstore"
	v1 "github.com/thomaskrut/tekf/eventstore/pb/protos/v1"
	"google.golang.org/grpc"
)

func main() {

	// Connect to EventStoreDB
	settings, err := esdb.ParseConnectionString("esdb://eventstore-db-service:2113?tls=false")

	if err != nil {
		log.Fatalf("Failed to parse EventStoreDB connection string: %v", err)
	}

	db, err := esdb.NewClient(settings)
	if err != nil {
		log.Fatalf("Failed to connect to EventStoreDB: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterBookingEventServiceServer(s, eventstore.NewBookingEventServiceServer(db))

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting server on port :9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
