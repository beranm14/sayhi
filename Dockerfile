FROM golang:latest AS build

WORKDIR /opt
COPY main.go /opt/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM scratch
COPY --from=0 /opt/main /main

ENV PORT 3000

EXPOSE 3000/tcp

ENTRYPOINT ["/main"]
