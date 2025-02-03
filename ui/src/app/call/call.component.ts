import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Sdp} from './Sdp';
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-call',
  templateUrl: './call.component.html',
  styleUrls: ['./call.component.css']
})
export class CallComponent implements OnInit {

  pcSender: any
  pcReceiver: any
  meetingId: string
  peerId: string
  userId: string
  URL = "/api/v1/webrtc/start-call"

  constructor(private http: HttpClient, private route: ActivatedRoute) {
  }


  ngOnInit() {
    // use http://localhost:4200/call;meetingId=07927fc8-af0a-11ea-b338-064f26a5f90a;userId=alice;peerId=bob
    // and http://localhost:4200/call;meetingId=07927fc8-af0a-11ea-b338-064f26a5f90a;userId=bob;peerId=alice
    // start the call
    this.meetingId = this.route.snapshot.paramMap.get("meetingId");
    this.peerId = this.route.snapshot.paramMap.get("peerId");
    this.userId = this.route.snapshot.paramMap.get("userId")

    this.pcSender = new RTCPeerConnection({
      iceServers: [
        {urls: "stun:stun.l.google.com:19302"},
        {urls: "stun:stun.l.google.com:5349"},
        {urls: "stun:stun1.l.google.com:3478"},
        {urls: "stun:stun1.l.google.com:5349"},
        {urls: "stun:stun2.l.google.com:19302"},
        {urls: "stun:stun2.l.google.com:5349"},
        {urls: "stun:stun3.l.google.com:3478"},
        {urls: "stun:stun3.l.google.com:5349"},
        {urls: "stun:stun4.l.google.com:19302"},
        {urls: "stun:stun4.l.google.com:5349"}
      ]
    })
    this.pcReceiver = new RTCPeerConnection({
      iceServers: [
        {urls: "stun:stun.l.google.com:19302"},
        {urls: "stun:stun.l.google.com:5349"},
        {urls: "stun:stun1.l.google.com:3478"},
        {urls: "stun:stun1.l.google.com:5349"},
        {urls: "stun:stun2.l.google.com:19302"},
        {urls: "stun:stun2.l.google.com:5349"},
        {urls: "stun:stun3.l.google.com:3478"},
        {urls: "stun:stun3.l.google.com:5349"},
        {urls: "stun:stun4.l.google.com:19302"},
        {urls: "stun:stun4.l.google.com:5349"}
      ]
    })

    this.pcSender.onicecandidate = (event: { candidate: null; }) => {
      if (event.candidate === null) {
        this.http.post<Sdp>('http://192.168.0.200:8080' + this.URL,
          {
            "session": {"sdp": btoa(JSON.stringify(this.pcSender.localDescription))},
            "meetingId": this.meetingId,
            "userId": this.userId,
            "peerId": this.peerId,
            "isSender": true
          }
        ).subscribe(response => {
          console.log("SDP: " + JSON.stringify(response["sdp"]))
          this.pcSender.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(response["sdp"]))))
        });
      }
    }
    this.pcReceiver.onicecandidate = (event: { candidate: null; }) => {
      if (event.candidate === null) {
        this.http.post<Sdp>('http://192.168.0.200:8080' + this.URL,
          {
            "session": {"sdp": btoa(JSON.stringify(this.pcReceiver.localDescription))},
            "meetingId": this.meetingId,
            "userId": this.userId,
            "peerId": this.peerId,
            "isSender": false
          }
        ).subscribe(response => {
          console.log("SDP: " + JSON.stringify(response["sdp"]))
          this.pcReceiver.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(response["sdp"]))))
        });
      }
    }
  }

  startCall() {
    // sender part of the call
    if (this.userId === "alice") {
      navigator.mediaDevices.getUserMedia({video: true, audio: false}).then((stream) => {
        const senderVideo: any = document.getElementById('senderVideo');
        senderVideo.srcObject = stream;
        const tracks = stream.getTracks();
        for (let i = 0; i < tracks.length; i++) {
          this.pcSender.addTrack(stream.getTracks()[i]);
        }
        this.pcSender.createOffer().then(d => this.pcSender.setLocalDescription(d))
      })
    }
    // you can use event listner so that you inform he is connected!
    this.pcSender.addEventListener('connectionstatechange', event => {
      if (this.pcSender.connectionState === 'connected') {
        console.log("horray!")
      }
    });

    // receiver part of the call
    this.pcReceiver.addTransceiver('video', {'direction': 'recvonly'})

    this.pcReceiver.createOffer()
      .then((d: any) => this.pcReceiver.setLocalDescription(d))

    this.pcReceiver.ontrack = function (event: { streams: any[]; }) {
      const receiverVideo: any = document.getElementById('receiverVideo');
      receiverVideo.srcObject = event.streams[0]
      receiverVideo.autoplay = true
      receiverVideo.controls = true
    }

  }

}
