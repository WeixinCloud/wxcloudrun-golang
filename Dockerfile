# 选择基础镜像。如需更换，请到(https://hub.docker.com/_/golang?tab=tags) 选择后替换后缀。
FROM golang:1.17.1-alpine3.14 as builder

# Create app directory
RUN mkdir /app

# Add file to /app/
ADD . /app/

# Build the binary
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Run service on container startup
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]
