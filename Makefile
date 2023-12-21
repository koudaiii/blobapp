DARWIN_TARGET_ENV=GOOS=darwin GOARCH=arm64
LINUX_TARGET_ENV=GOOS=linux GOARCH=amd64
APP_NAME=blobapp
BUILD=go build
DOCKER_BUILD=docker build
DOCKER_BUILD_OPTS=--no-cache
DOCKER_RMI=docker rmi -f
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GITHUB_USER=koudaiii
TAG=$(GITHUB_USER)/$(APP_NAME):$(GIT_COMMIT)

.PHONY: build
build:
	CGO_ENABLED=0 $(LINUX_TARGET_ENV)  $(BUILD) -o bin/$(APP_NAME) -ldflags "-s -w"


.PHONY: run
run:
	go run main.go

.PHONY: docker_image
docker_image: clean build
	$(DOCKER_BUILD) -t $(TAG) . $(DOCKER_BUILD_OPTS)

.PHONY: clean
clean:
	$(DOCKER_RMI) -f $(TAG)
