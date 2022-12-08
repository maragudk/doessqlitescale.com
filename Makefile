.PHONY: build
build: clean copy-assets build-wasm
	mkdir -p _site
	go run ./cmd/build

.PHONY:build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o _site/app.wasm ./cmd/wasm

.PHONY: clean
clean:
	rm -rf _site

.PHONY: copy-assets
copy-assets:
	mkdir -p _site
	cp `go env GOROOT`/misc/wasm/wasm_exec.js _site/
	cp public/* _site/

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: start
start: build
	go run ./cmd/server

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

