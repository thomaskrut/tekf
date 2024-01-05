package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.concurrent.ExecutionException;

@Service
public class CalendarModel {

    private final EventStoreDBClient client;
    private final Calendar calendar;
    private final HashMap<String, Booking> bookings;

    public CalendarModel() throws ExecutionException, InterruptedException {
        EventStoreDBClientSettings settings = EventStoreDBConnectionString.parseOrThrow("esdb://localhost:2113?tls=false");
        this.client = EventStoreDBClient.create(settings);
        this.calendar = new Calendar();
        this.bookings = new HashMap<>();
        readStream(0);
    }

    public Calendar getCalendar() {
        return calendar;
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
                    case "EVENT_TYPE_CREATE_BOOKING" -> bookings.put(booking.getId(), booking);
                    //case "EVENT_TYPE_UPDATE_BOOKING" -> update(booking);
                    case "EVENT_TYPE_DELETE_BOOKING" -> bookings.remove(booking.getId());
                }

            } catch (IOException e) {
                throw new RuntimeException(e);
            }

        });

        bookings.values().forEach(calendar::addBooking);
    }

}
