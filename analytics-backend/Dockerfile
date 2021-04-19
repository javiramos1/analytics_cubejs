FROM golang:1.14-alpine3.11 AS builder

RUN apk update && \
    apk add --no-cache --update alpine-sdk git

ARG CGO_ENABLED=0

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -o dist/analytics-backend -v -ldflags "-w -s" .

FROM gcr.io/distroless/static

COPY --from=builder /app/dist/analytics-backend /usr/local/bin/analytics-backend
COPY --from=builder /app/.env /etc/.env

ENTRYPOINT ["/usr/local/bin/analytics-backend"]