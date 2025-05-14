FROM golang:1.24-bookworm AS builder
ARG CMD=apiWrapper

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /build
COPY cmd/ cmd/
COPY config/ config/
COPY internal/ internal/

WORKDIR /build/cmd/$CMD
RUN CGO_ENABLED=0 go build -o main

FROM debian:bookworm-slim 
ARG CMD=apiWrapper
WORKDIR /prod

RUN apt-get update && \
    apt-get install -y ca-certificates

COPY ./certs/ca-certificates/* /usr/local/share/ca-certificates/
RUN update-ca-certificates


COPY --from=builder /build/cmd/$CMD/main .

EXPOSE 8000

CMD [ "./main" ]