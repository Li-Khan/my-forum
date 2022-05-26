FROM golang:1.18 AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o forum ./cmd/

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./forum"]