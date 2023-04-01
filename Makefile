
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