# 项目
基于 gin 框架搭建的 go demo 项目


# 安装依赖
```
go get
```
# 数据库
需要启动一个mysql数据库，数据库配置 config/application.yaml
```
server:
  port: 8080
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: ginessential
  username: root
  password: admin123
  charset: utf8
  loc: Local
```

# 启动

```
go run main.go
```