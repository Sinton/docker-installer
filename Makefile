# Go parameters
GOBUILD=go build
GOCLEAN=go clean
GOARCH=amd64
BINARY_NAME=docker-install

export GO111MODULE=on
export CGO_ENABLED=0

all: build

clean:
	$(GOCLEAN)

build: clean
	$(GOBUILD)

release: clean
	GOOS=windows $(GOARCH) $(GOBUILD) && zip Windows-$(GOARCH).zip ./$(BINARY_NAME).exe && rd /s /q ./$(BINARY_NAME).exe
	GOOS=darwin $(GOARCH) $(GOBUILD) && zip MacOS-$(GOARCH).zip ./$(BINARY_NAME) && rm -rf ./$(BINARY_NAME)
	GOOS=linux $(GOARCH) $(GOBUILD) && zip Linux-$(GOARCH).zip ./$(BINARY_NAME) && rm -rf ./$(BINARY_NAME)
