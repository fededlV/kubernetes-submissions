# ---------- Stage 1: build ----------
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o log-writer .

# ---------- Stage 2: run ----------
FROM scratch
COPY --from=build /app/log-writer /log-writer
ENTRYPOINT ["/log-writer"]