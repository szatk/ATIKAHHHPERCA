
FROM golang:1.17.3-alpine3.14 AS builder

WORKDIR /atikahperca
COPY ./ ./
RUN go mod download

RUN go build -o main


#2
FROM alpine:3.14
WORKDIR /atikahperca
COPY --from=builder /atikahperca/main .
COPY .env /atikahperca


EXPOSE 8080

CMD [ "./main" ]
