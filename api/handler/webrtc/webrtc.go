package webrtc

import (
	"boilerplate/api/service/webrtc/query"
)

type ConnectHandler struct {
	startCall      query.WebrtcService
	peerConnection *query.PeerConnectionService
}

func NewWebrtcHandler(callSvc query.WebrtcService, peerSvc *query.PeerConnectionService) *ConnectHandler {
	return &ConnectHandler{
		startCall:      callSvc,
		peerConnection: peerSvc,
	}
}
