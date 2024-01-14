package checkinout

import (
	"time"

	pb "github.com/thomaskrut/tekf/checkinout/pb/protos/v1"
)

type Booking struct {
	Id     string    `json:"id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Guests int       `json:"guests"`
	Name   string    `json:"name"`
	UnitId int       `json:"unitId"`
}

type Checkins map[string]bool
type Checkouts map[string]bool

type State struct {
	Checkins
	Checkouts
	today string
}

func (s *State) Apply(event *pb.BookingEvent) error {
	switch event.EventType {
	case pb.EventType_EVENT_TYPE_BOOKING_CREATED:
		return s.applyCreate(event.Booking)
	case pb.EventType_EVENT_TYPE_BOOKING_UPDATED:
		return s.applyUpdate(event.Booking)
	case pb.EventType_EVENT_TYPE_BOOKING_DELETED:
		return s.applyDelete(event.Booking.Id)
	case pb.EventType_EVENT_TYPE_BOOKING_CHECKED_IN:
		return s.applyCheckin(event.Booking.Id)
	case pb.EventType_EVENT_TYPE_BOOKING_CHECKED_OUT:
		return s.applyCheckout(event.Booking.Id)
	}
	return nil
}

func (s *State) applyCreate(booking *pb.Booking) error {
	if booking.From == s.today {
		s.Checkins[booking.Id] = false
	}

	if booking.To == s.today {
		s.Checkouts[booking.Id] = false
	}

	return nil
}

func (s *State) applyUpdate(booking *pb.Booking) error {

	delete(s.Checkins, booking.Id)
	delete(s.Checkouts, booking.Id)

	if booking.From == s.today {
		s.Checkins[booking.Id] = false
	}

	if booking.To == s.today {
		s.Checkouts[booking.Id] = false
	}
	return nil
}

func (s *State) applyDelete(id string) error {
	delete(s.Checkins, id)
	delete(s.Checkouts, id)
	return nil
}

func (s *State) applyCheckin(id string) error {
	status, exists := s.Checkins[id]
	if !exists {
		return ErrBookingNotFound
	}

	if status {
		return ErrBookingAlreadyCheckedIn
	}

	s.Checkins[id] = true
	return nil
}

func (s *State) applyCheckout(id string) error {
	status, exists := s.Checkouts[id]
	if !exists {
		return ErrBookingNotFound
	}

	if status {
		return ErrBookingAlreadyCheckedOut
	}

	s.Checkouts[id] = true
	return nil
}
