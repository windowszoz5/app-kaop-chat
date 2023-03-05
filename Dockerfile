FROM golang:1.17 as builder

WORKDIR /build

ADD . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPRIVATE=gitee.com
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .
RUN ls

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["ls"]
CMD ["./main"]