FROM golang:1.22-alpine AS builder
WORKDIR /zender
COPY ["main.go", "go.mod", "go.sum", "./"]
COPY pkg/ pkg/
RUN go mod download \
    && go install honnef.co/go/tools/cmd/staticcheck@2024.1 \
    && go vet ./... \
    && staticcheck ./... \
    && go test ./... \
    && go build 

FROM alpine:latest AS release
ENV GIN_MODE=release
COPY certs/ certs/
COPY config/ config/
COPY web/dist web/dist
COPY --from=builder /zender/zender /
ENTRYPOINT ["/zender"]