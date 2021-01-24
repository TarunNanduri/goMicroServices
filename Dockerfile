FROM golang:1.15.3 as builder

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 go build -o microServices main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/microServices .

CMD /app/microServices