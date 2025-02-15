# account-bank-backend
assignment test backend
##
git clone https://github.com/your-repo/account-bank-backend.git
##
cd account-bank-backend
##

install go 1.23.5
##
go install github.com/gin-gonic/gin@latest
##
go mod tidy
##
go install github.com/air-verse/air@v1.52.2
##

--------------------------------------------
".env"
--------------------------------------------
HTTP_PORT=8080
--------------------------------------------
GLOBAL_ENDPOINT=account-bank-backend
--------------------------------------------
DATABASE_URL=postgresql://postgres:postgres@db:5432/account-bank-backend?sslmode=disable
-------------------------------------------

//run
##
docker-compose up -d --build
##
docker-compose up -d

-------------------------------------------
ในการเรียกใช้งาน ทดสอบผ่าน postman
##
post ==> http://localhost:4000/api/account/details
response --> {
    "data": [
        {
            "accountId": "000020ece1a211ef95a30242ac180002",
            "accountNumber": "568-2-62729",
            "color": "#24c875",
            "isMainAccount": true,
            "progress": 42,
            "amount": 86048.06,
            "flags": [
                {
                    "flagType": "system",
                    "flagValue": "Flag4",
                    "createdAt": "2025-02-02T20:12:17"
                },
                {
                    "flagType": "system",
                    "flagValue": "Flag5",
                    "createdAt": "2025-02-02T20:12:17"
                }
            ]
        },
        {
            "accountId": "000024eae1a211ef95a30242ac180002",
            "accountNumber": "568-2-94760",
            "color": "#24c875",
            "progress": 82,
            "amount": 93311.34
        }
    ],
    "status": 200
}
##
post ==> http://localhost:4000/api/debit-card/info
response --> {
    "data": {
        "name": "My Debit Card",
        "status": "Active",
        "issuer": "TestLab",
        "number": "6821 5668 7876 2379",
        "color": "#00a1e2",
        "borderColor": "#ffffff"
    },
    "status": 200
}
##
get ==> http://localhost:4000/api/user/000018b0e1a211ef95a30242ac180002
response --> {
    "data": {
        "name": "User_000018b0e1a211ef95a30242ac180002",
        "userId": "000018b0e1a211ef95a30242ac180002"
    },
    "status": 200
}
##
get ==> http://localhost:4000/api/banner/000018b0e1a211ef95a30242ac180002
response --> {
    "banner": {
        "bannerId": "000018cfe1a211ef95a30242ac180002",
        "userId": "000018b0e1a211ef95a30242ac180002",
        "title": "Want some money?",
        "description": "You can start applying",
        "image": "https://dummyimage.com/54x54/999/fff"
    },
    "status": 200
}
##
get ==> http://localhost:4000/api/transactions/000018b0e1a211ef95a30242ac180002
response --> {
    "data": {
        "transactions": [
            {
                "name": "Transaction_135017",
                "image": "https://dummyimage.com/54x54/999/fff",
                "isBank": false
            }
        ]
    },
    "status": 200
}

-------------------------------------------
# 🙏 Thank You! 🎉
