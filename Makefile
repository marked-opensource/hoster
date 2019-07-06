run:
	go run main.go

test:
	go test -v -count 1 ./...

fmt:
	go fmt ./...
