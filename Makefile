
GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
GOIMPORTS     = $(GOBIN)/goimports

.PHONY: lint
lint: verify-import-alias
	bash hack/verify-staticcheck.sh

.PHONY: verify-import-alias
 verify-import-alias:
	bash hack/verify-import-alias.sh

.PHONY: verify-vendor
verify-vendor:
	bash hack/verify-vendor.sh

.PHONY: verify-all
verify-all: lint verify-vendor

.PHONY: genall
genall:
	cd ./api/ && $(MAKE) genall

.PHONY: build
build:
	go build -o ./bin/hub cmd/main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/hub-linux cmd/main.go

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/hub-darwin cmd/main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/hub-windows cmd/main.go

build-all: build-linux build-darwin build-windows
