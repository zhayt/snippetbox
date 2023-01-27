run:
	find ./cmd/web/ -not -name "*_test.go" | xargs -I % go run %

test-humanData:
	go test -v -run='^TestHumanDate/UTC|CET' ./cmd/web

test-ping:
	go test -v -run="^TestPing" ./cmd/web

test-middleware:
	 go test -v -run="^TestMiddleware" ./cmd/web

all-tests:
	go test -v ./...