package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutionException;

@Service
public class BookingsModel {

    private EventStoreDBClient client;
    private List<Booking> bookings = new ArrayList<>();

    public BookingsModel() throws ExecutionException, InterruptedException {
        EventStoreDBClientSettings settings = EventStoreDBConnectionString.parseOrThrow("esdb://localhost:2113?tls=false");
        this.client = EventStoreDBClient.create(settings);

        this.bookings = this.Read(0);
    }

    public List<Booking> getBookings() {
        return bookings;
    }

    public List<Booking> Read(int lastKnownVersion) throws ExecutionException, InterruptedException {

        ReadStreamOptions options = ReadStreamOptions.get()
                .forwards()
                .fromRevision(lastKnownVersion);

        ReadResult result = client.readStream("bookings-stream", options).get();

        ObjectMapper mapper = new ObjectMapper();

        result.getEvents().forEach(event -> {
            try {

                // Events need to be applied to state

                Booking booking = mapper.readValue(event.getEvent().getEventData(), Booking.class);
                bookings.add(booking);
            } catch (IOException e) {
                throw new RuntimeException(e);
            }

        });

        return bookings;
    }

}
