FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o log-reader .

FROM scratch
COPY --from=build /app/log-reader /log-reader
EXPOSE 3000
ENTRYPOINT ["/log-reader"]