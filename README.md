# SBCC - SB 控制中心

> **SBCC** = **SB** **C**ontrol **C**enter = **SB 控制中心**

一个基于 Go 语言构建的高性能、模块化控制中心后端框架。

---

## 📁 项目结构

```
sbcc/
├── .vscode/          # VS Code 配置文件
├── modbus/           # 主应用模块
│   ├── env/          # 环境配置系统
│   ├── gin/          # Gin 框架示例模块（演示 Chi + Gin 混合使用）
│   ├── grom/         # GORM ORM 数据库层
│   ├── home/         # 主页模块
│   ├── main/         # 程序入口
│   ├── sql/          # 原生 SQL 数据库连接层
│   ├── sqlx/         # SQLX 增强数据库层
│   ├── sub/          # 订阅 API 模块
│   ├── web/          # Web 引擎底座（Chi 路由）
│   ├── go.mod        # Go 依赖管理
│   └── go.sum        # 依赖校验文件
├── .gitignore        # Git 忽略配置
├── LICENSE           # 许可证文件
├── README.md         # 项目说明文档
├── build.bat         # Windows 构建脚本
└── go.work           # Go 工作区配置
```

---

## 📋 目录作用说明

| 目录 | 作用 | 状态 |
|------|------|------|
| `modbus/main/` | 程序入口，负责启动所有模块 | ✅ 核心 |
| `modbus/web/` | **Web 引擎底座**，使用 Chi 路由库，提供中间件、限流、日志等基础能力 | ✅ 核心 |
| `modbus/env/` | **环境配置系统**，自动创建 `.env` 文件，支持环境变量覆盖 | ✅ 核心 |
| `modbus/sql/` | **原生 SQL 连接层**，支持 SQLite/MySQL/PostgreSQL，带自动重连机制 | ✅ 核心 |
| `modbus/sqlx/` | **SQLX 增强层**，在原生连接基础上包装，支持命名参数等高级特性 | ✅ 扩展 |
| `modbus/grom/` | **GORM ORM 层**，提供对象关系映射能力，支持事务、预加载等 | ✅ 扩展 |
| `modbus/home/` | **主页模块**，处理根路径 `/` 的请求，演示如何挂载子路由 | ✅ 示例 |
| `modbus/sub/` | **订阅 API 模块**，提供 Clash 配置订阅服务，支持流量统计、到期时间等 | ✅ 业务 |
| `modbus/gin/` | **Gin 混合示例**，演示如何在 Chi 主路由上挂载 Gin 子应用 | 🧪 实验 |

---

## 🔥 核心特性

- **多数据库支持**: SQLite / MySQL / PostgreSQL 一键切换
- **三层数据库架构**: 原生 SQL → SQLX → GORM，按需选择
- **Chi + Gin 混合**: 支持在同一个项目中使用两个框架
- **自动配置**: `.env` 文件自动生成，无需手动创建
- **优雅启动**: 端口占用检测、模块状态反馈
- **高可用设计**: 数据库自动重连机制

---

## 🚀 快速开始

```bash

# 启动应用
go run modbus/main/run.go
```

启动后访问: http://localhost:9081

---

## 📝 配置说明

启动后会自动创建 `data/.env` 文件，主要配置项：

| 配置项 | 默认值 | 说明 |
|--------|--------|------|
| `WEB_PORT` | 9081 | Web 服务端口 |
| `DB_TYPE` | sqlite | 数据库类型 (sqlite/mysql/pgsql) |
| `DB_NAME` | data/sbcc.db | 数据库名/文件路径 |
| `DB_HOST` | 127.0.0.1 | 数据库主机 |
| `DB_PORT` | 5432 | 数据库端口 |
| `DB_USER` | postgres | 数据库用户名 |
| `DB_PASSWORD` | your_password | 数据库密码 |

---

## 🛠️ 技术栈

- **语言**: Go 1.26+
- **Web 框架**: [Chi](https://github.com/go-chi/chi)
- **ORM**: [GORM](https://gorm.io/)
- **SQL 增强**: [SQLX](https://github.com/jmoiron/sqlx)
- **限流**: [httprate](https://github.com/go-chi/httprate)

---

## 📄 许可证

Apache License 2.0
