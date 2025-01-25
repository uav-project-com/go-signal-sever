package dto

type Sdp struct {
	Sdp string `json:"sdp"`
}

type CallInfo struct {
	MeetingID string `json:"meetingId"`
	UserId    string `json:"userId"`
	PeerId    string `json:"peerId"`
	IsSender  *bool  `json:"isSender"`
	Session   *Sdp   `json:"sdp"`
}
