GO ?= go

.PHONY: all
all: test lint examples

.PHONY: install
install: examples

.PHONY: examples
examples:
	cd examples; $(GO) install ./...

.PHONY: test
test:
	echo "" > /tmp/coverage.txt
	set -e; for dir in `find . -type f -name "go.mod" | sed 's@/[^/]*$$@@' | sort | uniq`; do ( set -xe; \
	  cd $$dir; \
	  $(GO) test -mod=readonly -v -cover -coverprofile=/tmp/profile.out -covermode=atomic -race ./...; \
	  if [ -f /tmp/profile.out ]; then \
	    cat /tmp/profile.out >> /tmp/coverage.txt; \
	    rm -f /tmp/profile.out; \
	  fi); done
	mv /tmp/coverage.txt .

.PHONY: go-mod-tidy-all
go-mod-tidy-all:
	set -e; for dir in `find . -type f -name "go.mod" | sed 's@/[^/]*$$@@' | sort | uniq`; do ( set -xe; \
	  cd $$dir; \
	  go mod tidy); done

.PHONY: go-mod-update-all
go-mod-update-all:
	set -e; for dir in `find . -type f -name "go.mod" | sed 's@/[^/]*$$@@' | sort | uniq`; do ( set -xe; \
	  cd $$dir; \
	  go get -u); done

.PHONY: lint
lint:
	golangci-lint run --verbose ./...
