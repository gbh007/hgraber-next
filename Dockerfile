FROM alpine:3.20

# добавление ssl сертификатов и пакета временых зон
RUN apk update && apk add ca-certificates tzdata

RUN mkdir /app


ARG BINARY_PATH
COPY ${BINARY_PATH} /app/main

RUN chmod +x /app/main

WORKDIR /app

ENTRYPOINT ["/app/main"]