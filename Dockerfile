FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o exoplannet-app main.go

EXPOSE 8080

ENTRYPOINT [ "./exoplannet-app" ]
