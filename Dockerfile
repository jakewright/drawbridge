FROM golang:latest
COPY . /go/src/drawbridge

# Install Glide
RUN curl https://glide.sh/get | sh

WORKDIR /go/src/drawbridge

RUN glide install
RUN go build -o bin/drawbridge

# Make the app run when the container is started
CMD ["/go/src/drawbridge/bin/drawbridge"]
