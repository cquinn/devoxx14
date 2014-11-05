EXECUTABLE := webhellogo

UNIX_EXECUTABLES := \
	darwin/amd64/$(EXECUTABLE) \
	linux/amd64/$(EXECUTABLE)
WIN_EXECUTABLES := \
	windows/amd64/$(EXECUTABLE).exe

ALL_EXECUTABLES := $(UNIX_EXECUTABLES:%=bin/%) $(WIN_EXECUTABLES:%=bin/%)

DOCKER_IMAGE := cquinn/$(EXECUTABLE)
LAST_TAG := latest

GOFILES := $(shell git ls-files | egrep \.go$ | egrep -v '^(cmd|Godeps)/' | egrep -v _test.go$)
PACKAGES := .

all: $(EXECUTABLE) $(ALL_EXECUTABLES)

build: clean $(ALL_EXECUTABLES)

clean:
	rm -rf ./bin
	rm -f $(EXECUTABLE)

# Build binaries
$(EXECUTABLE): $(GOFILES)
	go build -o "$@"
bin/darwin/amd64/$(EXECUTABLE): $(GOFILES)
	GOARCH=amd64 GOOS=darwin go build -o "$@"
bin/linux/amd64/$(EXECUTABLE): $(GOFILES)
	GOARCH=amd64 GOOS=linux go build -o "$@"
bin/windows/amd64/$(EXECUTABLE).exe: $(GOFILES)
	GOARCH=amd64 GOOS=windows go build -o "$@"

docker_image: all
	docker build -t $(DOCKER_IMAGE):$(LAST_TAG) .

docker_release: docker_image
	docker tag $(DOCKER_IMAGE):$(LAST_TAG) $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE):$(LAST_TAG)
	docker push $(DOCKER_IMAGE):latest


.PHONY: all build clean docker_image docker_release
