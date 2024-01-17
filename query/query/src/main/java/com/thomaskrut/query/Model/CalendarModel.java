package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.thomaskrut.query.Controller.CalendarController;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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
    Logger logger = LoggerFactory.getLogger(CalendarModel.class);

    @Value("${eventstore-db.url}")
    private String eventstoreDbUrl;

    private long lastKnownRevision;

    public CalendarModel() {
        this.lastKnownRevision = -1;
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
            logger.info("Processing event: " + event.getEvent().getRevision() + " " + event.getEvent().getEventType());
            try {
                Booking booking = mapper.readValue(event.getEvent().getEventData(), Booking.class);
                switch (event.getEvent().getEventType()) {
                    case "EVENT_TYPE_BOOKING_CREATED", "EVENT_TYPE_BOOKING_UPDATED" ->
                        bookings.put(booking.getId(), booking);
                    case "EVENT_TYPE_BOOKING_DELETED" -> bookings.remove(booking.getId());
                }

            } catch (IOException e) {
                throw new RuntimeException(e);
            }

        });
        calendar.clear();
        bookings.values().forEach(calendar::addBooking);
    }

    public void update() throws ExecutionException, InterruptedException {
        readStream(this.lastKnownRevision + 1);
    }
}
