build : .mod
	gofmt -w ./
	go build

.mod :
	go mod download
	touch .mod

mod :
	go mod download

.PHONY : clean
clean:
	rm ./outliner
	rm .mod