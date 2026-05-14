FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o notex .

FROM alpine:3.19

RUN apk add --no-cache sqlite-libs ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/notex .
COPY --from=builder /app/backend/frontend ./backend/frontend
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh notex && mkdir -p /data

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "./entrypoint.sh"]
