# syntax=docker/dockerfile:1

FROM golang:1.21.4
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /booking cmd/booking/main.go
CMD ["/booking"]