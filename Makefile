.PHONY: fe-install
fe-install:
	cd ui && \
	npm install

.PHONY: fe-dev
fe-dev:
	cd ui && \
	npm run dev

.PHONY: fe-build
fe-build:
	cd ui && \
	npm install && \
	npm run build

.PHONY: be-dev
be-dev:
	go run main.go server

.PHONY: ent-gen
ent-gen:
	go generate ./...

.PHONY: be-build
be-build:
	go build -o build/verinotes --tags ui

.PHONY: build
build: fe-build ent-gen be-build

.PHONY: run
run:
	./build/verinotes server

.PHONY: lint
lint:
	golangci-lint run

#
# Security Scans
#

.PHONY: scan
scan: osv-go gosec osv-ui govulncheck-scan

.PHONY: osv-go
osv-go:
	osv-scanner -L go.mod --format json | jq '.results[].packages[].vulnerabilities[]'
.PHONY: osv-ui
osv-ui:
	osv-scanner -L ui/package-lock.json --format json | jq '.results[].packages[].vulnerabilities[]'
.PHONY: gosec
gosec:
	gosec -fmt=json ./...

.PHONY: govulncheck-scan
govulncheck-scan:
	govulncheck ./...

