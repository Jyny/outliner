EXECUTABLE=outliner
WINDOWS=$(EXECUTABLE)_windows.exe
LINUX=$(EXECUTABLE)_linux
DARWIN=$(EXECUTABLE)_darwin
VERSION=$(shell cat VERSION)
XPKG="github.com/jyny/outliner/pkg/cmd"

all : build shell_completion
	cp ./build/outliner_$(shell go env GOOS) ./outliner
	@echo build version: $(VERSION)

shell_completion :
	mkdir -p build
	go run completion/gen.go

build : mod embedded fmt
	env GOOS=linux GOARCH=amd64 go build -o ./build/$(LINUX) -ldflags="-X $(XPKG).version=$(VERSION)"  .
	env GOOS=darwin GOARCH=amd64 go build -o ./build/$(DARWIN) -ldflags="-X $(XPKG).version=$(VERSION)"  .
	env GOOS=windows GOARCH=amd64 go build -o ./build/$(WINDOWS) -ldflags="-X $(XPKG).version=$(VERSION)"  .

embedded :
	@pushd pkg/agent > /dev/null && \
 	go run gen/gen.go \
	&& popd > /dev/null

fmt :
	gofmt -w ./

mod :
	go mod download

.PHONY : clean
clean:
	rm -r ./build
	rm outliner