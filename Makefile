BINARY_NAME=purrsom_watch
GOOS=windows
GOARCH=amd64

ifeq ($(GOOS),windows)
	FILE_EXT=.exe
else
	FILE_EXT=
endif

build:
	go build -C cmd/watch -o ../../bin/$(BINARY_NAME)$(FILE_EXT)
	@echo Successfully built: bin/$(BINARY_NAME)$(FILE_EXT)

.PHONY: build