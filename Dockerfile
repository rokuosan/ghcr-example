FROM golang:1.24.0

WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN go build -o sample-app .

CMD ["./sample-app"]
