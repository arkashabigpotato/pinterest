FROM golang:1.17-alpine3.13 as build

# Copy go.mod
RUN mkdir /home/pinterest
COPY ./go.mod /home/pinterest
WORKDIR /home/pinterest

# Download libs
RUN go mod download

# Copy all project files
COPY . /home/pinterest

# Run build application
RUN CGO_ENABLED=0 go build -o service ./cmd/service/

# Start application
FROM alpine:3.14.6

# Copy files to new alpine image
RUN mkdir pinterest
COPY --from=build /home/pinterest/service .
RUN mkdir static
COPY --from=build /home/pinterest/static ./static

# Run service
CMD ["./service"]
