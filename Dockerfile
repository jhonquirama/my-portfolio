# Etapa de build
FROM golang:1.22.1-alpine AS builder

# Instala git por si usas módulos externos
RUN apk add --no-cache git

WORKDIR /app

# Copiamos los archivos de dependencia primero (para caché más eficiente)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Ahora copiamos el resto del código
COPY . .

# Compilamos el binario estático para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Etapa final: imagen liviana
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8007

# Comando por defecto
CMD ["./main"]