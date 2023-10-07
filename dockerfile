FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY /templates ./templates
COPY /static ./static
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /shortener

EXPOSE 3579

CMD [ "/shortener" ]