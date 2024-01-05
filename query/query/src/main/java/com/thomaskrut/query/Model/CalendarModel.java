package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.HashMap;
import java.util.concurrent.ExecutionException;

@Service
public class CalendarModel {

    private EventStoreDBClient client;
    private Calendar calendar;
    private HashMap<String, Booking> bookings;

    @Value("${eventstore-db.url}")
    private String eventstoreDbUrl;

    private long lastKnownRevision;

    public CalendarModel() {
        this.lastKnownRevision = 0;
    }

    public Calendar getCalendar() {
        return calendar;
    }


    public void readStream(long fromRevision) throws ExecutionException, InterruptedException {

        if (this.client == null) {
            EventStoreDBClientSettings settings = EventStoreDBConnectionString.parseOrThrow(this.eventstoreDbUrl);
            this.client = EventStoreDBClient.create(settings);
            this.calendar = new Calendar();
            this.bookings = new HashMap<>();
        }

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
                    case "EVENT_TYPE_CREATE_BOOKING", "EVENT_TYPE_UPDATE_BOOKING" -> bookings.put(booking.getId(), booking);
                    case "EVENT_TYPE_DELETE_BOOKING" -> bookings.remove(booking.getId());
                }

            } catch (IOException e) {
                throw new RuntimeException(e);
            }

        });

        bookings.values().forEach(calendar::addBooking);
    }

    public void update() throws ExecutionException, InterruptedException {
        readStream(this.lastKnownRevision + 1);
    }
}
