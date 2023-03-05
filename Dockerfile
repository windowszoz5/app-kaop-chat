FROM golang:1.17 as builder

WORKDIR /build

COPY . .
RUN env -w GO111MODULE=on
RUN env -w GOPROXY=https://goproxy.io,direct
RUN env -w GOPRIVATE=gitee.com
RUN go mod tidy
RUN build -o main main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]