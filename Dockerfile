FROM golang:1.22-alpine
# Enable CGO for go-sqlite3
ENV CGO_ENABLED=1
# Install required tools for cgo
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /products

EXPOSE 8081

CMD [ "/products" ]