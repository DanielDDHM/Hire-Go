# Usar uma imagem base do Go
FROM golang:1.23-alpine

# Configurar o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o arquivo go.mod e go.sum primeiro para aproveitar o cache de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod tidy

# Copiar todo o restante do código para o container
COPY . .

# Compilar o binário
RUN go build -o main cmd/main.go

# Expor a porta em que a aplicação roda
EXPOSE 4000

# Comando para rodar a aplicação
CMD ["./main"]
