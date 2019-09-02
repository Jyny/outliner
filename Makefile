build : .mod
	gofmt -w ./
	go build

.mod :
	go mod download
	touch .mod

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