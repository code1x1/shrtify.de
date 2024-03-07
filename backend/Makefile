.PHONY: build clean
clean:
	rm -f ./dist/server

build: clean
	CGO_ENABLED=0 go build -gcflags=all="-N -l" -o dist/server ./src
