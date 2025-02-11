FROM golang:1.23.5 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN swag init -g main.go -o docs

RUN go build -o account-bank-backend

EXPOSE 8080

CMD ["/app/account-bank-backend"]