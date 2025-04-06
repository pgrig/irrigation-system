# syntax=docker/dockerfile:1

FROM golang:1.24.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /irrigation-system

EXPOSE 8080

CMD ["/irrigation-system"]