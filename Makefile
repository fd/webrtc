
DEPOT_TOOLS_REPO="https://chromium.googlesource.com/chromium/tools/depot_tools.git"
LIB_WEBRTC_REPO="http://webrtc.googlecode.com/svn/trunk"
LIBWEBRTC_REVISION=r5459
HOST_ARCH=x64
TARGET_ARCH=x64

export GYP_GENERATORS=ninja
export GYP_DEFINES=host_arch=$(HOST_ARCH) target_arch=$(TARGET_ARCH)
export PATH := $(realpath vendor/depottools):$(PATH)

build: vendor/libwebrtc/out

vendor/libwebrtc:
	mkdir -p vendor/libwebrtc

vendor/libwebrtc/.gclient: vendor/libwebrtc
	(cd vendor/libwebrtc; gclient config ${LIB_WEBRTC_REPO})

vendor/libwebrtc/trunk:
	(cd vendor/libwebrtc; gclient sync -f -n  -D -j1 -r${LIBWEBRTC_REVISION})

vendor/libwebrtc/out:
	(cd vendor/libwebrtc; gclient runhooks -j1)
	# (cd vendor/libwebrtc; ninja -C trunk/out/Release)

vendor/depottools:
	rm -rf vendor/depottools
	git clone ${DEPOT_TOOLS_REPO} vendor/depottools

env:
	(cd vendor/libwebrtc; which gclient)
