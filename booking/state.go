package booking

import (
	"time"

	pb "github.com/thomaskrut/tekf/booking/pb/protos/v1"
)

type Booking struct {
	Id     string    `json:"id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Guests int       `json:"guests"`
	Name   string    `json:"name"`
}

type UnitBookings map[int][]Booking

type State struct {
	LastEventSequenceNumber int
	UnitBookings            UnitBookings
}

func (s *State) Apply(event *pb.BookingEvent) {
	switch event.EventType {
	case pb.EventType_EVENT_TYPE_CREATE_BOOKING:
		s.applyCreateBooking(event)
	}
}

func (s *State) applyCreateBooking(event *pb.BookingEvent) {
	booking := Booking{
		Id:     event.Booking.GetId(),
		From:   event.Booking.From.AsTime(),
		To:     event.Booking.GetTo().AsTime(),
		Guests: int(event.Booking.GetGuests()),
		Name:   event.Booking.GetName(),
	}

	unitId := int(event.Booking.GetUnitId())
	if s.UnitBookings == nil {
		s.UnitBookings = make(UnitBookings)
	}

	s.UnitBookings[unitId] = append(s.UnitBookings[unitId], booking)
}

func (s *State) checkAvailability(unitId int, from time.Time, to time.Time) bool {
	if s.UnitBookings == nil {
		return true
	}

	if s.UnitBookings[unitId] == nil {
		return true
	}

	for _, booking := range s.UnitBookings[unitId] {
		if booking.From.Before(to) && booking.To.After(from) {
			return false
		}
	}

	return true
}
