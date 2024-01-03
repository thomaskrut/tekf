package booking

import (
	"time"

	pb "github.com/thomaskrut/tekf/pb/protos/v1"
)

type Booking struct {
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
		From:   event.GetFrom().AsTime(),
		To:     event.GetTo().AsTime(),
		Guests: int(event.GetGuests()),
		Name:   event.GetName(),
	}

	unitId := int(event.GetUnitId())
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

		if booking.From.Equal(from) || booking.To.Equal(to) {
			return false
		}

		if booking.From.Before(from) && booking.To.After(from) {
			return false
		}
		if booking.From.Before(to) && booking.To.After(to) {
			return false
		}
		if booking.From.After(from) && booking.To.Before(to) {
			return false
		}

		if booking.From.After(from) && booking.From.Before(to) {
			return false
		}

		if booking.To.After(from) && booking.To.Before(to) {
			return false
		}
	}

	return true
}
