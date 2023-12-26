FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY tri/ ./tri/

RUN GOARCH=amd64 GOOS=linux go build -o appexe cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/appexe .

CMD [ "/app/appexe" ]
