package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const fileName = "data/.env"

// Init 初始化配置文件
func Init(config [][]string) {
	// 1. 确保目录存在（创建 data 文件夹）
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("创建目录失败:", err)
			return
		}
	}

	// 2. 读取现有内容（如果文件不存在，contentStr 为空，不影响后续判断）
	existingContent, _ := os.ReadFile(fileName)
	contentStr := string(existingContent)

	// 3. 打开文件（如果不存在则创建，存在则准备追加）
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开/创建文件失败:", err)
		return
	}
	defer f.Close()

	// 4. 遍历配置数组
	for i := 0; i < len(config); i++ {
		line := config[i]
		comment := line[0] // 注释
		key := line[1]     // 参数名
		val := line[2]     // 数据内容

		// 检查 line[1] (Key) 是否已存在
		if strings.Contains(contentStr, key+"=") {
			// 如果已有该参数，跳过不处理
			continue
		}

		// 如果没有，则在文件尾部顺序写入
		// 按照你要求的格式：#注释 \n Key=Value \n\n
		data := fmt.Sprintf("#%s\n%s=%s\n\n", comment, key, val)

		_, err := f.WriteString(data)
		if err != nil {
			fmt.Printf("追加参数 [%s] 失败: %v\n", key, err)
		}
	}
}

// Get 读取值：环境变量优先级高于配置文件
func Get(key string) string {
	// 1. 【新增】首先尝试从系统环境变量获取
	// 比如你在 Linux 下执行: WEB_PORT=8080 ./main
	if envVal := os.Getenv(key); envVal != "" {
		return envVal
	}

	// 2. 如果环境变量没有，再打开文件读取
	f, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// 过滤注释
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// 匹配 Key=
		if strings.HasPrefix(line, key+"=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return parts[1]
			}
		}
	}

	return ""
}
