FROM golang:latest as builder

WORKDIR /rosary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/rosary

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go/bin/rosary /usr/local/bin

WORKDIR /rosary

ENV ROSARY_ADDR=:6724 ROSARY_BEAT_MS=400 ROSARY_POLL_PERIOD_MS=100
EXPOSE 6724

ENTRYPOINT ["rosary"]