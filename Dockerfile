FROM node:20-alpine AS vue-builder
WORKDIR /vue3
COPY ["web/index.html","web/env.d.ts","web/*.json","web/*.mts","./"]
COPY web/public/ public/
COPY web/src/ src/
RUN npm install --legacy-peer-deps \
    && npm run build 

FROM golang:1.22-alpine AS go-builder
WORKDIR /golang
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
COPY config/ config/
COPY --from=vue-builder /vue3/dist/ web/dist/
COPY --from=go-builder /golang/zender /
ENTRYPOINT ["/zender"]