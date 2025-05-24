FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /products

EXPOSE 8081

CMD [ "/products" ]