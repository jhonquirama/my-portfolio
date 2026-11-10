# Etapa de build
FROM golang:1.22.1-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Etapa final
FROM alpine:latest

# Instala AWS CLI (opcional si lo necesitas en el entrypoint)
RUN apk --no-cache add ca-certificates curl unzip bash \
 && curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
 && unzip awscliv2.zip && ./aws/install \
 && rm -rf awscliv2.zip aws

WORKDIR /app

ENV APPNAME=main

COPY --from=builder /app/main .
COPY --from=builder /app/entrypoint.sh .

RUN chmod +x /app/entrypoint.sh

EXPOSE 8007

CMD ["/app/entrypoint.sh"]
