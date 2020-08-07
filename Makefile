.PHONY: build
build:
	go build -o ./bin/app ./cmd/server/main.go

.PHONY: run
run:
	./bin/app

.PHONY: build-and-run
build-and-run: build run

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: test-bench
test-bench:
	go test -bench=. ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

format:
	gofmt -w ./..

docker-build:
	docker build -t order-service-image:1.0.0 .

docker-run:
	docker run -d -p 8081:8081 --name order-service-container order-service-image:1.0.0

.DEFAULT_GOAL := build-and-run
