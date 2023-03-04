FROM golang:1.17 as builder

WORKDIR /build

COPY . .
USER root
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN gi env -w GOPRIVATE=gitee.com
RUN go mod tidy
RUN build -o main main.go



FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]