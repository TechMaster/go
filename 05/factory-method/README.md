## Factory Pattern

## Factory Pattern là gì

Factory Pattern là một pattern thuộc nhóm creational patterns và nó là một trong những pattern được sử dụng phổ biến nhất. 

Factory Pattern đưa toàn bộ logic của việc tạo mới object vào trong factory, che giấu logic của việc khởi tạo đối tượng

Client chỉ tương tác với factory struct và cho biết loại instance cần được khởi tạo. Factory struct sẽ tương tác với các struct tương ứng và sau đó trả về một instance cụ thể với struct đó

## Ví dụ

Chức năng tạo bài viết với 2 thể loại bài viết là : Bài viết thông thường, bài viết tin hot