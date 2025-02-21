# Stage 1: Gerar documentação Swagger
FROM golang:1.24-alpine AS swagger
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    export PATH=$PATH:$(go env GOPATH)/bin && \
    swag init -g cmd/main.go 

# Stage 2: Build da aplicação
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
COPY --from=swagger /app/docs ./docs
# Instalar dependências de build
RUN apk add --no-cache gcc musl-dev
# Build da aplicação
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Stage 3: Imagem final
FROM alpine:latest
WORKDIR /app
# Instalar SQLite
RUN apk add --no-cache sqlite
# Copiar o binário compilado
COPY --from=builder /app/main .
# Copiar o arquivo .env
COPY .env .
# Criar diretório para o banco de dados
RUN mkdir -p /app/data

# Export a porta do container
EXPOSE ${CONTAINER_PORT}

# Comando para executar a aplicação
CMD ["./main"] 