# Go-DDD-Layout
<div>

  ![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
  ![Release](https://img.shields.io/badge/release-1.0.0-green.svg)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
</div>

<p style="font-size: 15px">
  这是一个简易的  Go-DDD-Layout 项目，该项目实现了 User 的 CURD 基础操作，区分管理员角色和普通用户权限，通过它可以了解一个工程化 Go 项目应该如何构建，以及领域驱动设计（Domain-Driven-Design）架构带给我们哪些思考。
</p>

简体中文 | [English](./README.md)

## 环境部署

### docker 手动方式

```bash
➜  ~ docker rm -f $(docker ps -aq)
➜  ~
➜  ~ docker pull mysql:latest
➜  ~
➜  ~ docker images
REPOSITORY                                      TAG       IMAGE ID       CREATED       SIZE
mysql                                           latest    5d2fb452c483   9 days ago    622MB
➜  ~
➜  ~ docker run --name testdb -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql:latest
b2b36cb05f1e2d2c6c877557e5502802adf57ea43c57b35b0b885a1db6bd7c19
➜  ~
➜  ~ docker ps -a
CONTAINER ID   IMAGE          COMMAND                   CREATED          STATUS          PORTS                                                  NAMES
b2b36cb05f1e   mysql:latest   "docker-entrypoint.s…"   17 seconds ago   Up 17 seconds   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   testdb

➜  ~ docker cp ~/go-ddd/infrastructure/database/data.sql testdb:/data.sql
Successfully copied 12.8kB to testdb:/data.sql
➜  ~
➜  ~ docker exec -it testdb bash
bash-4.4# ls -lh data.sql
-rw-r--r-- 1 502 games 11K Nov  3 10:01 data.sql
bash-4.4#
bash-4.4# mysql -u root -p -h 127.0.0.1 -P 3306 -p123456
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.2.0 MySQL Community Server - GPL

Copyright (c) 2000, 2023, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> CREATE DATABASE mydb;
Query OK, 1 row affected (0.02 sec)
mysql>
mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mydb               |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.01 sec)
mysql>
mysql> USE mydb;
Database changed
mysql>
mysql> source data.sql;
...
Query OK, 1 row affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)
mysql>
mysql> SHOW TABLES;
+-------------------+
| Tables_in_mydb    |
+-------------------+
| t_rbac_role       |
| t_rbac_user       |
| t_rbac_user_roles |
+-------------------+
3 rows in set (0.00 sec)
mysql>
mysql> select count(1) from t_rbac_user;
+----------+
| count(1) |
+----------+
|       11 |
+----------+
1 row in set (0.00 sec)
mysql>
mysql> \q
Bye
bash-4.4# exit
exit
```

### 使用 docker-compose

```bash
➜  ./build.sh
docker-compose up -d --build
[+] Building 17.4s (15/15) FINISHED                                                                                                                                                         docker:orbstack
 => [app internal] load .dockerignore                                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                                        0.0s
 => [app internal] load build definition from Dockerfile                                                                                                                                               0.0s
 => => transferring dockerfile: 998B                                                                                                                                                                   0.0s
 => [app internal] load metadata for docker.io/library/golang:1.21                                                                                                                                     2.2s
 => [app builder 1/4] FROM docker.io/library/golang:1.21@sha256:b113af1e8b06f06a18ad41a6b331646dff587d7a4cf740f4852d16c49ed8ad73                                                                       0.0s
 => [app internal] load build context                                                                                                                                                                  0.0s
 => => transferring context: 10.34kB                                                                                                                                                                   0.0s
 => CACHED [app builder 2/4] WORKDIR /build                                                                                                                                                            0.0s
 => [app builder 3/4] COPY . /build                                                                                                                                                                    0.0s
 => [app builder 4/4] RUN go build -o app .                                                                                                                                                           15.1s
 => CACHED [app stage-1 2/7] WORKDIR /app                                                                                                                                                              0.0s
 => CACHED [app stage-1 3/7] RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime                                                                                                               0.0s
 => CACHED [app stage-1 4/7] RUN echo "Asia/Shanghai" > /etc/timezone                                                                                                                                  0.0s
 => CACHED [app stage-1 5/7] COPY --from=builder /build/app .                                                                                                                                          0.0s
 => CACHED [app stage-1 6/7] COPY .env .                                                                                                                                                               0.0s
 => CACHED [app stage-1 7/7] RUN apt-get update && apt-get install -y nmap                                                                                                                             0.0s
 => [app] exporting to image                                                                                                                                                                           0.0s
 => => exporting layers                                                                                                                                                                                0.0s
 => => writing image sha256:27e5e792fad523c3c17a99520a23bb1adb51fa92e9810ce61368f4ef21c61f29                                                                                                           0.0s
 => => naming to docker.io/library/go-ddd                                                                                                                                                              0.0s
[+] Running 3/3
 ✔ Container mysql-db   Started                                                                                                                                                                           0.0s
 ✔ Container go-app     Started                                                                                                                                                                           0.0s
 ✔ Container nginx-web  Started                                                                                                                                                                           0.0s

# 注意添加host解析
➜  sudo vi /etc/hosts
# echo "127.0.0.1 go-ddd-layout.com" >> /etc/hosts

➜  docker ps -a
CONTAINER ID   IMAGE          COMMAND                   CREATED          STATUS          PORTS                                                  NAMES
a87de7ec1c88   nginx:latest   "/docker-entrypoint.…"   31 seconds ago   Up 30 seconds   0.0.0.0:80->80/tcp, :::80->80/tcp                      nginx-web
ccb49c2afaf9   go-ddd         "/bin/sh -c 'while !…"   31 seconds ago   Up 30 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp              go-app
7006bf52356a   mysql:latest   "docker-entrypoint.s…"   31 seconds ago   Up 30 seconds   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql-db
```

## 接口文档

swagger:

http://go-ddd-layout.com/swagger/index.html

## 目录树

使用 DDD 提升代码质量

```bash
.
├── .env                                    # 环境变量配置文件
├── .git                                    # Git版本控制目录
├── .gitignore                              # Git忽略文件列表
├── Dockerfile                              # 构建Docker镜像 
├── Makefile                                # 项目构建和管理的Makefile
├── README.md                               # 项目的说明文档
├── application                             # 应用层目录
│   ├── event                               # 事件相关代码
│   └── service                             # 应用服务目录
│       └── user                            # 用户应用服务目录
│           └── application_service.go      # 用户应用服务实现代码
├── build.sh                                # 构建脚本
├── docker-compose.yml                      # Docker Compose配置文件
├── docs                                    # 文档目录
│   ├── docs.go                             # 文档生成代码
│   ├── swagger.json                        # Swagger JSON文件
│   └── swagger.yaml                        # Swagger YAML文件
├── domain                                  # 领域层目录
│   └── user                                # 用户领域目录
│       ├── entity                          # 领域实体目录
│       │   ├── role.go                     # 角色实体定义
│       │   ├── user.go                     # 用户实体定义
│       │   └── valueobj                    # 值对象目录
│       │       └── value_object.go         # 值对象定义
│       ├── repository                      # 领域仓储目录
│       │   ├── role.go                     # 角色仓储接口定义
│       │   └── user.go                     # 用户仓储接口定义
│       └── service                         # 领域服务目录
│           └── domain_service.go           # 领域服务实现代码
├── go.mod                                  # Go模块文件
├── go.sum                                  # Go依赖版本控制文件
├── infrastructure                          # 基础设施层目录
│   ├── common                              # 公共代码目录
│   │   ├── auth                            # 身份认证相关代码
│   │   │   └── header_token.go             # 头部令牌身份认证实现
│   │   ├── context                         # 上下文相关代码
│   │   │   ├── getter.go                   # 上下文获取函数
│   │   │   └── store.go                    # 上下文存储函数
│   │   ├── cookie                          # Cookie操作代码
│   │   │   └── cookie.go                   # Cookie操作实现
│   │   ├── errors                          # 错误处理代码
│   │   │   └── errors.go                   # 自定义错误类型和处理函数
│   │   ├── jwt                             # JWT相关代码
│   │   │   ├── claims.go                   # JWT声明结构定义
│   │   │   └── token.go                    # JWT令牌生成和解析实现
│   │   ├── orm                             # ORM操作相关代码
│   │   │   └── pagination.go               # 分页查询功能实现
│   │   ├── random                          # 随机数生成相关代码
│   │   │   └── generate_password.go        # 生成随机密码实现
│   │   └── response                        # 响应处理代码
│   │       └── response.go                 # 响应处理函数实现
│   ├── database                            # 数据库相关代码
│   │   ├── data.sql                        # 数据库初始化脚本
│   │   └── db.go                           # 数据库连接和操作实现
│   ├── persistence                         # 持久化层目录
│   │   ├── index.go                        # 索引文件
│   │   └── user                            # 用户持久化目录
│   │       ├── converter                   # 转换器目录（说明：领域对象Entity与PO之间的转化）
│   │       │   └── user.go                 # 用户转换器实现
│   │       ├── po                          # 持久化对象目录
│   │       │   ├── role.go                 # 角色持久化对象定义
│   │       │   └── user.go                 # 用户持久化对象定义
│   │       └── user_repo.go                # 用户持久化实现
│   └── readme.md                           # 基础设施层说明文档
├── interfaces                              # 接口层目录（说明：尽管Go官方不建议目录起名为复数(s)，但interface为关键词，因此这里可以为复数形式）
│   ├── adapter                             # 适配器目录，适配Web框架及协议接入，比如：gin/fiber等
│   │   ├── initialize                      # 初始化目录
│   │   │   ├── app.go                      # 应用初始化实现
│   │   │   ├── engine.go                   # 路由引擎实现
│   │   │   └── service.go                  # 应用服务实现
│   │   ├── middleware                      # 中间件目录
│   │   │   ├── auth.go                     # 身份认证中间件实现
│   │   │   ├── cors.go                     # 跨域中间件实现
│   │   │   ├── error.go                    # 错误处理中间件实现
│   │   │   └── logging.go                  # 日志记录中间件实现
│   │   └── router                          # 路由配置目录
│   │       └── router.go                   # 路由配置实现
│   ├── assembler                           # 转换器目录（说明：领域对象Entity与DTO之间的转化）
│   │   └── user                            # 用户转换器目录
│   │       └── user.go                     # 用户转换器实现
│   ├── controller                          # 控制器目录
│   │   ├── enter.go                        # 入口控制器
│   │   ├── public                          # 公共控制器目录
│   │   │   └── hello_world.go              # HelloWorld公共控制器实现
│   │   ├── sys                             # 系统控制器目录
│   │   │   ├── enter.go                    # 系统入口控制器
│   │   │   ├── login.go                    # 登录控制器实现
│   │   │   ├── logout.go                   # 注销控制器实现
│   │   │   └── menu.go                     # 菜单控制器实现
│   │   └── user                            # 用户控制器目录
│   │       ├── create_user.go              # 创建用户控制器实现
│   │       ├── delete_user.go              # 删除用户控制器实现
│   │       ├── enter.go                    # 用户入口控制器
│   │       ├── update_password.go          # 更新密码控制器实现
│   │       ├── update_user.go              # 更新用户控制器实现
│   │       ├── user_detail.go              # 用户详情控制器实现
│   │       ├── user_info.go                # 用户信息控制器实现
│   │       └── user_list.go                # 用户列表控制器实现
│   └── dto                                 # 数据传输对象目录
│       ├── sys                             # 系统DTO目录
│       │   ├── login.go                    # 登录DTO定义
│       │   └── menu.go                     # 菜单DTO定义
│       └── user                            # 用户DTO目录
│           └── user.go                     # 用户DTO定义
├── main.go                                 # 主入口文件
└── nginx.conf                              # nginx配置文件
```

## 文章

在[这里](https://juejin.cn/post/7298160530292703244)写了一篇关于这个项目的文章

## 项目扩展

如果你的项目非常注重安全性，需要处理敏感信息如 JWT 的私钥和数据库口令时，固定并放置在环境变量中仍然存在一些潜在的安全风险。

为了进一步增强安全性，我们可以采用更加安全的方式来处理这些敏感信息：

- 使用密钥管理系统（KMS）：例如AWS KMS、Google Cloud KMS或HashiCorp Vault，通过API从这些系统中获取私钥。

- 使用公钥/私钥对：考虑使用非对称加密算法，如RSA或ECDSA，而不是对称加密算法。将私钥存储在安全位置，而公钥可以在应用程序中使用。

当然，无论选择哪种方式，都需要遵循一些基本原则，包括限制对密码的访问权限、定期轮换密码、监控和审计密码的使用等。

对于密钥的轮换和管理，Redis可以作为一个可行的选择。以下是在使用Redis时实现密钥轮换的一些思路：

- 存储密钥：将当前有效的密钥存储在Redis中，可以使用Redis的字符串数据类型进行存储。

- 定期轮换：根据预定的时间间隔生成新的密钥，并将新密钥存储到Redis中。

- 平滑过渡期：在旧密钥失效之前，验证和解析仍然接受新旧密钥，以确保能够处理由旧密钥签发的令牌。

- 更新令牌签发：生成新的令牌时，使用存储在Redis中的新密钥进行签名。

- 撤销旧密钥：过渡期结束后，从Redis中删除旧密钥。

- 定期执行：设置定时任务或使用定时器，在规定的时间间隔内执行密钥轮换操作。

使用Redis作为密钥存储的好处包括高性能、可靠性和灵活性。但需要注意确保Redis实例的安全性和可靠性非常重要。请采取适当的措施来保护Redis实例，例如使用访问控制列表（ACL）或访问密码，以防止未经授权的访问。此外，还应备份和监控Redis实例，以确保数据的完整性和可用性。
