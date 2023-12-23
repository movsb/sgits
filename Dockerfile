FROM golang:alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build

FROM alpine:3.18
WORKDIR /workspace
RUN apk update && apk add git-daemon && rm -rf /var/cache/apk
COPY --from=builder /app/sgits .
CMD ["./sgits"]
