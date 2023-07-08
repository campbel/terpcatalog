FROM golang:alpine AS builder

COPY go.mod go.sum /src/terpcatalog/
WORKDIR /src/terpcatalog
RUN go mod download

COPY . /src/terpcatalog
RUN CGO_ENABLED=0 GOOS=linux go build -o /terpcatalog .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /terpcatalog /terpcatalog
ENTRYPOINT ["/terpcatalog"]