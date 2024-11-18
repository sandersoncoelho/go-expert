FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

FROM scratch
COPY --from=builder /app/cmd/ordersystem/.env .
COPY --from=builder /app/server .
CMD ["./server"]