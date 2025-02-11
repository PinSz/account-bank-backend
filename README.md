# account-bank-backend
assignment test backend

install go 1.23.5
go install github.com/gin-gonic/gin@latest
go install github.com/air-verse/air@v1.52.2

docker-compose up -d --build
docker-compose up -d

go run main.go

http://localhost:4000/api/account-balances
http://localhost:4000/api/account-details
