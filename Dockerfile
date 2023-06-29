FROM golang:1.18.8-alpine3.16 AS BuildStage
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o promjson -ldflags '-extldflags "-static"' ./cmd/main.go

FROM amd64/alpine:3.16
WORKDIR /opt
COPY --from=BuildStage /build/promjson .
ENTRYPOINT ["./promjson"]