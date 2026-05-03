# 构建阶段
FROM golang:1.26-alpine AS builder

WORKDIR /app

# 安装 git（某些依赖需要）
RUN apk add --no-cache git

# 复制 modbus 目录下的所有文件（包含 go.mod）
COPY modbus/ .

# 下载依赖
RUN go mod download

# 构建可执行文件（go.mod 在当前目录，所以路径是 ./main/run.go）
RUN CGO_ENABLED=0 GOOS=linux go build -o sbcc ./main/run.go

# 运行阶段
FROM alpine:3.19

WORKDIR /app

# 安装 CA 证书（用于 HTTPS 请求）
RUN apk add --no-cache ca-certificates tzdata

# 从构建阶段复制可执行文件和静态资源
COPY --from=builder /app/sbcc .
COPY --from=builder /app/static ./static
COPY --from=builder /app/data ./data

# 创建必要目录
RUN mkdir -p /app/data

# 暴露端口
EXPOSE 9081

# 设置时区
ENV TZ=Asia/Shanghai

# 启动命令
CMD ["./sbcc"]