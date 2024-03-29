package checkinout

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/matryer/is"
	pb "github.com/thomaskrut/tekf/checkinout/pb/protos/v1"
)

var (
	today              = time.Now().Format("2006-01-02")
	yesterday          = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	futureBookingFrom  = time.Now().AddDate(0, 0, 10).Format("2006-01-02")
	futureBookingTo    = time.Now().AddDate(0, 0, 11).Format("2006-01-02")
	futureBookingFrom2 = time.Now().AddDate(0, 1, 10).Format("2006-01-02")
	futureBookingTo2   = time.Now().AddDate(0, 1, 11).Format("2006-01-02")
)

type mockEventStoreClient struct{}

func (m *mockEventStoreClient) Write(context.Context, *pb.BookingEvent) error {
	return nil
}

func (m *mockEventStoreClient) ReadLatest(context.Context) ([]*pb.BookingEvent, error) {
	return nil, nil
}

func (m *mockEventStoreClient) ReadAll(context.Context) ([]*pb.BookingEvent, error) {
	return []*pb.BookingEvent{
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "abc123",
				From:   today,
				To:     today,
				Guests: 2,
				Name:   "Mr. Guest",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "efg456",
				From:   yesterday,
				To:     futureBookingFrom,
				Guests: 2,
				Name:   "Mr. Person",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "hij789",
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Person",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "xyz987",
				From:   today,
				To:     today,
				Guests: 2,
				Name:   "Mr. Guest Already Checked In and Out",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CHECKED_IN,
			Booking: &pb.Booking{
				Id: "xyz987",
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CHECKED_OUT,
			Booking: &pb.Booking{
				Id: "xyz987",
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "uvw123",
				From:   today,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Person",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_DELETED,
			Booking: &pb.Booking{
				Id: "uvw123",
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "jkl012",
				From:   today,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Person",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_UPDATED,
			Booking: &pb.Booking{
				Id:     "jkl012",
				From:   futureBookingFrom2,
				To:     futureBookingTo2,
				Guests: 2,
				Name:   "Mr. Person",
				UnitId: 1,
			},
		},
	}, nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(string, []byte) error {
	return nil
}

func TestBookingCommandHandler_HandleCheckinCommand(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantErr     bool
		wantErrType error
	}{
		{
			name: "happy path",
			id:   "abc123",
		},
		{
			name:        "checkin not today",
			id:          "efg456",
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
		{
			name:        "already checked in",
			id:          "xyz987",
			wantErr:     true,
			wantErrType: ErrBookingAlreadyCheckedIn,
		},
		{
			name:        "booking deleted",
			id:          "uvw123",
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
		{
			name:        "booking changed",
			id:          "jkl012",
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			h := NewCommandHandler(&mockPublisher{}, &mockEventStoreClient{})
			h.LoadState()

			err := h.HandleCheckinCommand(tt.id)
			if tt.wantErr {
				is.True(err != nil)
				is.True(errors.Is(err, tt.wantErrType))
				return
			}
			is.NoErr(err)
		})
	}
}

func TestBookingCommandHandler_HandleCheckoutCommand(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantErr     bool
		wantErrType error
	}{
		{
			name: "happy path",
			id:   "abc123",
		},
		{
			name:        "checkout not today",
			id:          "efg456",
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
		{
			name:        "already checked out",
			id:          "xyz987",
			wantErr:     true,
			wantErrType: ErrBookingAlreadyCheckedOut,
		},
		{
			name:        "booking deleted",
			id:          "uvw123",
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
		{
			name:        "booking changed",
			id:          "jkl012",
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			h := NewCommandHandler(&mockPublisher{}, &mockEventStoreClient{})
			h.LoadState()

			err := h.HandleCheckoutCommand(tt.id)
			if tt.wantErr {
				is.True(err != nil)
				is.True(errors.Is(err, tt.wantErrType))
				return
			}
			is.NoErr(err)
		})
	}
}
