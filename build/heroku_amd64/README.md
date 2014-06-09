# Instructions

```bash
mkdir -p $GOPATH/src/github.com/fd
git clone git://github.com/fd/webrtc.git $GOPATH/src/github.com/fd/webrtc
cd $GOPATH/src/github.com/fd/webrtc

# when you want to rebuild the the static libs
build/linux_amd64/bin/deps
build/linux_amd64/bin/build

apt-get install build-essential
apt-get install libexpat-dev libssl-dev libnss3-dev libx11-dev libxext-dev libxss-dev
go get .
```
