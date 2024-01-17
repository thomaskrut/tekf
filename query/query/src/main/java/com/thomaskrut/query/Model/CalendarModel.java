package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.time.LocalDate;
import java.util.HashMap;
import java.util.concurrent.ExecutionException;

@Service
public class CalendarModel extends ReadModel {


    private final Calendar calendar;
    private final HashMap<String, Booking> bookings;

    public CalendarModel() {
        calendar = new Calendar(units);
        bookings = new HashMap<>();
        logger =  LoggerFactory.getLogger(CalendarModel.class);
    }

    public Calendar getCalendar() {
        return calendar;
    }

    @Override
    public void readStream(long fromRevision) throws ExecutionException, InterruptedException {

        ReadStreamOptions options = ReadStreamOptions.get()
                .forwards()
                .fromRevision(fromRevision);

        ReadResult result = client.readStream("bookings-stream", options).get();

        ObjectMapper mapper = new ObjectMapper();

        result.getEvents().forEach(event -> {

            lastKnownRevision = event.getEvent().getRevision();
            logger.info("Processing event: " + event.getEvent().getRevision() + " " + event.getEvent().getEventType());
            try {
                Booking booking = mapper.readValue(event.getEvent().getEventData(), Booking.class);
                switch (event.getEvent().getEventType()) {
                    case "EVENT_TYPE_BOOKING_CREATED", "EVENT_TYPE_BOOKING_UPDATED" ->
                        bookings.put(booking.getId(), booking);
                    case "EVENT_TYPE_BOOKING_DELETED" -> bookings.remove(booking.getId());
                }

            } catch (IOException e) {
                logger.error("Error processing event: " + e.getMessage(), e);
            }

        });
        calendar.clear();
        bookings.values().forEach(calendar::addBooking);
    }

    @Override
    public void update(LocalDate date) throws ExecutionException, InterruptedException {
        readStream(this.lastKnownRevision + 1);
    }


}
