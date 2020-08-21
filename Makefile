COMMIT_REVISION := $(shell git log --pretty=%h -1)
REVISION_FLAG := "-X github.com/lintingzhen/commitizen-go/cmd.revision=${COMMIT_REVISION}"
TARGET := commitizen-go
GOFILES := $(wildcard *.go) $(wildcard cmd/*.go) $(wildcard git/*.go) $(wildcard commit/*.go)

ifeq ($(OS),Windows_NT)
	GOOS := windows
	COPY := copy
else
	COPY := cp
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS := linux
	else ifeq ($(UNAME_S),Darwin)
		GOOS := darwin
	endif
endif

GIT_EXEC_PATH := $(shell git --exec-path)

all: ${TARGET}
install: 
	$(COPY) commitizen-go $(GIT_EXEC_PATH)/git-cz
clean:
	rm -rf ${TARGET}
    

commitizen-go: $(GOFILES)
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -o $@ -ldflags ${REVISION_FLAG}

.PHONY: all install clean 
