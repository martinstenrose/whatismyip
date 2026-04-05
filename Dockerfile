FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o whatismyip .

FROM alpine:3.21
RUN addgroup -g 1000 app && adduser -D -u 1000 -G app app
WORKDIR /
COPY --from=builder /app/whatismyip /whatismyip
USER app
EXPOSE 8080
ENTRYPOINT [ "/whatismyip" ]
