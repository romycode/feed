run-feeder:
	@ go run ./cmd/feeder/main.go

run-client:
	@ go run ./cmd/client/main.go

test:
	@ go test -count=1 ./internal/ ./internal/platform/storage