COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)/coverage
endif

all: generate fmt build test cover install
install-deps:
	glide install
build: generate
	if [ ! -d bin ]; then mkdir bin; fi
	go build -v -o bin/gencheck ./gencheck
fmt:
	gofmt -l -w -s $$(find . -type f -name '*.go' -not -path "./vendor/*")
test: generate gen-test
	if [ ! -d coverage ]; then mkdir coverage; fi
	go test -v ./generator -race -cover -coverprofile=$(COVERAGEDIR)/generator.coverprofile
	go test -v ./ -race -cover -coverprofile=$(COVERAGEDIR)/gencheck.coverprofile
	go test -v ./internal/example -race -cover -coverprofile=$(COVERAGEDIR)/example.coverprofile
cover:
	go tool cover -html=$(COVERAGEDIR)/generator.coverprofile -o $(COVERAGEDIR)/generator.html
	go tool cover -html=$(COVERAGEDIR)/gencheck.coverprofile -o $(COVERAGEDIR)/gencheck.html
	go tool cover -html=$(COVERAGEDIR)/gencheck.coverprofile -o $(COVERAGEDIR)/example.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	go clean
	rm -f bin/gencheck
	rm -rf coverage/

# go-bindata will take all of the template files and create readable assets from them in the executable.
# This way the templates are readable in atom (or another text editor with golang template language support)
generate:
	go-bindata -o generator/assets.go -pkg=generator template/*.tmpl

gen-test: build install
	go generate $$(glide novendor)

install:
	go install ./gencheck

reinstall: build
	go install ./gencheck

ci-docker:
	docker build -t gencheck-builder .

PWD=$(shell pwd)
WORKDIR=/go/src/github.com/abice/gencheck
circleci: ci-docker
	if [ ! -d "$(COVERAGEDIR)" ]; then mkdir -p "$(COVERAGEDIR)"; fi
	docker run --rm --workdir=$(WORKDIR) -v $(PWD):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "glide install \
	&& go test -v ./generator -race -cover -coverprofile=coverage/generator.coverprofile \
	&& go test -v ./ -race -cover -coverprofile=coverage/gencheck.coverprofile \
	&& go test -v ./internal/example -race -cover -coverprofile=coverage/example.coverprofile"

circlecover:
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "go tool cover -html=coverage/generator.coverprofile -o coverage/generator.html \
	&& go tool cover -html=coverage/gencheck.coverprofile -o coverage/gencheck.html \
	&& go tool cover -html=coverage/gencheck.coverprofile -o coverage/example.html"

circlecoveralls:
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "gover coverage coverage/coveralls.coverprofile \
		&& goveralls -coverprofile=coverage/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)"

ci-test:
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "glide install && go build -v -o bin/gencheck ./gencheck"

phony: clean tc build
