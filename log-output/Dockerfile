FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o log-output .

FROM scratch
COPY --from=build /app/log-output /log-output
ENTRYPOINT ["/log-output"]