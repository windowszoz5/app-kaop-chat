FROM golang:1.16 as builder

WORKDIR /build

COPY . .
RUN go build -o main main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]