FROM golang:latest

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/jakewright/drawbridge
COPY . .

RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go install .

FROM alpine:latest
RUN mkdir /config
WORKDIR /root/
COPY --from=0 /go/bin/drawbridge .
CMD ["./drawbridge", "/config/config.yaml"]
