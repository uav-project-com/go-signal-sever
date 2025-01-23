package webrtc

import (
	"boilerplate/api/service/webrtc/command"
)

type ConnectHandler struct {
	startCall command.WebrtcService
}

func NewWebrtcHandler(service command.WebrtcService) *ConnectHandler {
	return &ConnectHandler{startCall: service}
}
