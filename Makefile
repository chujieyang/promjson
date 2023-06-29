all:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' cmd/main.go 
	@docker build -t harbor.yusur.tech/telemetry/promjson:1.0 -f Dockerfile .