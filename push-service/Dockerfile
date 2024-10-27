# Устанавливаем базовый образ
FROM golang:1.22.4-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем зависимости и устанавливаем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка исполняемого файла
RUN go build -o /push-service ./cmd/push/main.go

# Создаем минимальный образ для запуска
FROM alpine:latest
WORKDIR /root/

# Копируем скомпилированное приложение из builder'а
COPY --from=builder /push-service .

# Экспорт порта
EXPOSE 50054

# Команда для запуска приложения
CMD ["./push-service"]
