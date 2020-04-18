FROM golang:latest as builder

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/jakewright/drawbridge
COPY . .

RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go install .

FROM alpine:latest as prod
RUN mkdir /config
WORKDIR /root/
COPY --from=builder /go/bin/drawbridge .
CMD ["./drawbridge", "/config/config.yaml"]
