build : .mod pregen
	gofmt -w ./
	go build

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
	rm ./outliner
	rm .mod