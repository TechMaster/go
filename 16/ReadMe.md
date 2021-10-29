# API Gateway

## Giới thiệu

API Gateway là một phần mềm máy chủ làm nhiệm vụ điều hướng chuyển các request đến (incoming request) tới đúng các dịch vụ phía sau. API Gateway cũng tương đương như reverse proxy.
![](api_gateway.jpg)

![What Is an API Gateway?](https://youtu.be/hYgP0cBORVg)

Chức năng chính của API Gateway hiện nay gồm có:
1. Điều hướng: route
2. Tự phát hiện dịch vụ: discover service
3. Cân bằng tải: load balancing
4. Giao diện quản trị: control dashboard
5. Viết, nạp thêm plugin
6. Viết, nạp thêm middleware. Middleware can thiệp incoming request, biến đổi nó trước điều hướng: Authentication, change header, rate limiter, circuit breaker, log...


## Những API Gateway phổ biến
1. HA-Proxy: viết hoàn toàn bằng C, C++. Rất chuyên làm Reverse Proxy. Chạy cực nhanh, tốn ít RAM, ổn định. Bản community rất ít chức năng, giao diện kém. Bản enterprize thu phí có đầy đủ giao diện quản trị, hỗ trợ Docker, Kubernetes
2. Nginx: viết hoàn toàn bằng C, C++, có hỗ trợ viết plugin bằng LUA.
Tốc độ tốt, dễ cấu hình. Bản community ít chức năng, không có giao diện điều khiển.
3. [Traefik](traefik.io): viết bằng Go. Tốc độ chỉ bằng 70% so với HA-Proxy, bộ nhớ tiêu thụ khi tải lớn gần gấp 2 so với HA-Proxy. Có thể viết thêm plugin cho Traefik 2.x sử dụng Go.
4. Kong: viết trên [OpenResty](https://openresty.org/en/) và ngôn ngữ lập trình LUA

Bên cạnh API Gateway còn có Service Mesh. Cách hoạt động của Service Mesh khác với API Gateway ở chỗ, Service Mesh route các lời gọi giữa service với nhau.

![](https://dz2cdn1.dzone.com/storage/temp/10249320-screen-shot-2018-09-17-at-115358-am.png)

## Giới thiệu Traefik

Tại sao tôi chọn Traefik?
Trả lời: câu trả lời cũng tương tự như tại sao bạn chọn xe Yamaha mà không chọn Honda hay Suzuki. Traefik không phải là API Gateway có hiệu suất hoạt động tốt nhất. Nó xếp sau HAProxy và Nginx. Nhưng phiên bản miễn phí của Traefik dễ cấu hình, dễ quản trị nhất, phù hợp với các dự án nhỏ.
![](https://doc.traefik.io/traefik/assets/img/traefik-architecture.png)

Giao diện quản trị thân thiện
![](dashboard.jpg)

Xem được tất cả cấu hình chuyển hướng

![](router.jpg)

Tích hợp chặt chẽ với Docker và Docker Swarm để cân băng tải và tăng khả năng sẵn sàng: kết nối thông qua Docker socks để nhận biết thay đổi Docker container replica.

![](loadbalancing.jpg)

Phiên bản Traefik mới nhất tại tháng 10/2021 hỗ trợ:
- HTTP/3
- Private và Provider Plugins
- Load Balacing Health Check
- K8s Ingress 1.22+
- Router Metrics
- Consul Connect

## Traefik căn bản
### 1. Entry Point định nghĩa cổng vào (port number)
Entry Point để hứng các request dạng TCP hoặc UDP. Phổ biến là 80 (http), 443 (https), 8080 (dashboard)
![](https://doc.traefik.io/traefik/assets/img/entrypoints.png)
```yaml
services:
  gateway:
    image: traefik:v2.5
    ports:      
      - "80:80"
      - "8080:8080"
```

hoặc
```yaml
services:
  gateway:
    image: traefik:v2.5
    networks:
      - techmaster
    command:
      - '--entryPoints.web.address=:80'
      - '--entryPoints.api.address=:9999'
      - "--entrypoints.websecure.address=:443"
```

### 2. Routers (chuyển hướng)
![](https://doc.traefik.io/traefik/assets/img/routers.png)

Cấu hình sử dụng file provider
```yaml
http:
  routers:
    to-whoami:
      rule: "Host(`localhost`) && PathPrefix(`/whoami`)"
      middlewares:
        - test-user
      service: whoami-dockerprovider@docker
```
Cấu hình sử dụng Docker label
```yaml
  mainsite:
    build: mainsite
    image: mainsite:latest
    labels:
      - "traefik.http.routers.mainsite.rule=Host(`iris.com`)"
      - "traefik.http.services.mainsite.loadbalancer.server.port=9001"
    depends_on:
      - "redis"
```

### 3. Service dịch vụ phía sau Traefik
![](https://doc.traefik.io/traefik/assets/img/services.png)

```yaml
services:
  main:
    image: registry.techmaster.com/mainsite:latest
    networks:
      - techmaster
      - techmaster_backend     
    deploy:
      labels:
        - "traefik.enable=true"
        - "traefik.http.routers.main.rule=Host(`techmaster.com`)"
        - "traefik.http.services.main.loadbalancer.server.port=8079" 
        - "traefik.docker.network=techmaster"

  admin:
    image: registry.techmaster.com/techmaster-admin-side:latest
    networks:
      - techmaster
      - techmaster_backend 
    deploy:
      labels:
        - "traefik.enable=true"
        - "traefik.http.routers.admin.rule=Host(`techmaster.com`) && PathPrefix(`/admin`)"
        - "traefik.http.services.admin.loadbalancer.server.port=8081" 
        - "traefik.docker.network=techmaster"        

  teacher:
    image: registry.techmaster.com/techmaster-teacher-side:latest
    networks:
      - techmaster
      - techmaster_backend       
    deploy:
      labels:
        - "traefik.enable=true"
        - "traefik.http.routers.teacher.rule=Host(`techmaster.com`) && PathPrefix(`/teacher`)"
        - "traefik.http.services.teacher.loadbalancer.server.port=8082" 
        - "traefik.docker.network=techmaster"

  user:
    image: registry.techmaster.com/techmaster-user-side:latest
    networks:
      - techmaster
      - techmaster_backend     
    deploy:
      labels:
        - "traefik.enable=true"
        - "traefik.http.routers.user.rule=Host(`techmaster.com`) && PathPrefix(`/user`)"
        - "traefik.http.services.user.loadbalancer.server.port=8086" 
        - "traefik.docker.network=techmaster"
```

### 4. Middleware xử lý trung gian
![](https://doc.traefik.io/traefik/assets/img/middleware/overview.png)

Ví dụ về Forward Auth middleware
![](forward_auth.jpg)

### 5. Tích hợp SSL Lets Encrypt vào Traefik
[Tích hợp Lets Encrypt vào Traefik V2.5+](https://techmaster.vn/posts/36749/tich-hop-lets-encrypt-vao-traefik-v2-5)

## Các bài lab thực hành Traefik

1. [Sử dụng File Provider](https://github.com/TechMaster/LearnTraefik/tree/main/FileProvider)
2. [Sử dụng Docker Provider](https://github.com/TechMaster/LearnTraefik/tree/main/DockerProvider)
3. [ForwardAuth middleware](https://github.com/TechMaster/LearnTraefik/tree/main/ForwardAuth)
4. [Gateway Authentication sử dụng Session - Cookie](https://github.com/TechMaster/LearnTraefik/tree/main/ForwardAuth)
5. [Gateway Roled Based Access Control Auth](https://github.com/TechMaster/LearnTraefik/tree/main/GatewayRBAC)
6. [Bảo mật Traefik Dashboard](https://github.com/TechMaster/LearnTraefik/tree/main/SecureDashBoard)
7. [Google OAuth Authentication](https://github.com/TechMaster/LearnTraefik/tree/main/GoogleForwardAuth)