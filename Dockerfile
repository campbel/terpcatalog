FROM node:alpine AS nodebuilder

COPY admin/app /src/admin/app
WORKDIR /src/admin/app
RUN npm install && npm run build  

COPY catalog/app /src/catalog/app/
WORKDIR /src/catalog/app
RUN npm install && npm run build

FROM golang:alpine AS builder

COPY go.mod go.sum /src/terpcatalog/
WORKDIR /src/terpcatalog
RUN go mod download

COPY --from=nodebuilder /src/admin/app/dist/ /src/terpcatalog/admin/app/dist/
COPY --from=nodebuilder /src/catalog/app/dist/ /src/terpcatalog/catalog/app/dist/
COPY . /src/terpcatalog
RUN CGO_ENABLED=0 GOOS=linux go build -o /terpcatalog .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /terpcatalog /terpcatalog
ENTRYPOINT ["/terpcatalog"]