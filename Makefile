
DEPOT_TOOLS_REPO="https://chromium.googlesource.com/chromium/tools/depot_tools.git"
LIB_WEBRTC_REPO="http://webrtc.googlecode.com/svn/trunk"
LIBWEBRTC_REVISION=r6226
HOST_ARCH=x64
TARGET_ARCH=x64

export BUILD_OS=$(shell go env GOOS)
export BUILD_ARCH=$(shell go env GOARCH)
export GYP_GENERATORS=ninja
export GYP_DEFINES=host_arch=$(HOST_ARCH) target_arch=$(TARGET_ARCH)
export PATH := $(realpath vendor/depottools):$(PATH)

SRC_H_FILES := $(shell find vendor/libwebrtc/trunk -type f -name '*.h')
DST_H_FILES := $(subst vendor/libwebrtc/trunk/,build/${BUILD_OS}_${BUILD_ARCH}/include/,$(SRC_H_FILES))

SRC_A_FILES := $(shell find vendor/libwebrtc/trunk/out/Release -type f -name '*.a')
DST_A_FILES := $(subst vendor/libwebrtc/trunk/out/Release/,build/${BUILD_OS}_${BUILD_ARCH}/lib/,$(SRC_A_FILES))

build: $(DST_H_FILES) $(DST_A_FILES)

vendor/libwebrtc/trunk/DEPS: vendor/depottools/gclient
	mkdir -p vendor/libwebrtc
	(cd vendor/libwebrtc; gclient config ${LIB_WEBRTC_REPO})
	(cd vendor/libwebrtc; gclient sync -f -n  -D -j1 -r${LIBWEBRTC_REVISION})
	find vendor -type d -name .git -exec bash -c 'echo "* -text" > {}/info/attributes' \;
	(cd vendor/libwebrtc; gclient runhooks -j1)
	(cd vendor/libwebrtc; ninja -C trunk/out/Release)

vendor/libwebrtc/trunk/out/Release/%.a: vendor/libwebrtc/trunk/DEPS vendor/depottools/gclient

vendor/depottools/gclient:
	rm -rf vendor/depottools
	git clone ${DEPOT_TOOLS_REPO} vendor/depottools
	find vendor -type d -name .git -exec bash -c 'echo "* -text" > {}/info/attributes' \;

env:
	echo $(DST_H_FILES)

build/${BUILD_OS}_${BUILD_ARCH}/include/%.h: vendor/libwebrtc/trunk/%.h
	@mkdir -p $(dir $@)
	cp $< $@

build/${BUILD_OS}_${BUILD_ARCH}/lib/%.a: vendor/libwebrtc/trunk/out/Release/%.a
	@mkdir -p $(dir $@)
	cp $< $@


run-example:
	go run _examples/simple/main.go
