FROM golang:1.9.0 AS builder
WORKDIR /go/src/github.com/github.com/YAPPi-incubator/go-app
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

FROM busybox:1.27
COPY --from=builder /go/src/github.com/github.com/YAPPi-incubator/go-app .

EXPOSE 8080
ENTRYPOINT ["/go-app"]