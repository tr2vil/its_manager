DATE := $(shell date +%FT%T%z)
BRANCH := $(shell git branch | cut -d ' ' -f2)
USER := $(shell whoami)
GIT_HASH := $(shell git describe --always HEAD)
GOPATH := $(shell env| grep GOPATH | cut -d '=' -f 2)
GOBIN := $(shell env| grep GOBIN | cut -d '=' -f 2)

CC = go build
TARGET = its_manager
default: amd64
#LDFLAGS='-s -X github.com/prometheus/common/version.Branch=$(BRANCH) \
-X github.com/prometheus/common/version.BuildUser=$(USER) \
-X github.com/prometheus/common/version.BuildDate=$(DATE) \
-X github.com/prometheus/common/version.Version=$(GIT_HASH) \
-X github.com/prometheus/common/version.Revision=$(GIT_HASH)'

$(TARGET) : its_manager.go
	$(CC) -o $(TARGET) sock_stat_exporter.go

freebsd:
	GOOS=freebsd GOARCH=386 $(CC) -o $(TARGET)_$@ $^
amd64:
	GOOS=linux GOARCH=amd64 $(CC) -o $(TARGET)_$@ $^
darwin:
	GOOS=darwin GOARCH=amd64 $(CC) -o $(TARGET)_$@ $^
all: freebsd amd64 darwin
install:
	cp $(TARGET)_* $(GOBIN)
clean:
	rm -f $(TARGET)_*
dist_clean:
	rm -f $(GOBIN)/$(TARGET)_*
