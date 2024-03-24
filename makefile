GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=forge-template

default: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

build-for-all-environments:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux-amd64 -v
	GOOS=linux GOARCH=386 $(GOBUILD) -o $(BINARY_NAME)-linux-386 -v
	GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_NAME)-linux-arm -v
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BINARY_NAME)-linux-arm64 -v
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows-amd64 -v
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(BINARY_NAME)-windows-386 -v
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows-arm -v
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(BINARY_NAME)-windows-arm64 -v
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-darwin-amd64 -v
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BINARY_NAME)-darwin-arm64 -v

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_NAME)-linux-amd64
	rm -rf $(BINARY_NAME)-linux-386
	rm -rf $(BINARY_NAME)-linux-arm
	rm -rf $(BINARY_NAME)-linux-arm64
	rm -rf $(BINARY_NAME)-windows-amd64
	rm -rf $(BINARY_NAME)-windows-386
	rm -rf $(BINARY_NAME)-windows-arm
	rm -rf $(BINARY_NAME)-windows-arm64
	rm -rf $(BINARY_NAME)-darwin-amd64
	rm -rf $(BINARY_NAME)-darwin-arm64