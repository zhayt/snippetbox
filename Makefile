run:
	find ./cmd/web/ -not -name "*_test.go" | xargs -I % go run %

test:
	go test -v ./cmd/web

all-test:
	go test ./...