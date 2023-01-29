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

cover-test:
	go test -cover ./...

coverprofile-test:
	go test -coverprofile=/tmp/profile.out ./...

tool-cover-test:
	go tool cover -func=/tmp/profile.out

html-cover-test:
	go test -covermode=count -coverprofile=/tmp/profile.out -short ./...
	go tool cover -html=/tmp/profile.out

unhashed-test:
	go test -v -count=1 -short ./...