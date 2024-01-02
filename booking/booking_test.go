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
	today             = time.Now().Format("2006-01-02")
	tomorrow          = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	yesterday         = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	futureBookingFrom = time.Now().AddDate(0, 0, 10).Format("2006-01-02")
	futureBookingTo   = time.Now().AddDate(0, 0, 11).Format("2006-01-02")
)

type mockEventStoreClient struct{}

func (m *mockEventStoreClient) Write(context.Context, *pb.BookingEvent) error {
	return nil
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

			err := h.HandleCreateBookingCommand(CreateBookingCommand{
				UnitId: 1,
				From:   futureBookingFrom,
				To:     futureBookingTo,
				Guests: 2,
				Name:   "Mr. Guest",
			})
			is.NoErr(err)

			err = h.HandleCreateBookingCommand(tt.command)
			if tt.wantErr {
				is.True(err != nil)
				is.True(errors.Is(err, tt.wantErrType))
				return
			}
			is.NoErr(err)
		})
	}
}
