run:
	go run main.go

test:
	go test ./...

test/cover:
	go test -coverprofile coverage.out ./...
	go tool cover -html=coverage.out
