# base go image for compile
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./cakeStoreApp .

RUN chmod +x /app/cakeStoreApp

# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/cakeStoreApp /app

CMD [ "/app/cakeStoreApp" ]