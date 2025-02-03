## rtcpPLIInterval
Biến `rtcpPLIInterval` trong đoạn code trên đại diện cho một khoảng thời gian (kiểu `time.Duration`) để gửi các gói RTCP **Picture Loss Indication (PLI)** đến một nguồn phát video từ xa (publisher). Mục đích chính là đảm bảo nguồn phát gửi lại keyframe thường xuyên để khôi phục video nếu có sự cố mất gói hoặc lỗi khung hình.

### **Giải thích chi tiết**:
- **RTCP PLI**:
    - Đây là một loại gói tin thuộc giao thức RTCP (Real-Time Control Protocol), được sử dụng trong truyền thông video thời gian thực.
    - Gói PLI được gửi từ người nhận (receiver) đến nguồn phát (publisher) để yêu cầu gửi lại một keyframe (khung hình chính). Keyframe rất quan trọng vì nó là khung hình hoàn chỉnh, không phụ thuộc vào dữ liệu từ các khung hình khác.

- **Mục đích của `rtcpPLIInterval`**:
    - Thiết lập một khoảng thời gian cố định (interval) để gửi gói PLI định kỳ, nhằm giảm thiểu gián đoạn khi khung hình bị mất hoặc không đồng bộ.
    - Trong đoạn code, `time.NewTicker(rtcpPLIInterval)` được sử dụng để tạo một ticker, phát tín hiệu sau mỗi khoảng thời gian `rtcpPLIInterval`.
    - Gói PLI được gửi qua phương thức `peerConnection.WriteRTCP`.

- **Cách hoạt động trong đoạn code**:
    - Một goroutine (hàm chạy song song) được khởi tạo để định kỳ gửi gói PLI dựa trên `rtcpPLIInterval`.
    - Nếu xảy ra lỗi trong quá trình gửi gói RTCP, lỗi sẽ được in ra console thông qua `fmt.Println`.

### **Ví dụ giá trị cho `rtcpPLIInterval`**:
Trong thực tế, giá trị này có thể được cấu hình, ví dụ:
- **500ms**: Gửi yêu cầu keyframe mỗi 500ms.
- **1s**: Gửi mỗi giây, phù hợp với độ trễ thấp trong video streaming.

### **Tối ưu hóa**:
Trong đoạn code có gợi ý rằng việc gửi PLI định kỳ có thể "lãng phí tài nguyên". Một cách tối ưu hơn là lắng nghe các sự kiện RTCP từ người xem (viewers), chỉ gửi PLI khi cần thiết (ví dụ, khi nhận được yêu cầu NACK hoặc PLI từ phía người xem). Điều này giúp tiết kiệm băng thông và tài nguyên xử lý.

### **Purpose of Clockrate**

For 30 frames per second (FPS) real-time video, the typical **clockrate** used in WebRTC and RTP standards is **90,000 Hz**.

In WebRTC, the `clockrate` parameter of `NewRTPCodec` specifies the sampling rate, in Hertz (Hz), of the media being encoded or decoded. This value is crucial for synchronizing media streams and interpreting RTP timestamps, as it defines how RTP timestamps are calculated and incremented over time.


1. **Timestamp Calculation:**
    - RTP timestamps are derived based on the clockrate. For example, with a clockrate of 48,000 Hz (common for audio), the RTP timestamp increases by 48,000 units per second of audio.

2. **Media Synchronization:**
    - The clockrate ensures that media from multiple sources (e.g., audio and video) remains synchronized during playback.

3. **Codec Specification:**
    - Different codecs operate at specific clock rates. For example:
        - Audio codecs like Opus typically use 48,000 Hz.
        - Video codecs like VP8 or H.264 generally use 90,000 Hz.

4. **Interoperability:**
    - When different endpoints communicate, the clockrate helps ensure proper timing alignment and seamless streaming between devices.

### **Common Clockrate Values**
- **Audio:**
    - 8,000 Hz: Common for narrowband codecs (e.g., G.711).
    - 16,000 Hz: Used for wideband audio (e.g., Speex).
    - 48,000 Hz: Preferred for high-quality codecs like Opus.

- **Video:**
    - 90,000 Hz: A standard value for most video codecs, aligning with real-time video playback rates.

### **Example in WebRTC**
Here's an example of creating a new RTP codec with a clockrate in Go:
```go
codec := webrtc.NewRTPCodec(
    webrtc.RTPCodecTypeAudio,
    "opus",    // Codec name
    48000,     // Clockrate in Hz
    2,         // Channels
    "",        // Feedback mechanisms (e.g., "transport-cc")
    0,         // Payload type
)
```

In this example:
- The clockrate of 48,000 Hz aligns with the Opus codec's default sampling rate.
- RTP timestamps increment by 48,000 per second of audio.

By setting the appropriate clockrate, you ensure the codec operates as expected and that media streams remain synchronized.

### NewRTPCodec - Clockrate: Why 90,000 Hz?
- **Industry Standard**: The RTP specification (RFC 3550) defines 90,000 Hz as the standard clockrate for video to provide high precision for timestamping and synchronization.
- **Precision**: This value allows for fractional frame durations to be expressed accurately, especially for common video frame rates such as 30 FPS and 60 FPS.
    - At 30 FPS:
        - Each frame duration is approximately \( \frac{1}{30} \) seconds, or 33.33 milliseconds.
        - Using a 90,000 Hz clockrate, the timestamp increment per frame is:
          \[
          \text{Timestamp increment} = \frac{\text{Clockrate}}{\text{Frame rate}} = \frac{90,000}{30} = 3,000
          \]
          So, the RTP timestamp for each video frame increases by 3,000 units.

### Clockrate Examples for Video Codecs
- **H.264**: Commonly uses 90,000 Hz.
- **VP8/VP9**: Typically uses 90,000 Hz.
- **AV1**: Also aligns with the 90,000 Hz standard in most RTP implementations.

### How It Affects Real-Time Video
The clockrate ensures:
1. **Accurate Timestamps**: RTP timestamps increment consistently with frame intervals.
2. **Synchronization**: Audio and video streams can be synchronized since audio clockrates (e.g., 48,000 Hz for Opus) can be aligned with video using the RTP standard's timing mechanisms.

By using the standard 90,000 Hz clockrate for 30 FPS video, you achieve precise frame timing and compatibility with most RTP-based systems.