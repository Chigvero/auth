FROM golang:1.23-alpine3.20 as builder

WORKDIR /github.com/Chigvero/source
COPY . /github.com/Chigvero/source

RUN go mod tidy
RUN go mod download
RUN go build -o ./bin/authServer cmd/authServer/main.go

FROM alpine:3.20
WORKDIR /root/
COPY --from=builder  /github.com/Chigvero/source/bin/authServer .
COPY --from=builder /github.com/Chigvero/source/prod.env .
COPY --from=builder /github.com/Chigvero/source/prod_entrypoint.sh .
RUN chmod +x prod_entrypoint.sh
ENTRYPOINT ["sh","prod_entrypoint.sh"]
