FROM golang:1.7.3


RUN go get -u github.com/jteeuwen/go-bindata/... \
    && go get github.com/Masterminds/glide \
    && go get golang.org/x/tools/cmd/cover \
    && go get github.com/mattn/goveralls \
    && go get github.com/modocache/gover
