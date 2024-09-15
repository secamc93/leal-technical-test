# Etapa de construcción: Compila el binario estático usando Go.
FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . /app
WORKDIR /app/cmd
RUN go build -o /app/leal-technical-test .

# Etapa final: Preparar la imagen de ejecución.
FROM alpine:3.20 AS runner


WORKDIR /app
COPY --from=builder /app/leal-technical-test /app/leal-technical-test
COPY .env /app/.env

EXPOSE 60000

CMD ["/app/leal-technical-test"]


