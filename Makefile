COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
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

WORKDIR=/go/src/github.com/abice/gencheck
circleci: ci-docker
	if [ ! -d coverage ]; then mkdir coverage; fi
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "glide install \
	&& go test -v ./generator -race -cover -coverprofile=$(COVERAGEDIR)/generator.coverprofile \
	&& go test -v ./ -race -cover -coverprofile=$(COVERAGEDIR)/gencheck.coverprofile \
	&& go test -v ./internal/example -race -cover -coverprofile=$(COVERAGEDIR)/example.coverprofile"

circlecover:
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "go tool cover -html=$(COVERAGEDIR)/generator.coverprofile -o $(COVERAGEDIR)/generator.html \
	&& go tool cover -html=$(COVERAGEDIR)/gencheck.coverprofile -o $(COVERAGEDIR)/gencheck.html \
	&& go tool cover -html=$(COVERAGEDIR)/gencheck.coverprofile -o $(COVERAGEDIR)/example.html"

circlecoveralls:
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile \
		&& goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)"

ci-test:
	docker run --rm --workdir=$(WORKDIR) -v $$(pwd):$(WORKDIR) --entrypoint=/bin/sh gencheck-builder -c "glide install && go build -v -o bin/gencheck ./gencheck"

	make test
	post:
	- cd "$SRC_DIR" && make cover
	- cd "$SRC_DIR" && make coveralls

phony: clean tc build
