package booking

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestState_checkAvailability(t *testing.T) {
	tests := []struct {
		name     string
		bookings UnitBookings
		from     time.Time
		to       time.Time
		want     bool
	}{
		{
			name: "no bookings",
			from: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "no bookings for unit",
			bookings: UnitBookings{
				0: []Booking{
					{
						From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "no bookings for unit in time range",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
					},
					{
						From: time.Date(2021, 3, 10, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 3, 21, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "booking for unit in time range",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "booking for unit in time range 2",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 23, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 11, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "booking for unit in time range 3",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 23, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 22, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 25, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "booking starts on same day as existing booking ends",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 23, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 23, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 24, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "booking ends on same day as existing booking starts",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 23, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "booking for unit in time range 4",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 15, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "booking for unit in time range 5",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 19, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 15, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "booking for unit in time range 6",
			bookings: UnitBookings{
				1: []Booking{
					{
						From: time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2021, 1, 18, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			from: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2021, 1, 15, 0, 0, 0, 0, time.UTC),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			s := &State{
				UnitBookings: tt.bookings,
			}
			is.Equal(s.checkAvailability(1, tt.from, tt.to), tt.want)
		})
	}
}
