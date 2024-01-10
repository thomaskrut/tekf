package com.thomaskrut.query.Model;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.LocalDate;

public class Booking {
    @JsonProperty("id")
    private String Id;
    @JsonProperty("from")
    private String from;
    @JsonProperty("to")
    private String to;
    @JsonProperty("guests")
    private int guests;
    @JsonProperty("name")
    private String name;
    @JsonProperty("unit_id")
    private int unitId;

    public static Booking emptyBooking(String bookingId) {
        Booking booking = new Booking();
        booking.setId(bookingId);
        return booking;
    }

    public String getId() {
        return Id;
    }

    public void setId(String id) {
        Id = id;
    }

    public String getFrom() {
        return from;
    }

    public void setFrom(String from) {
        this.from = from;
    }

    public String getTo() {
        return to;
    }

    public void setTo(String to) {
        this.to = to;
    }

    public int getGuests() {
        return guests;
    }

    public void setGuests(int guests) {
        this.guests = guests;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getUnitId() {
        return unitId;
    }

    public void setUnitId(int unitId) {
        this.unitId = unitId;
    }

    public LocalDate getFromAsDate() {
        return LocalDate.parse(from);
    }

    public LocalDate getToAsDate() {
        return LocalDate.parse(to);
    }

    public String getColor() {
        int hash = this.Id.hashCode();
        int lowerBound = 0xFF9999; // Light red lower bound
        int upperBound = 0xFFCCCC; // Light red upper bound
        int range = upperBound - lowerBound;
        int colorInt = Math.abs(hash % range) + lowerBound;
        String colorHex = String.format("#%06X", colorInt);
        return colorHex;
    }

}
