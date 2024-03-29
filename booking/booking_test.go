package booking

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/matryer/is"
	pb "github.com/thomaskrut/tekf/booking/pb/protos/v1"
)

var (
	today              = time.Now().Format("2006-01-02")
	tomorrow           = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
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

func (m *mockEventStoreClient) ReadAll(context.Context) ([]*pb.BookingEvent, error) {
	return []*pb.BookingEvent{
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "abc123",
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Guest",
				UnitId: 1,
			},
		},
		{
			EventType: pb.EventType_EVENT_TYPE_BOOKING_CREATED,
			Booking: &pb.Booking{
				Id:     "efg456",
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

func TestBookingCommandHandler_HandleCreateBookingCommand(t *testing.T) {
	tests := []struct {
		name        string
		command     CreateBookingCommand
		wantErr     bool
		wantErrType error
	}{
		{
			name: "happy path",
			command: CreateBookingCommand{
				UnitId: 1,
				From:   today,
				To:     tomorrow,
				Guests: 2,
				Name:   "Thomas",
			},
		},
		{
			name: "booking in the past",
			command: CreateBookingCommand{
				UnitId: 1,
				From:   yesterday,
				To:     tomorrow,
				Guests: 2,
				Name:   "Thomas",
			},
			wantErr:     true,
			wantErrType: ErrInvalidDateRange,
		},
		{
			name: "to after from",
			command: CreateBookingCommand{
				UnitId: 1,
				From:   tomorrow,
				To:     today,
				Guests: 2,
				Name:   "Thomas",
			},
			wantErr:     true,
			wantErrType: ErrInvalidDateRange,
		},
		{
			name: "invalid unit id",
			command: CreateBookingCommand{
				UnitId: 9999,
				From:   today,
				To:     tomorrow,
				Guests: 2,
				Name:   "Thomas",
			},
			wantErr:     true,
			wantErrType: ErrInvalidUnitId,
		},
		{
			name: "invalid guest count",
			command: CreateBookingCommand{
				UnitId: 1,
				From:   today,
				To:     tomorrow,
				Guests: 0,
				Name:   "Thomas",
			},
			wantErr:     true,
			wantErrType: ErrInvalidGuestCount,
		},
		{
			name: "invalid name",
			command: CreateBookingCommand{
				UnitId: 1,
				From:   today,
				To:     tomorrow,
				Guests: 2,
				Name:   "",
			},
			wantErr:     true,
			wantErrType: ErrInvalidGuestName,
		},
		{
			name: "dates unavailable",
			command: CreateBookingCommand{
				UnitId: 1,
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Guest Person",
			},
			wantErr:     true,
			wantErrType: ErrUnitNotAvailable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			h := NewBookingCommandHandler(&mockPublisher{}, &mockEventStoreClient{})
			h.LoadState()

			err := h.HandleCreateBookingCommand(tt.command)
			if tt.wantErr {
				is.True(err != nil)
				is.True(errors.Is(err, tt.wantErrType))
				return
			}
			is.NoErr(err)
		})
	}
}

func TestBookingCommandHandler_HandleDeleteBookingCommand(t *testing.T) {
	tests := []struct {
		name        string
		command     DeleteBookingCommand
		wantErr     bool
		wantErrType error
	}{
		{
			name: "happy path",
			command: DeleteBookingCommand{
				Id: "abc123",
			},
		},
		{
			name: "booking not found",
			command: DeleteBookingCommand{
				Id: "xyz456",
			},
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			h := NewBookingCommandHandler(&mockPublisher{}, &mockEventStoreClient{})
			h.LoadState()

			err := h.HandleDeleteBookingCommand(tt.command)
			if tt.wantErr {
				is.True(err != nil)
				is.True(errors.Is(err, tt.wantErrType))
				return
			}
			is.NoErr(err)
		})
	}
}

func TestBookingCommandHandler_TestUpdateBookingCommand(t *testing.T) {
	tests := []struct {
		name        string
		command     UpdateBookingCommand
		wantErr     bool
		wantErrType error
	}{
		{
			name: "happy path",
			command: UpdateBookingCommand{
				Id:     "abc123",
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Guest Person New Name",
				UnitId: 1,
			},
		},
		{
			name: "booking not found",
			command: UpdateBookingCommand{
				Id: "xyz456",
			},
			wantErr:     true,
			wantErrType: ErrBookingNotFound,
		},
		{
			name: "invalid guest count",
			command: UpdateBookingCommand{
				Id:     "abc123",
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 0,
				Name:   "Mr. Guest Person",
				UnitId: 1,
			},
			wantErr:     true,
			wantErrType: ErrInvalidGuestCount,
		},
		{
			name: "invalid name",
			command: UpdateBookingCommand{
				Id:     "abc123",
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "",
				UnitId: 1,
			},
			wantErr:     true,
			wantErrType: ErrInvalidGuestName,
		},
		{
			name: "dates unavailable",
			command: UpdateBookingCommand{
				Id:     "abc123",
				From:   futureBookingFrom2,
				To:     futureBookingTo2,
				Guests: 2,
				Name:   "Mr. Guest Person",
				UnitId: 1,
			},
			wantErr:     true,
			wantErrType: ErrUnitNotAvailable,
		},
		{
			name: "invalid date range",
			command: UpdateBookingCommand{
				Id:     "abc123",
				From:   futureBookingTo,
				To:     futureBookingFrom,
				Guests: 2,
				Name:   "Mr. Guest Person",
				UnitId: 1,
			},
			wantErr:     true,
			wantErrType: ErrInvalidDateRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			h := NewBookingCommandHandler(&mockPublisher{}, &mockEventStoreClient{})
			h.LoadState()

			err := h.HandleUpdateBookingCommand(tt.command)
			if tt.wantErr {
				is.True(err != nil)
				is.True(errors.Is(err, tt.wantErrType))
				return
			}
			is.NoErr(err)
		})
	}
}
