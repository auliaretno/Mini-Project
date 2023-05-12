FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /project-petshop

EXPOSE 3306

CMD [ "/project-petshop" ]