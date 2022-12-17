#Build stage
FROM golang:1.18-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY static/ ./static/
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x start.sh
RUN chmod +x wait-for.sh
COPY db/migration ./db/migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]