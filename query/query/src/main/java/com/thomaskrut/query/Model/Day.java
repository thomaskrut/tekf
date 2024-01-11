package com.thomaskrut.query.Model;

import java.time.LocalDate;
import java.util.HashMap;
import java.util.List;

public class Day {
    private final LocalDate date;

    private String weekday;
    private HashMap<Integer, Booking> bookings;
    private boolean isFirstDayOfMonth;

    public Day(LocalDate date, List<Integer> units) {
        this.date = date;

        this.weekday = date.getDayOfWeek().name();

        this.isFirstDayOfMonth = date.getDayOfMonth() == 1;

        bookings = new HashMap<>();
        units.forEach(u -> {
            bookings.put(u, Booking.emptyBooking(""));
        });
    }

    public void clear() {
        bookings.forEach((unitId, booking) -> {
            bookings.put(unitId, Booking.emptyBooking(""));
        });
    }

    public LocalDate getDate() {
        return date;
    }

    public String getWeekday() {
        return weekday;
    }

    public boolean isFirstDayOfMonth() {
        return isFirstDayOfMonth;
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
