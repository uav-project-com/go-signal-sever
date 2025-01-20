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