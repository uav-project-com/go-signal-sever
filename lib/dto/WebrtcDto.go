package dto

type Sdp struct {
	Sdp string
}

type CallInfo struct {
	MeetingID string `json:"meetingId"`
	UserId    string `json:"userId"`
	PeerId    string `json:"peerId"`
	IsSender  *bool  `json:"isSender"`
}
