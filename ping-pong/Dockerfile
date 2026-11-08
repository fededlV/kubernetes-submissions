FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o ping-pong .

FROM scratch
COPY --from=build /app/ping-pong /ping-pong
ENTRYPOINT ["/ping-pong"]