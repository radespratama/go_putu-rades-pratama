FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN go build -o /dist

EXPOSE 9000

CMD ["/dist"]