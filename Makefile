GO=go
NAME := ttt
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'
LECTURES_JSON := $(shell cat data/lectures.json | tr -d ' \n' | sed 's/"/\\\\"/g')
COURSES_JSON := $(shell cat data/courses.json | tr -d ' \n' | sed 's/"/\\\\"/g')

all: test build wasm

update_version:
	@for i in README.md; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done

	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' $(NAME).go > a
	@mv a $(NAME).go
	@echo "Replace version to \"${VERSION}\""

update_wasm_json:
	@sed -e 's/const LecturesJSON = .*/const LecturesJSON = "$(LECTURES_JSON)"/g' \
	     -e 's/const CoursesJSON = .*/const CoursesJSON = "$(COURSES_JSON)"/g' standalonedatastore.go > a
	@mv a standalonedatastore.go

setup: update_version update_wasm_json
	git submodule update --init

test: build setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v wasm)

build: setup
	cd cmd/$(NAME) ; $(GO) build -o "../../$(NAME)" -v

wasm: setup
	(cd cmd/wasm ; GOOS=js GOARCH=wasm $(GO) build -o ../../web/main.wasm)

lint: setup
	$(GO) vet $$(go list ./... | grep -v wasm)
	golint cmd/ttt cmd/wasm .
	gosimple 

dist: build
	mkdir -p $(DIST)
	cp -r $(NAME) README.md LICENSE data $(DIST)
	zip -r $(DIST).zip $(DIST)/*

install: test build
	$(GO) install $(LDFLAGS)

clean:
	$(GO) clean
	rm -rf $(NAME) web/main.wasm coverage.out $(DIST) $(DIST).zip
