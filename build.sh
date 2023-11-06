#!/bin/bash

# 执行 docker_build 目标
make docker_build

# 配置本地host域名解析（可能需要sudo权限）
echo "127.0.0.1 go-ddd-layout.com" >> /etc/hosts

#EOF