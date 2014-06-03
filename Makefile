BUILD_OS   = $(shell go env GOOS)
BUILD_ARCH = $(shell go env GOARCH)

build-libwebrtc:
	build/$(BUILD_OS)_$(BUILD_ARCH)/bin/build

run-example:
	go run _examples/simple/main.go

