# Instructions

```bash
mkdir -p $GOPATH/src/github.com/fd
git clone git://github.com/fd/webrtc.git $GOPATH/src/github.com/fd/webrtc
cd $GOPATH/src/github.com/fd/webrtc

# when you wan to rebuild the the static libs
apt-get install git -yy
build/linux_amd64/bin/deps
build/linux_amd64/bin/build

go get .
```
