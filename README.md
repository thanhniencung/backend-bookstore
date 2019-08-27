Các bước cài đặt server:

Bước 1. Cài đặt Database 
- 	Cài đặt pgAdmin, link: https://www.pgadmin.org/
-	Sau khi cài xong thì tạo user để truy cập vào database(có thể dụng default or tự tạo)
-	Tạo một database bên trong pgAdmin (tham khảo trong clip)
-	Chạy script khởi tạo các table trong database 

Bước 2. Cài đặt Golang, link: https://golang.org/dl/

Bước 3. Cài đặt git (tool này cần để Golang download các thư viện về)

Bước 4. Chạy source server
-	Chỉnh thông tin kết nối database trong cmd/main.go (tham khảo video)
-	Chạy server:
-		Cách 1: đi tới thư mục cmd/ rồi gõ lệnh: go run main.go 
		Cách 2: Nếu bạn nào đã cài lệnh Make trong máy thì chỉ cần đứng ở thư mục 
		gốc rồi gõ lệnh: make run 