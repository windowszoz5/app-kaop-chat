FROM golang:1.17 as builder

WORKDIR /build

ADD . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPRIVATE=gitee.com
RUN go mod tidy
RUN go build -ldflags="-s -w" -o main main.go
RUN ls

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["ls"]
CMD ["./main"]