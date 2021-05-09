.PHONY: pre-commit

pre-commit:
	go mod tidy
	go mod vendor
	go vet
	go fmt ./...
