FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY /templates ./templates
COPY /static ./static
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./shortener

FROM scratch

COPY --from=builder /app/shortener /shortener
COPY --from=builder /app/templates /templates
COPY --from=builder /app/static /static
COPY --from=builder /app/.env /.env

EXPOSE 3579

CMD [ "/shortener" ]