FROM golang:1.16 as builder

WORKDIR /build

COPY . .
USER root
RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn,direct
RUN export GOPRIVATE=gitee.com
RUN go mod download
RUN build -o main main.go

FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]