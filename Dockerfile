FROM golang:latest AS builder
RUN mkdir /minit
ADD . /minit
WORKDIR /minit
RUN CGO_ENABLED=0 GOOS=linux go build -o minit -ldflags "-s -w" cmd/minit/main.go

FROM alpine:latest AS production
COPY --from=builder /minit .
CMD ["./minit"]
EXPOSE 8080