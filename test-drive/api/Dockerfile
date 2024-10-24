FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o backend .

EXPOSE 8080

CMD ["./backend"]