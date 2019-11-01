GO=go
NAME := ziraffe
VERSION := 1.0.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)'
	-X 'main.revision=$(REVISION)'
LECTURES_JSON := $(shell cat data/lectures.json | tr -d ' \n' | sed 's/"/\\\\"/g')
COURSES_JSON := $(shell cat data/courses.json | tr -d ' \n' | sed 's/"/\\\\"/g')

all: test build wasm

update_version:
	@for i in README.md; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done

	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' ziraffe.go > a
	@mv a ziraffe.go
	@echo "Replace version to \"${VERSION}\""

update_wasm_json:
	@sed -e 's/const LECTURES_JSON = .*/const LECTURES_JSON = "$(LECTURES_JSON)"/g' \
	     -e 's/const COURSES_JSON = .*/const COURSES_JSON = "$(COURSES_JSON)"/g' cmd/wasm/wasmdatastore.go > a
	@mv a cmd/wasm/wasmdatastore.go

setup: update_version update_wasm_json
	git submodule update --init

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v wasm)

build: setup
	cd cmd/ziraffe    ; go build -o "ziraffe" -v

wasm: setup
	(cd cmd/wasm ; GOOS=js GOARCH=wasm go build -o main.wasm && cp main.wasm ../../web)

lint: setup
	$(GO) vet $$(go list ./... | grep -v vendor)
	for pkg in $$(go list ./... | grep -v vendor); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

install: test build
	$(GO) install $(LDFLAGS)

clean:
	$(GO) clean
	rm -rf $(NAME)
