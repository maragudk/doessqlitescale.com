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
copy-assets: copy-sqlite
	mkdir -p _site
	cp `go env GOROOT`/misc/wasm/wasm_exec.js _site/
	cp public/* _site/

.PHONY: copy-sqlite
copy-sqlite: sqlite.zip
	mkdir -p _site
	unzip sqlite.zip
	mv sqlite-wasm-3400000/jswasm/sqlite3.js _site
	mv sqlite-wasm-3400000/jswasm/sqlite3.wasm _site
	rm -rf sqlite-wasm-3400000

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: lint
lint:
	golangci-lint run

sqlite.zip:
	curl -sLO https://sqlite.org/2022/sqlite-wasm-3400000.zip
	mv sqlite-wasm-3400000.zip sqlite.zip

.PHONY: start
start: build
	go run ./cmd/server

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

