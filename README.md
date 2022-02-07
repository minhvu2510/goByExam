### Interface:
Interface xác định hành vi của một đối tượng. Interface chỉ xác định những gì đối tượng phải làm còn cách làm thế nào thì tuỳ thuộc vào đối tượng đó.

Trong Go, interface là một tập khai báo các phương thức (methods). Khi một khiểu dữ liệu (struct) định nghĩa đầy đủ các phương thúc(method) trong một interface thì nó dc gọi là implement interface đó (thực hiện interface đó). Interface xác định rằng: Một kiểu dữ liệu nên có những phương thức nào và cách nó implement những phương thức đó.

Ví dụ : WashingMachine là một interface chứa các khai báo phương thức như Cleaning() và Drying(). Bất kì kiểu dữ liệu nào định nghĩa ra phương thức Cleaning() và Drying() đều được gọi là implement interface WashingMachine.

 Công dụng (vi du file interface.go): Chương trình tính lương các loiạ nhân viên của một công ty. Mỗi loại có một cách tính lương riêng. Hàm totalExpnse thẻ hiện tiện ích của interface. Nó sẽ nhận đầu vào là 1 slice các interface. Trong interface chứa method tính lương của từng loại nhân viên. => truyền tất cả nhân viên vào sẽ tính dc lương phải trả. ưu điểm của hàm này là có thẻ mở rộng đến bất kì loại nhân viên moiws nào mà ko cân phải thay đổi code

 ### Goroutines 
 Goroutines là các hàm hoăc phương thức chạy đồng thời với các hàm/phương thức khác.Goroutines có thể dc coi là những luồng gọn nhẹ. Chi phí tạo Goroutins thấp hơn so với 1 luồng.(1 ứng dụng có thể có hàng ngàn Goroutinens chạy đồng thời). ưu điểm so với luồng

`*` Chi phí thấp hơn luồng (chiếm vài kb trong stack)
`*` Dc ghép với kênh với một số ít hơn các luồng HDH. Có thể chỉ có một luồng trong một ct với hàng ngàn goroutines.
`*` Goroutines trao đổi thông qua các kênh. Các kênh được thiết kế để ngăn ngừa các khả năng xung đột xảy ra khi truy cập bộ nhớ chia sẻ sử dụng Goroutines. 

File goroutines.go. Nhập từ khóa go vào trước hàm hoặc phương thức, chúng ta sẽ có một Goroutine chạy đồng thời.

### Channel
 Các goroutins chạy độc lập để các goroutine giao tiếp với nhau thì có Channels. Tương tự như cách nước chảy từ đầu này sang đầu kia trong đường ống, dữ liệu gửi từ một đầu và nhận từ đầu kia băng channels.

 Mỗi channel có một loại liên kết với nó. Loại này là loại dữ liệu mà channel được phép vận chuyển. Không có loại khác được phép vận chuyển bằng cách sử dụng channel. (chan T là một channel loại T). File chanel.go

### Race condition
Hiện tượng nhiều tiến trình cùng truy cập và muốn thay đổi giá trị của biens, nhưng không theo quy tắc nào khiến kết quả không như mong muốn. Để xử lý vấn đề nyaf cần dùng đến atomic, mutext, channel

### Waitgroup 
Công cụ quản lý luồng chạy của goroutines: go rt chính chờ go rt con có tín hiệu x/lý xong mới chạy tiếp

### Atomic 
Package của go đảm bỏa cho tại một thời điểm chỉ có 1 tiến trình duy nhất đọc ghi một biến (chỉ hỗ trợ các kiển số int32, int64, unit32,...)

### Mutex 
Giống atomic nhưng khác ở chỗ atomic chỉ lock 1 biến của mutex lock cả 1 đoạn code

 ### Worker Pool 

 Mục đích là để ta có thể quản lý các worker, quản lý việc phân phối task và đặc biệt là kiểm soát được những tài nguyên dùng chung giữa các worker. Ví dụ như các worker chạy đồng thời và cùng truy xuất vào 1 file hoặc dùng chung một API.
worker pool : tập hợp các goroutines chò task để xử lý
Tại sao dùng worker pool : quản lý được goroutines dễ dàng, tiết kiệm thời gian cấp phát go rt, kiểm soát size không bị exhausted resource