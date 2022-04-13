cuc: tidy *.go
	go fmt ./...
	go vet ./...
	go build ./...

.PHONY: tidy
tidy: go.*
	go mod tidy

.PHONY: run
run: cuc
	./cuc 
