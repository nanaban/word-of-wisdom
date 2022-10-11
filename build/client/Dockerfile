## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -v -o client ./cmd/client/...

## Deploy
FROM debian:buster-slim as deploy
RUN set -x &&  \
    apt-get update &&  \
    DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /app/client /app/client

EXPOSE 3000

CMD ["/app/client", "-a", "host.docker.internal:3000", "-n", "1000", "-c", "24"]