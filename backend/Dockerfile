FROM golang:1.25-alpine3.23

RUN apk add --no-cache docker-cli shadow

WORKDIR /app
COPY . .

RUN go build -o statuspagebackend .

CMD ["/app/statuspagebackend"]