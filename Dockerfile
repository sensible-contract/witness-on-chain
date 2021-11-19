FROM golang:1.15-alpine AS build
ARG GO_OS="linux"
ARG GO_ARCH="amd64"
WORKDIR /build/
COPY . .

# Build binary output
RUN GOPROXY=https://goproxy.cn,direct GOOS=${GO_OS} GOARCH=${GO_ARCH} go build -v -o witnessonchain main.go

FROM alpine:latest
COPY --from=build /build/witnessonchain witnessonchain
ENV LISTEN 0.0.0.0:8000
EXPOSE 8000
CMD ["./witnessonchain"]
