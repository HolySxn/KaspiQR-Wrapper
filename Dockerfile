FROM golang:1.24-bookworm AS build
ARG CMD=apiWrapper

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /app
COPY . .

WORKDIR /app/cmd/$CMD
RUN go build -o main



EXPOSE 8000

CMD [ "./main" ]