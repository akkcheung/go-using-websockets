FROM golang:1.18-alpine3.16 AS builder
RUN mkdir /build
ADD go.mod go.sum server.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/go-using-websockets /app/ 
COPY websockets.html /app/
WORKDIR /app
CMD ["./go-using-websockets"]
