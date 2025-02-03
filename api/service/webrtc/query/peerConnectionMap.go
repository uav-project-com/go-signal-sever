package query

import (
	env "boilerplate/lib/environment"
	"github.com/pion/webrtc/v2"
	"sync"
)

type PeerConnectionService struct {
	Connections   sync.Map
	TrackChannels map[string]chan *webrtc.Track
	Config        webrtc.Configuration
	Api           *webrtc.API
}

func NewPeerConnectionService() *PeerConnectionService {
	m := webrtc.MediaEngine{}

	// Setup the codecs you want to use.
	// Only support VP8(video compression), this makes our proxying code simpler
	// 90000: 90000Hz - tần số lấy mẫu của video cho realtime 30fps
	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))
	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))
	peerConnectionConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{{URLs: env.GetStrings("STUN")}},
	}

	// chatgpt: https://chatgpt.com/share/6794fe47-0698-8002-8834-fb4951073255
	return &PeerConnectionService{
		Connections:   sync.Map{},                          // TODO: what chatgpt var using for?
		TrackChannels: make(map[string]chan *webrtc.Track), // peerConnectionMap
		Config:        peerConnectionConfig,
		Api:           api,
	}
}

// AddConnection Example: Adding a connection
func (s *PeerConnectionService) AddConnection(id string, pc *webrtc.PeerConnection) {
	s.Connections.Store(id, pc)
}

// GetConnection Example: Retrieving a connection
func (s *PeerConnectionService) GetConnection(id string) (*webrtc.PeerConnection, bool) {
	val, ok := s.Connections.Load(id)
	if !ok {
		return nil, false
	}
	return val.(*webrtc.PeerConnection), true
}
