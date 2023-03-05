FROM golang:1.17 as builder

WORKDIR /build

ADD . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GOPRIVATE=gitee.com
RUN go mod tidy
RUN go env -w CGO_ENABLED=0
RUN go build -o main main.go
RUN ls

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]