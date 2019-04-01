COMMIT_REVISION := $(shell git log --pretty=%h -1)
REVISION_FLAG := "-X main.revision=${COMMIT_REVISION}"
TARGET := commitizen

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
	$(COPY) commitizen $(GIT_EXEC_PATH)/git-cz
clean:
	rm -rf ${TARGET}
    

commitizen: commitizen.go arguments.go answers.go git.go
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -o $@ -ldflags ${REVISION_FLAG}

.PHONY: all install clean 
