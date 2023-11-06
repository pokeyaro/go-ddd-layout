# Go-DDD-Layout
<div>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
![Release](https://img.shields.io/badge/release-1.0.0-green.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
</div>

<p style="font-size: 15px">
  This is a simple Go-DDD-Layout project that implements basic CRUD operations for User, differentiating between the roles of administrator and regular user. Through this project, you can learn how to build a structured Go project and explore the considerations brought by the Domain-Driven Design (DDD) architecture.
</p>

English | [简体中文](./README_zh.md)

## Preparation

### Docker Deployment (Manual)

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

### Using Docker-compose

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

# Note: Adding Host Resolution!
➜  sudo vi /etc/hosts
# echo "127.0.0.1 go-ddd-layout.com" >> /etc/hosts

➜  docker ps -a
CONTAINER ID   IMAGE          COMMAND                   CREATED          STATUS          PORTS                                                  NAMES
a87de7ec1c88   nginx:latest   "/docker-entrypoint.…"   31 seconds ago   Up 30 seconds   0.0.0.0:80->80/tcp, :::80->80/tcp                      nginx-web
ccb49c2afaf9   go-ddd         "/bin/sh -c 'while !…"   31 seconds ago   Up 30 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp              go-app
7006bf52356a   mysql:latest   "docker-entrypoint.s…"   31 seconds ago   Up 30 seconds   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql-db
```

## API Documentation

swagger:

http://go-ddd-layout.com/swagger/index.html

## Directory Tree

```bash
.
├── .env                                    # Environment variable configuration file
├── .git                                    # Git version control directory
├── .gitignore                              # Git ignore file list
├── Dockerfile                              # Docker image build file
├── Makefile                                # Makefile for project building and management
├── README.md                               # Project documentation
├── application                             # Application layer directory
│   ├── event                               # Event-related code
│   └── service                             # Application service directory
│       └── user                            # User application service directory
│           └── application_service.go      # Implementation of user application service
├── build.sh                                # Build script
├── docker-compose.yml                      # Docker Compose configuration file
├── docs                                    # Documentation directory
│   ├── docs.go                             # Documentation generation code
│   ├── swagger.json                        # Swagger JSON file
│   └── swagger.yaml                        # Swagger YAML file
├── domain                                  # Domain layer directory
│   └── user                                # User domain directory
│       ├── entity                          # Domain entity directory
│       │   ├── role.go                     # Role entity definition
│       │   ├── user.go                     # User entity definition
│       │   └── valueobj                    # Value object directory
│       │       └── value_object.go         # Value object definition
│       ├── repository                      # Domain repository directory
│       │   ├── role.go                     # Role repository interface definition
│       │   └── user.go                     # User repository interface definition
│       └── service                         # Domain service directory
│           └── domain_service.go           # Implementation of domain service
├── go.mod                                  # Go module file
├── go.sum                                  # Go dependency version control file
├── infrastructure                          # Infrastructure layer directory
│   ├── common                              # Common code directory
│   │   ├── auth                            # Authentication-related code
│   │   │   └── header_token.go             # Implementation of header token authentication
│   │   ├── context                         # Context-related code
│   │   │   ├── getter.go                   # Context retrieval functions
│   │   │   └── store.go                    # Context storage functions
│   │   ├── cookie                          # Cookie operation code
│   │   │   └── cookie.go                   # Cookie operation implementation
│   │   ├── errors                          # Error handling code
│   │   │   └── errors.go                   # Custom error types and handling functions
│   │   ├── jwt                             # JWT-related code
│   │   │   ├── claims.go                   # JWT claims structure definition
│   │   │   └── token.go                    # JWT token generation and parsing implementation
│   │   ├── orm                             # ORM-related code
│   │   │   └── pagination.go               # Implementation of pagination queries
│   │   ├── random                          # Random number generation related code
│   │   │   └── generate_password.go        # Implementation of generating random passwords
│   │   └── response                        # Response handling code
│   │       └── response.go                 # Implementation of response handling functions
│   ├── database                            # Database-related code
│   │   ├── data.sql                        # Database initialization script
│   │   └── db.go                           # Database connection and operation implementation
│   ├── persistence                         # Persistence layer directory
│   │   ├── index.go                        # Index file
│   │   └── user                            # User persistence directory
│   │       ├── converter                   # Converter directory (Note: Conversion between domain objects and PO)
│   │       │   └── user.go                 # User converter implementation
│   │       ├── po                          # Persistence object directory
│   │       │   ├── role.go                 # Role persistence object definition
│   │       │   └── user.go                 # User persistence object definition
│   │       └── user_repo.go                # User persistence implementation
│   └── readme.md                           # Infrastructure layer documentation
├── interfaces                              # Interface layer directory (Note: Although Go officially discourages naming directories in plural form, 'interface' is a keyword, so it can be named in plural form here)
│   ├── adapter                             # Adapter directory for web frameworks and protocol integration, e.g., gin/fiber, etc.
│   │   ├── initialize                      # Initialization directory
│   │   │   ├── app.go                      # Application initialization implementation
│   │   │   ├── engine.go                   # Router engine implementation
│   │   │   └── service.go                  # Application service implementation
│   │   ├── middleware                      # Middleware directory
│   │   │   ├── auth.go                     # Authentication middleware implementation
│   │   │   ├── cors.go                     # Cross-origin resource sharing (CORS) middleware implementation
│   │   │   ├── error.go                    # Error handling middleware implementation
│   │   │   └── logging.go                  # Logging middleware implementation
│   │   └── router                          # Router configuration directory
│   │       └── router.go                   # Router configuration implementation
│   ├── assembler                           # Assembler directory (Note: Conversion between domain objects and DTO)
│   │   └── user                            # User assembler directory
│   │       └── user.go                     # User assembler implementation
│   ├── controller                          # Controller directory
│   │   ├── enter.go                        # Entry controller
│   │   ├── public                          # Public controller directory
│   │   │   └── hello_world.go              # HelloWorld public controller implementation
│   │   ├── sys                             # System controller directory
│   │   │   ├── enter.go                    # System entry controller
│   │   │   ├── login.go                    # Login controller implementation
│   │   │   ├── logout.go                   # Logout controller implementation
│   │   │   └── menu.go                     # Menu controller implementation
│   │   └── user                            # User controller directory
│   │       ├── create_user.go              # Create user controller implementation
│   │       ├── delete_user.go              # Delete user controller implementation
│   │       ├── enter.go                    # User entry controller
│   │       ├── update_password.go          # Update password controller implementation
│   │       ├── update_user.go              # Update user controller implementation
│   │       ├── user_detail.go              # User detail controller implementation
│   │       ├── user_info.go                # User info controller implementation
│   │       └── user_list.go                # User list controller implementation
│   └── dto                                 # Data Transfer Object (DTO) directory
│       ├── sys                             # System DTO directory
│       │   ├── login.go                    # Login DTO definition
│       │   └── menu.go                     # Menu DTO definition
│       └── user                            # User DTO directory
│           └── user.go                     # User DTO definition
├── main.go                                 # Main entry file
└── nginx.conf                              # Nginx configuration file
```

## Other

An article was written about this project [here](https://juejin.cn/post/7298160530292703244)

## Project Extension

If your project places a strong emphasis on security and deals with sensitive information such as JWT private keys and database passwords, there are still potential security risks associated with storing them fixedly in environment variables.

To further enhance security, we can adopt more secure approaches to handle these sensitive pieces of information:

- Use Key Management Systems (KMS): For example, AWS KMS, Google Cloud KMS, or HashiCorp Vault. Obtain the private keys from these systems through APIs.

- Use Public/Private Key Pairs: Consider using asymmetric encryption algorithms such as RSA or ECDSA instead of symmetric encryption algorithms. Store the private key in a secure location while the public key can be used within the application.

Regardless of the approach chosen, it is essential to follow some basic principles, including restricting access to the passwords, regularly rotating the passwords, and monitoring and auditing their usage.

For key rotation and management, Redis can be a viable choice. Here are some considerations for implementing key rotation when using Redis:

- Storing Keys: Store the currently valid key in Redis using the string data type provided by Redis.

- Regular Rotation: Generate a new key based on a predetermined time interval and store the new key in Redis.

- Smooth Transition Period: During the transition period before the old key expires, continue accepting and parsing tokens issued with both the old and new keys to ensure the ability to handle tokens issued with the old key.

- Update Token Issuance: When generating new tokens, sign them using the new key stored in Redis.

- Revoke Old Keys: Once the transition period ends, remove the old key from Redis.

- Periodic Execution: Set up a scheduled task or use a timer to perform the key rotation operation at regular intervals.

The benefits of using Redis as a key storage include high performance, reliability, and flexibility. However, it is crucial to ensure the security and reliability of the Redis instance. Take appropriate measures to protect the Redis instance, such as using Access Control Lists (ACLs) or access passwords to prevent unauthorized access. Additionally, backup and monitor the Redis instance to ensure data integrity and availability.
