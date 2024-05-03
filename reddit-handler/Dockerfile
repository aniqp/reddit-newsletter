FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY . /app
RUN go mod download

CMD ["go", "run", "./cmd/main.go"]
