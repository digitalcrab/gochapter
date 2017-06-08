NAME=homework2
SOURCE=cmd/server/main.go
DEPEND=github.com/Masterminds/glide

all: clean depend build

clean:
	rm -rf build/

depend:
	go get -u -v $(DEPEND)
	glide install

build:
	go build -o build/$(NAME) $(SOURCE)
