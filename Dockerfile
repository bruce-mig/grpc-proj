FROM golang:1.21.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /greeter

EXPOSE 8080

CMD ["/greeter"]