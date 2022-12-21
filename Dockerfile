FROM golang:1.19.4-alpine AS builder

WORKDIR /usr/src/app
COPY . .

#RUN go build -o /go/bin/cep-retriever cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /cep-retriever cmd/main.go

FROM scratch
COPY --from=builder etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /cep-retriever /cep-retriever

EXPOSE 80
EXPOSE 8080
ENTRYPOINT [ "/cep-retriever" ]

