package com.thomaskrut.query.Model;

import com.eventstore.dbclient.*;

import jakarta.annotation.PostConstruct;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;

import java.time.LocalDate;
import java.util.List;
import java.util.concurrent.ExecutionException;

public abstract class ReadModel {

    @Value("${eventstore-db.url}")
    protected String eventstoreDbUrl;

    protected long lastKnownRevision;

    protected EventStoreDBClient client;

    protected Logger logger;

    protected final List<Integer> units;

    public ReadModel() {
        logger =  LoggerFactory.getLogger(ReadModel.class);
        this.units = List.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20);
        this.lastKnownRevision = -1;
    }

    @PostConstruct
    public void initClient() {
        if (eventstoreDbUrl == null || eventstoreDbUrl.isEmpty()) {
            throw new IllegalArgumentException("eventstoreDbUrl cannot be null or empty");
        }
        EventStoreDBClientSettings settings = EventStoreDBConnectionString.parseOrThrow(eventstoreDbUrl);
        this.client = EventStoreDBClient.create(settings);
    }

    public abstract void readStream(long fromRevision) throws ExecutionException, InterruptedException;

    public abstract void update(LocalDate date) throws ExecutionException, InterruptedException;
}
