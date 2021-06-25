# app server case

- 为了学习 Kubernetes 工作资源使用姿势，我们自己定义一个服务，用来跟随官方文档的脚步，自己尝试
- 服务本身比较简单，只是一个 HTTP 服务，随着工作资源的复杂，一步一步丰富服务内容

## 使用姿势

- 服务启动：go run main.go
- 通过浏览器访问：http://localhost:8080

## 镜像打包

- make image v=1.0.0

## 启动容器

- docker run --rm -it -p 8080:8080 app:1.0.0
