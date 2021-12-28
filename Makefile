.PHONY: test
test: 
	go test ./... -cover

.PHONY: run
run: 
	go run example/main.go