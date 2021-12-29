.PHONY: test
test: 
	go test ./... -cover

.PHONY: run
run: 
	go run example/main.go

.PHONY: scripts
scripts:
	cd scripts && pipenv run jupyter notebook