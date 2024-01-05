package booking

import (
	"log"
	"time"

	pb "github.com/thomaskrut/tekf/booking/pb/protos/v1"
)

type Booking struct {
	Id     string    `json:"id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Guests int       `json:"guests"`
	Name   string    `json:"name"`
	UnitId int       `json:"unitId"`
}

type UnitBookings map[int][]Booking

type State struct {
	LastEventSequenceNumber int
	UnitBookings            UnitBookings
}

func (s *State) Apply(event *pb.BookingEvent) error {
	switch event.EventType {
	case pb.EventType_EVENT_TYPE_CREATE_BOOKING:
		return s.applyCreateBooking(event)
	case pb.EventType_EVENT_TYPE_DELETE_BOOKING:
		return s.applyDeleteBooking(event)
	}
	return nil
}

func (s *State) applyDeleteBooking(event *pb.BookingEvent) error {
	booking := s.getBooking(event.Booking.GetId())
	if booking == nil {
		log.Println("Illegal state: Booking not found")
		return nil
	}

	unitId := booking.UnitId

	for i, b := range s.UnitBookings[unitId] {
		if b.Id == booking.Id {
			log.Println("Deleted booking:", booking.Id)
			s.UnitBookings[unitId] = append(s.UnitBookings[unitId][:i], s.UnitBookings[unitId][i+1:]...)
			break
		}
	}

	return nil
}

func (s *State) applyCreateBooking(event *pb.BookingEvent) error {

	fromTime, err := time.Parse("2006-01-02", event.Booking.From)
	if err != nil {
		return ErrUnableToParseDate
	}

	toTime, err := time.Parse("2006-01-02", event.Booking.To)
	if err != nil {
		return ErrUnableToParseDate
	}

	booking := Booking{
		Id:     event.Booking.GetId(),
		From:   fromTime,
		To:     toTime,
		Guests: int(event.Booking.GetGuests()),
		Name:   event.Booking.GetName(),
		UnitId: int(event.Booking.GetUnitId()),
	}

	unitId := int(event.Booking.GetUnitId())
	if s.UnitBookings == nil {
		s.UnitBookings = make(UnitBookings)
	}

	s.UnitBookings[unitId] = append(s.UnitBookings[unitId], booking)

	return nil
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

func (s *State) getBooking(id string) *Booking {
	for _, bookings := range s.UnitBookings {
		for _, booking := range bookings {
			if booking.Id == id {
				return &booking
			}
		}
	}
	return nil
}
