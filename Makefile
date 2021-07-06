MAJOR=$(shell cat MAJOR)
MINOR=$(shell cat MINOR)
PATCH=$(shell cat PATCH)
VERSION=$(MAJOR).$(MINOR).$(PATCH)

all:	build

brew-deps:
	brew install go
	brew install protobuf

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/fzipp/gocyclo
#	go get -u ./...
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

#	go install google.golang.org/protobuf/cmd/protoc-gen-go
#go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto


clean:
	rm -f coverage.out peekaboo

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
	protoc -I=. --proto_path=pb --go_out=pkg --go-grpc_out=pkg --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative pb/v1/**/*.proto
	for f in $$(find pkg -type f -name \*.go); do \
		sed -i "" -e "s/,omitempty//g" -e 's!json:"\(.*\)"!json:"\1" csv:"\1" yaml:"\1"!g' $$f ;\
        done


build:
	go build -ldflags "-X main.version=$(VERSION)"

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)"

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)"

major:
	echo $$(( $(MAJOR) + 1 )) >MAJOR
	echo 1 >MINOR
	echo 1 >PATCH

minor:
	echo $$(( $(MINOR) + 1 )) >MINOR
	echo 1 >PATCH

patch:
	echo $$(( $(PATCH) + 1 )) >PATCH

.PHONY: deps clean fmt test coverage pbgen build linux darwin major minor patch
