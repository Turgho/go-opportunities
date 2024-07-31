# Usar a imagem oficial do Golang como base
FROM golang:1.22 AS build

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /api/app

# Copiar os arquivos do projeto para o diretório de trabalho
COPY go.mod .
COPY go.sum .

# Baixar as dependências
RUN go mod download

# Copiar o restante dos arquivos do projeto
COPY . .

# Compilar o binário
RUN go build -o /app cmd/main.go

# Cópia imagem do debian para executar arquivo
FROM gcr.io/distroless/base-debian12

WORKDIR /

# Copiar os arquivos do build
COPY --from=build app app

# Expor a porta que a aplicação usará
EXPOSE 8000

# Define usuário como nonroot
USER nonroot:nonroot

# Comando para rodar a aplicação
ENTRYPOINT [ "./app" ]