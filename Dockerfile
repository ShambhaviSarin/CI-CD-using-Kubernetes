FROM golang:1.20-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w' -o /out/helloworld ./

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/helloworld /usr/local/bin/helloworld
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/helloworld"]