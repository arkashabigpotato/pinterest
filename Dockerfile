FROM golang:1.17-alpine3.13

RUN mrdir /home/pinterest
COPY . /home/pinterest
WORKDIR /home/pinterest
RUN go mod download
RUN CGO_ENABLED=0 go build -o pinterest ./cmd/service/