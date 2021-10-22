package main

import (
	"fmt"

	"github.com/pion/webrtc/v3"
)

var peerConnectionConfig = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs: []string{"stun:stun.l.google.com:19302"},
		},
	},
}

func core_handeler(sdpchan chan string) {
	offer := webrtc.SessionDescription{}
	Decode(<-sdpchan, &offer)

	peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfig)

	if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		panic(err)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	<-gatherComplete

	sdpchan <- Encode(*peerConnection.LocalDescription())
}
