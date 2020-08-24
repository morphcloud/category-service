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

.PHONY: format
format:
	gofmt -w ./..

.PHONY: docker-build
docker-build:
	docker build -t hzhyvinskyi/morphcloud-order-service:1.0.0 .

.PHONY: docker-run
docker-run:
	docker run -d -p 8081:8081 --name morphcloud-order-service-container hzhyvinskyi/morphcloud-order-service:1.0.0

.PHONY: docker-push
docker-push:
	docker push hzhyvinskyi/morphcloud-order-service:1.0.0

.DEFAULT_GOAL := build-and-run
