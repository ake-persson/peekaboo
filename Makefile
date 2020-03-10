MAJOR=$(shell cat MAJOR)
MINOR=$(shell cat MINOR)
PATCH=$(shell cat PATCH)
VERSION=$(MAJOR).$(MINOR).$(PATCH)

all:	build

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/fzipp/gocyclo
	go get -u ./...

clean:
	rm -f coverage.out client/client agent/agent catalog/catalog

fmt:
#	gofmt -w system/*.go
	for f in $$(find pb -type f); do \
		prototool format -w $$f; \
	done


test:
	golint -set_exit_status ./...
	go vet ./...
	errcheck ./...
	misspell ./...
	ineffassign ./...
	gocyclo -over 15 ./...
#	go test -v -covermode=count

coverage:
#       go test -v -covermode=count -coverprofile=coverage.out
#       go tool cover -html=coverage.out

pbgen:
	for f in $$(find pb -type f -name \*.proto); do \
		protoc -I . --go_out=plugins=grpc:$$GOPATH/src $$f ;\
	done
	for f in $$(find pkg -type f -name \*.go); do \
		sed -i "" -e "s/,omitempty//g" $$f ;\
	done

build:
	cd client && go build -ldflags "-X main.version=$(VERSION)"
	cd agent && go build -ldflags "-X main.version=$(VERSION)"
	cd catalog && go build -ldflags "-X main.version=$(VERSION)"

linux:
	cd client && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)"
	cd agent && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)"
	cd catalog && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)"

major:
	echo $$(( $(MAJOR) + 1 )) >MAJOR
	echo 1 >MINOR
	echo 1 >PATCH

minor:
	echo $$(( $(MINOR) + 1 )) >MINOR
	echo 1 >PATCH

patch:
	echo $$(( $(PATCH) + 1 )) >PATCH

.PHONY: deps clean fmt test coverage pbfmt pbgen major minor patch
