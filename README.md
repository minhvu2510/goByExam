### Interface:
Interface xác định hành vi của một đối tượng. Interface chỉ xác định những gì đối tượng phải làm còn cách làm thế nào thì tuỳ thuộc vào đối tượng đó.

Trong Go, interface là một tập khai báo các phương thức (methods). Khi một khiểu dữ liệu (struct) định nghĩa đầy đủ các phương thúc(method) trong một interface thì nó dc gọi là implement interface đó (thực hiện interface đó). Interface xác định rằng: Một kiểu dữ liệu nên có những phương thức nào và cách nó implement những phương thức đó.

Ví dụ : WashingMachine là một interface chứa các khai báo phương thức như Cleaning() và Drying(). Bất kì kiểu dữ liệu nào định nghĩa ra phương thức Cleaning() và Drying() đều được gọi là implement interface WashingMachine.

 Công dụng (vi du file interface.go): Chương trình tính lương các loiạ nhân viên của một công ty. Mỗi loại có một cách tính lương riêng. Hàm totalExpnse thẻ hiện tiện ích của interface. Nó sẽ nhận đầu vào là 1 slice các interface. Trong interface chứa method tính lương của từng loại nhân viên. => truyền tất cả nhân viên vào sẽ tính dc lương phải trả. ưu điểm của hàm này là có thẻ mở rộng đến bất kì loại nhân viên moiws nào mà ko cân phải thay đổi code