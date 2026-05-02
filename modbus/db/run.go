package db

import (
	"database/sql"
	"fmt"
	"modbus/env"
	"time"

	_ "github.com/go-sql-driver/mysql" // 换成你实际使用的驱动
)

var DB *sql.DB

func Run() {
	// 一次性初始化数据库配置
	env.Init([][]string{
		{"DB_TYPE", "sqlite", "--- [数据库驱动配置] ---", "支持类型: sqlite (本地文件), pgsql/mysql (远程服务器)", "注意: 若选 sqlite，下方 DB_NAME 则为数据库文件名"},
		{"DB_NAME", "data/sbcc.db", "--- [数据库实例/文件名] ---", "SQLite 示例: data/sbcc.db | PG/MySQL 示例: my_trade_db"},
		{"DB_HOST", "127.0.0.1", "--- [网络配置] ---", "仅针对 pgsql/mysql，sqlite 模式下请忽略"},
		{"DB_PORT", "5432", "默认端口: pgsql(5432), mysql(3306)"},
		{"DB_USER", "postgres", "数据库用户名"},
		{"DB_PASSWORD", "your_password", "数据库密码"},
	})

	dsn := "123456@tcp(127.0.0.1:3306)/my_db?parseTime=True"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("⚠️ [DB] 数据库模块 配置错误！")
	}

	// 间隔2s
	time.Sleep(2 * time.Second)

	// 开始无限循环重连
	for {

		err = DB.Ping()
		if err == nil {
			fmt.Printf("✅ [DB] 数据库 连接成功！\n")
			break // 连上了，跳出循环
		}

		fmt.Printf("🔄 [DB] 数据库连接失败: %v。 9秒后尝试重连...\n", err)
		time.Sleep(9 * time.Second)
		// fmt.Printf("❌ [DB] 数据库 重连中...\n")
	}

	// 设置连接池（连上后再设置）
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
}
