package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.HashMap;
import java.util.concurrent.ExecutionException;

@Service
public class CalendarModel {

    private final EventStoreDBClient client;
    private final HashMap<String, Booking> calendar = new HashMap<>();

    public CalendarModel() throws ExecutionException, InterruptedException {
        EventStoreDBClientSettings settings = EventStoreDBConnectionString.parseOrThrow("esdb://localhost:2113?tls=false");
        this.client = EventStoreDBClient.create(settings);

        readStream(0);
    }

    public HashMap<String, Booking> getBookings() {
        return calendar;
    }

    private void add(Booking booking) {
        this.calendar.put(booking.getId(), booking);
    }

    private void update(Booking booking) {
        this.calendar.put(booking.getId(), booking);
    }

    private void delete(Booking booking) {
        this.calendar.remove(booking.getId());
    }

    public void readStream(int lastKnownVersion) throws ExecutionException, InterruptedException {

        ReadStreamOptions options = ReadStreamOptions.get()
                .forwards()
                .fromRevision(lastKnownVersion);

        ReadResult result = client.readStream("bookings-stream", options).get();

        ObjectMapper mapper = new ObjectMapper();

        result.getEvents().forEach(event -> {
            try {
                Booking booking = mapper.readValue(event.getEvent().getEventData(), Booking.class);
                switch (event.getEvent().getEventType()) {
                    case "EVENT_TYPE_CREATE_BOOKING" -> add(booking);
                    case "EVENT_TYPE_UPDATE_BOOKING" -> update(booking);
                    case "EVENT_TYPE_DELETE_BOOKING" -> delete(booking);
                }

            } catch (IOException e) {
                throw new RuntimeException(e);
            }

        });
    }

}
