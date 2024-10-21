package cli

import (
	"OceanCloud-Image/config"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Setup() {
	// 加载配置
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("加载配置时出错: %v", err)
	}
	var db *sql.DB
	if cfg.Database.Db_type == "mysql" {
		// 初始化 MySQL
		dbport := strconv.Itoa(cfg.Database.Port)
		sqlink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, dbport, cfg.Database.Db_name)
		db, err = sql.Open("mysql", sqlink)
		if err != nil {
			log.Fatalf("[ERROR] 数据库连接错误: %v", err)
		}
		defer db.Close()
	} else if cfg.Database.Db_type == "sqlite" {
		// 初始化 SQLite
		db, err = sql.Open("sqlite3", cfg.Database.Sqlite_file)
		if err != nil {
			log.Fatalf("[ERROR] 数据库连接错误: %v", err)
		}
		defer db.Close()
	} else {
		log.Fatalf("[ERROR] 不支持的数据库类型: %s。请检查配置文件。", cfg.Database.Db_type)
	}

	// 检查 app_image 表是否存在
	var count int

	// 检查表是否存在
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'app_image'").Scan(&count)
	if err != nil {
		log.Fatalf("[ERROR] 查询表存在性时出错: %v", err)
	}

	// 如果表存在，查询记录数
	if count > 0 {
		// 查询表记录数
		err = db.QueryRow("SELECT COUNT(*) FROM app_image").Scan(&count)
		if err != nil {
			log.Fatalf("[ERROR] 查询 app_image 表记录数时出错: %v", err)
		}
	}

	// 如果表不存在或为空，导入数据
	if count == 0 {
		// 导入 ./ocean.sql 文件
		importSQLFile(db, "./ocean.sql")
	} else {
		fmt.Println("您已安装过，无需执行安装")
	}
}

func importSQLFile(db *sql.DB, filePath string) {
	// 打开 SQL 文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("[ERROR] 读取 SQL 文件时出错: %v", err)
	}
	defer file.Close()

	// 使用 bufio.Scanner 逐行读取文件
	scanner := bufio.NewScanner(file)
	var sqlStatement strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		// 跳过空行和注释
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "--") {
			continue
		}
		// 将当前行添加到 SQL 语句中
		sqlStatement.WriteString(line + "\n")

		// 如果当前行是以分号结尾，执行 SQL 语句
		if strings.HasSuffix(line, ";") {
			_, err := db.Exec(sqlStatement.String())
			if err != nil {
				log.Fatalf("[ERROR] 执行 SQL 导入时出错: %v", err)
			}
			// 清空 SQL 语句构建器
			sqlStatement.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("[ERROR] 读取 SQL 文件时出错: %v", err)
	}

	log.Println("数据库导入成功！")
}
