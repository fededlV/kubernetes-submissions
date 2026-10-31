FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o todo-app .

FROM scratch
COPY --from=build /app/todo-app /todo-app
ENV PORT=3000
EXPOSE 3000
ENTRYPOINT ["/todo-app"]