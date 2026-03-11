FROM golang:1.22

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o server ./app/main.go

EXPOSE 8081

CMD ["./server"]