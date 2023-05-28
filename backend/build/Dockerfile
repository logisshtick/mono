FROM golang:1.20.4-alpine3.18 as builder
RUN apk add make

WORKDIR /app
COPY ../ ./
RUN go mod download

COPY . .
RUN make


FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/svlz /app/app

EXPOSE 4433
CMD ["./app"]
