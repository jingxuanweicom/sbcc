package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 换成你实际使用的驱动
)

var DB *sql.DB

func Run() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_db?parseTime=True"

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
