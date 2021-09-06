FROM golang:1.16 as builder

WORKDIR /

COPY . .
RUN go mod tidy
RUN go build -o app .

CMD ["./app"]


FROM virtuan/alpine-3
RUN apk add --no-cache tzdata


WORKDIR /root

COPY --from=builder /docs .
COPY --from=builder /app .
COPY --from=builder /start.sh .

EXPOSE $PORT

CMD ["sh", "start.sh"]
