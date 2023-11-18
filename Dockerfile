# 使用轻量的 Alpine 镜像作为基础镜像
FROM daocloud.io/library/ubuntu:20.04

# 设置工作目录
WORKDIR /app

# 复制本地编译好的可执行文件到工作目录
COPY . .



