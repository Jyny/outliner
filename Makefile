EXECUTABLE=outliner
WINDOWS=$(EXECUTABLE)_windows.exe
LINUX=$(EXECUTABLE)_linux
DARWIN=$(EXECUTABLE)_darwin
VERSION=$(shell git describe --tags --always --long --dirty)
XPKG="github.com/jyny/outliner/pkg/cmd"

release : .mod pregen fmt $(LINUX) $(DARWIN) $(WINDOWS)
	@echo version: $(VERSION)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o ./build/$(WINDOWS) -ldflags="-X $(XPKG).version=$(VERSION)"  .

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o ./build/$(LINUX) -ldflags="-X $(XPKG).version=$(VERSION)"  .

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -o ./build/$(DARWIN) -ldflags="-X $(XPKG).version=$(VERSION)"  .

.mod :
	go mod download
	touch .mod

.PHONY : pregen
pregen :
	@pushd pkg/agent > /dev/null && \
	go run gen/gen.go \
	&& popd > /dev/null

.PHONY : mod
mod :
	go mod download

.PHONY : fmt
fmt :
	gofmt -w ./

.PHONY : clean
clean:
	rm -r ./build
	rm .mod