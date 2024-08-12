FROM golang:alpine

WORKDIR /app

COPY . /app

ADD go.mod .

RUN go build -o main .

ENTRYPOINT ["./main"]
