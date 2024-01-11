package com.thomaskrut.query.Model;

import java.time.LocalDate;
import java.util.ArrayList;
import java.util.List;

public class Calendar {
    private final List<Day> days;

    private final List<Integer> units;

    public Calendar() {
        this.days = new ArrayList<>();
        this.units = List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20);

        LocalDate startDate = LocalDate.now().minusDays(14);
        LocalDate endDate = startDate.plusDays(100);

        startDate.datesUntil(endDate).forEach(date -> {
            Day day = new Day(date, units);
            days.add(day);
        });

    }

    public void clear() {
        days.forEach(Day::clear);
    }

    public void addBooking(Booking booking) {
        days.stream().filter(day -> day.getDate().isEqual(booking.getFromAsDate())
                || (day.getDate().isAfter(booking.getFromAsDate()) && day.getDate().isBefore(booking.getToAsDate())))
                .forEach(day -> {
                    day.addBooking(booking);
                });
    }

    public List<Day> getDays() {
        return days;
    }

     public List<Integer> getUnits() {
        return units;
    }

}

