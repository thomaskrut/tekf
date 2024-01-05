package com.thomaskrut.query.Model;

import java.time.LocalDate;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class Calendar {
    private final List<Day> days;

    private final List<Integer> units;

    public List<Integer> getUnits() {
        return units;
    }

    public Calendar() {
        this.days = new ArrayList<>();
        this.units = List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

        LocalDate startDate = LocalDate.now();
        LocalDate endDate = startDate.plusDays(100);

        startDate.datesUntil(endDate).forEach(date -> {
            Day day = new Day(date, units);
            days.add(day);
        });

    }

    public void addBooking(Booking booking) {
        days.stream().
                filter(day ->
                         day.getDate().isEqual(booking.getFromAsDate())
                        || (day.getDate().isAfter(booking.getFromAsDate()) && day.getDate().isBefore(booking.getToAsDate()))).
                forEach(day -> {
                            day.addBooking(booking);
                });
    }

    public List<Day> getDays() {
        return days;
    }

}

class Day {
    private final LocalDate date;

    private String weekday;
    private HashMap<Integer, Booking> bookings;

    public Day(LocalDate date, List<Integer> units) {
        this.date = date;

        this.weekday = date.getDayOfWeek().name();

        bookings = new HashMap<>();
        units.forEach(u -> {
            bookings.put(u, Booking.emptyBooking(""));
        });
    }

    public LocalDate getDate() {
        return date;
    }

    public String getWeekday() {
        return weekday;
    }

    public void setBookings(HashMap<Integer, Booking> bookings) {
        this.bookings = bookings;
    }

    public HashMap<Integer, Booking> getBookings() {
        return bookings;
    }

    public List<Integer> getUnits() {
        return bookings.keySet().stream().toList();
    }

    public void addBooking(Booking booking) {
        bookings.put(booking.getUnitId(), booking);
    }
}
