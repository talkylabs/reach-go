FROM golang:1.19

RUN mkdir /talkylabs
WORKDIR /talkylabs

COPY client ./client
COPY rest ./rest
COPY reach.go .
COPY reach_test.go .

# Fetch dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download
