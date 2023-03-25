FROM golang:latest AS build

WORKDIR /opt
COPY main.go /opt/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM scratch
COPY --from=0 /opt/main /main

ENV COUNT_OF_CATS 60

ENTRYPOINT ["/main"]
