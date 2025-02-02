FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./cmd/main ./cmd/main.go && chmod +x ./cmd/main

EXPOSE 8080

WORKDIR /app/cmd

RUN mkdir -p log && chmod 777 log

# Создаем папку backend/deploy, если ее нет
RUN mkdir -p /app/backend/deploy

# Принудительно копируем .env внутрь контейнера
COPY deploy/.env /app/backend/deploy/.env

CMD ["./main"]