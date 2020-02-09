FROM golang:alpine as builder

WORKDIR /builder
COPY ./ /builder

RUN apk --no-cache add ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fritzbox-cloudflare cmd/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /builder/fritzbox-cloudflare /fritzbox-cloudflare

EXPOSE 6221

ENTRYPOINT [ "/fritzbox-cloudflare" ]
