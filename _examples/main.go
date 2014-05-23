package main

import (
	"github.com/fd/webrtc"
)

func main() {
	webrtc.InitializeSSL()

	factory := webrtc.New()

	stream := factory.CreateMediaStream("s")

	_ = stream
}
