.PHONY: build
build: clean
	go run .

.PHONY: clean
clean:
	rm -rf _site

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

