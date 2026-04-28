test:
	go test ./...

run:
	go run ./cmd/server

ci:
	gofmt -w cmd/server/main.go internal/calc/calc.go internal/calc/calc_test.go tests/integration/server_test.go
	go vet ./...
	go test -race ./...
