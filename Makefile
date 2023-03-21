build:
	go build -v ./...

test:
	go test -v ./...

lint:
	golangci-lint run

bench:
	go test -benchmem -bench .

doc:
	godoc -http=:6060 -play -index

cover:
	#go tool cover -func=coverage.out
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...