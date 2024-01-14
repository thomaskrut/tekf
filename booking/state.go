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
	UnitBookings UnitBookings
}

func (s *State) Apply(event *pb.BookingEvent) error {
	switch event.EventType {
	case pb.EventType_EVENT_TYPE_BOOKING_CREATED:
		return s.applyCreateBooking(event)
	case pb.EventType_EVENT_TYPE_BOOKING_DELETED:
		return s.applyDeleteBooking(event)
	case pb.EventType_EVENT_TYPE_BOOKING_UPDATED:
		return s.applyUpdateBooking(event)
	}
	return nil
}

func (s *State) deleteBooking(bookingId string) error {
	booking := s.getBooking(bookingId)
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

func createBooking(event *pb.BookingEvent) (*Booking, error) {
	fromTime, err := time.Parse("2006-01-02", event.Booking.From)
	if err != nil {
		return nil, ErrUnableToParseDate
	}

	toTime, err := time.Parse("2006-01-02", event.Booking.To)
	if err != nil {
		return nil, ErrUnableToParseDate
	}

	return &Booking{
		Id:     event.Booking.GetId(),
		From:   fromTime,
		To:     toTime,
		Guests: int(event.Booking.GetGuests()),
		Name:   event.Booking.GetName(),
		UnitId: int(event.Booking.GetUnitId()),
	}, nil
}

func (s *State) applyDeleteBooking(event *pb.BookingEvent) error {
	return s.deleteBooking(event.Booking.Id)
}

func (s *State) applyUpdateBooking(event *pb.BookingEvent) error {

	if err := s.deleteBooking(event.Booking.Id); err != nil {
		return ErrBookingNotFound
	}

	booking, err := createBooking(event)
	if err != nil {
		return err
	}

	s.UnitBookings[booking.UnitId] = append(s.UnitBookings[booking.UnitId], *booking)

	return nil
}

func (s *State) applyCreateBooking(event *pb.BookingEvent) error {

	booking, err := createBooking(event)
	if err != nil {
		return err
	}

	s.UnitBookings[booking.UnitId] = append(s.UnitBookings[booking.UnitId], *booking)

	return nil
}

func (s *State) checkAvailability(unitId int, from time.Time, to time.Time, excludeId string) bool {
	if s.UnitBookings == nil {
		return true
	}

	if s.UnitBookings[unitId] == nil {
		return true
	}

	for _, booking := range s.UnitBookings[unitId] {
		if booking.Id == excludeId {
			continue
		}
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
