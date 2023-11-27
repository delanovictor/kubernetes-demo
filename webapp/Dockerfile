FROM golang:1.16-alpine

WORKDIR /app

#Instala dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

#Copia arquivos fonte
COPY main.go ./
COPY database ./database
COPY template ./template

#Gera executável
RUN go build -o /webapp

EXPOSE 8000

CMD [ "/webapp" ]