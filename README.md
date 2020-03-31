## Result:

### Basic use:

Truy cập vào website `localhost:6000` để sử dụng những tính năng cơ bản nhất của web

Tài khoản:

- email: `root@gmail.com`
- password: `1234567`

### Report solution:

1. Với top 5 phim ăn khách nhất trong 3 ngày gần đây:
    - Tính toán độ hot của phim thông qua số lượt xem và rating của phim
    - Sau đó lưu độ hot theo thứ tự vào trong field `algorithm_point`
    - Để thuận lợi và tăng tốc độ xử lý cho việc trả về thông tin liên quan đến phim
    - Điểm số này có thể  không quan trọng bằng điểm số được quyết định bằng tay bởi admin thông qua field `manual_point`
    - Ta sẽ có 1 job chạy 24h mỗi lượt cho việc update lại điểm số  `algorithm_point`
    - Cache data lại để tăng tính trải nghiệm cho người dùng sau khi đã lưu thông tin lại

2. Số lỗi xảy ra hằng ngày theo API response:

***Database Design:***

ID | Err_Code | Type | Created_At | API_Router 

- ID: lưu id của lỗi
- Err_Code: Lưu mã lỗi
- Type: Lưu dạng của lỗi (client, server, other)
- Created_At: Lỗi xuất hiện vào lúc nào
- API_Router: Xuất phát từ API nào

### Docker-compose for local dev:

**Requirement:**

- make pkg
- linux or macOS env

**Flow:**

Để setup docker-compose cho local dev cần thực hiện những cmd sau:

- make local-db (Cài đặt db ở local)
- make setup-package (Cài đặt package vendor cho golang dev)

### CI/CD with CircleCI:

Set up CI/CD với CircleCI và digitalocean

Qua 3 bước:

`test`:

circleCI runner sẽ sử dụng image có chứa go.
Install các package cần thiết và tiến hành chạy các test case có trong project.
Bước test tiến hành ở các branch

`build`:

***Điều kiện cần thiết:***

Thực hiện được bước `test` 

circleCI runner sẽ build project thành docker image.
Sau đó đẩy lên docker hub
Chỉ tiến hành ở branch `master`

`deploy`:

***Điều kiện cần thiết:***

Thực hiện được bước `build`

circleCI runner sẽ truy cập vào server có sẳn thông qua ssh service.
Sau đó tiến hành pull docker image từ docker hub về thống qua các file docker-compose có sẵn.
Cuối cùng tiền hành mở các docker container lên
Chỉ tiến hành ở branch `master`

### Monitor:

Sử dụng prometheus và grafana cho việc monitor server

Truy cập `localhost:3000/` (grafana)

Tài khoản:

- username: admin
- password: pass
