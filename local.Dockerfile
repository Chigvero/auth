FROM golang:1.23-alpine3.20 AS builder

WORKDIR /github.com/Chigvero/source
COPY . /github.com/Chigvero/source

RUN go mod tidy
RUN go mod download
RUN go build -o ./bin/authServer cmd/authServer/main.go

FROM alpine:3.20
WORKDIR /root/
COPY --from=builder  /github.com/Chigvero/source/bin/authServer .
COPY --from=builder /github.com/Chigvero/source/local.env .
COPY --from=builder /github.com/Chigvero/source/local_entrypoint.sh .
RUN chmod +x local_entrypoint.sh
ENTRYPOINT ["sh","local_entrypoint.sh"]