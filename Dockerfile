FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /project-petshop

EXPOSE 8000

CMD [ "/project-petshop" ]