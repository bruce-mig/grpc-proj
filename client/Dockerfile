FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go mod init github.com/bruce-mig/grpc-proj

COPY . .

RUN go mod tidy

RUN cd client
