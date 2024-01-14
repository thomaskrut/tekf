package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.LocalDate;
import java.util.HashMap;
import java.util.List;
import java.util.concurrent.ExecutionException;

@Service
public class DashboardModel {

    private long lastKnownRevision;
    private final List<Integer> units;
    private LocalDate today;
    private EventStoreDBClient client;
    private HashMap<String, Booking> checkIns;
    private HashMap<String, Booking> checkOuts;
    private HashMap<String, Booking> occupied;

    @Value("${eventstore-db.url}")
    private String eventstoreDbUrl;

    public DashboardModel() {
        this.lastKnownRevision = 0;
        this.units = List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20);
    }

    public void readStream(long fromRevision) throws ExecutionException, InterruptedException {

        if (this.client == null) {
            EventStoreDBClientSettings settings = EventStoreDBConnectionString.parseOrThrow(this.eventstoreDbUrl);
            this.client = EventStoreDBClient.create(settings);
            this.checkIns = new HashMap<>();
            this.checkOuts = new HashMap<>();
            this.occupied = new HashMap<>();
            this.today = LocalDate.of(LocalDate.now().getYear(), LocalDate.now().getMonth(),
                    LocalDate.now().getDayOfMonth());
        }

        System.out.println(today);

        ReadStreamOptions options = ReadStreamOptions.get()
                .forwards()
                .fromRevision(fromRevision);

        ReadResult result = client.readStream("bookings-stream", options).get();

        ObjectMapper mapper = new ObjectMapper();

        result.getEvents().forEach(event -> {

            this.lastKnownRevision = event.getEvent().getRevision();
            System.out.println(this.lastKnownRevision);
            try {
                Booking booking = mapper.readValue(event.getEvent().getEventData(), Booking.class);
                switch (event.getEvent().getEventType()) {
                    case "EVENT_TYPE_BOOKING_UPDATED", "EVENT_TYPE_BOOKING_CREATED" -> {

                        deleteBooking(booking);

                        if (booking.getFromAsDate().isEqual(today)) {
                            this.checkIns.put(booking.getId(), booking);
                        } else if (booking.getToAsDate().isEqual(today)) {
                            this.checkOuts.put(booking.getId(), booking);
                            this.occupied.put(booking.getId(), booking);
                        } else if (booking.getFromAsDate().isBefore(today) && booking.getToAsDate().isAfter(today)) {
                            this.occupied.put(booking.getId(), booking);
                        }
                    }

                    case "EVENT_TYPE_BOOKING_DELETED" -> deleteBooking(booking);

                    case "EVENT_TYPE_BOOKING_CHECKED_IN" -> this.checkIns.values().stream()
                            .filter(b -> b.getId().equals(booking.getId()))
                            .findFirst().ifPresent(b -> b.setCheckedIn(true));

                    case "EVENT_TYPE_BOOKING_CHECKED_OUT" -> this.checkOuts.values().stream()
                            .filter(b -> b.getId().equals(booking.getId()))
                            .findFirst().ifPresent(b -> b.setCheckedOut(true));

                }

            } catch (IOException e) {
                throw new RuntimeException(e);
            }

        });
    }

    public void update(LocalDate date) throws ExecutionException, InterruptedException {

        if (date.equals(today)) {
            readStream(this.lastKnownRevision + 1);
        } else {
            this.today = date;
            readStream(0);
        }

    }

    private void deleteBooking(Booking booking) {
        this.checkIns.remove(booking.getId());
        this.checkOuts.remove(booking.getId());
        this.occupied.remove(booking.getId());
    }

    public Booking getCheckinForUnit(int unit) {
        return this.checkIns.values().stream()
                .filter(b -> b.getUnitId() == unit)
                .findFirst().orElse(null);
    }

    public Booking getCheckoutForUnit(int unit) {
        return this.checkOuts.values().stream()
                .filter(b -> b.getUnitId() == unit)
                .findFirst().orElse(null);
    }

    public Booking getOccupiedForUnit(int unit) {
        return this.occupied.values().stream()
                .filter(b -> b.getUnitId() == unit)
                .findFirst().orElse(null);
    }

    public List<Integer> getUnits() {
        return units;
    }

}
