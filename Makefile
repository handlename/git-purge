VERSION=$(shell git describe)

all:
	go build -ldflags "-X main.version $(VERSION)"

install:
	go install -ldflags "-X main.version $(VERSION)"
