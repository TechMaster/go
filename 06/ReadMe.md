# Go Routinee

Kinh nghiệm:
- CRUD --> SQL DATABASE có transaction
- File upload đồng thời CRUD, request service khác
  - insert table ok --> commit
  - save file error

  - commit và rollback cho các tác vụ đa dạng sẽ phải do lập trình viên viết.
  - Có thể tận dụng go routine và channel để phối hợp khi các tác vụ này kéo dài, không rõ lúc nào mới hoàn thành
  - go routine bản chất là async. Nhưng chúng ta dùng package sync để đồng bộ 

  + sync.Mutex, sync.WaitGroup, sync.Once dùng để synchronize các asynchorous running routine

 - Kích thước buffer: đủ to để phía receiver kịp xử lý các message đến mà không bị nghẽn, đừng quá lớn vì sẽ tốn bộ nhớ.