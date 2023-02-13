.PHONY: be-dev
be-dev:
	go run main.go server

.PHONY: ent-gen
ent-gen:
	go generate ./...

.PHONY: be-build
be-build:
	go build -o build/verinotes

.PHONY: build
build: ent-gen be-build

.PHONY: run
run:
	./build/verinotes server

