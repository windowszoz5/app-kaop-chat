FROM golang:1.16 as builder

WORKDIR /build

COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main main.go

FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]