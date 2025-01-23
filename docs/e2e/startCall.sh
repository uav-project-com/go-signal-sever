curl --request POST \
  --header "Content-Type: application/json" \
  --data '{
    "meetingId": "07927fc8-af0a-11ea-b338-064f26a5f90a",
    "userId": "alice",
    "peerId": "bob",
    "isSender": true
  }' \
  http://localhost:8080/api/v1/webrtc/start-call -vv