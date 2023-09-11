FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o gores

FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/gores /

EXPOSE 6379

CMD ["/gores"]