package cli

//ImageManager
//Create by JiuXia2025 <2025226181@qq.com> <Blog.InekoXia.COM>

import (
	"OceanCloud-Image/config"
	"OceanCloud-Image/database"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func ListImages() {
	fmt.Println("数据库中所有的镜像：")
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

	images, err := database.GetAppImages(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, img := range images {
		fmt.Printf("镜像ID: %d, 类型: %s, 目录: %s, 名称: %s\n", img.ID, img.Type, img.Path, img.Name)
	}

}

func CreateImage(imageType, imagePath, displayName string) {
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
	// 查询当前最大 ID
	var maxID int
	err = db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM app_image").Scan(&maxID)
	if err != nil {
		fmt.Errorf("查询最大 ID 时出错: %w", err)
	}
	newID := maxID + 1 // 计算新的 ID

	// 创建新的 AppImage
	newImage := config.AppImage{ID: newID, Type: imageType, Path: imagePath, Name: displayName}
	_, err = database.AddAppImage(db, newImage)
	if err != nil {
		fmt.Errorf("添加新的 AppImage 错误: %w", err)
	}
	fmt.Printf("已添加新的镜像，ID: %d\n，镜像名：%s", newID, displayName)
}

func DeleteImage(imageID string) {
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
	fmt.Printf("正在删除镜像 ID: %s\n", imageID)
	// 删除镜像的逻辑
	// 执行删除操作
	id, err := strconv.ParseInt(imageID, 10, 64)
	if err := database.DeleteAppImage(db, id); err != nil {
		log.Fatal(err)
	}
	fmt.Println("删除 AppImage 完成，执行./ocimage listImage查看列表")
}

func InfoImage(imageID string) {
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
	id, err := strconv.ParseInt(imageID, 10, 64)
	image, err := database.InfoAppImage(db, id)
	if err != nil {
		log.Fatal(err)
	}

	// 打印查询结果
	fmt.Printf("镜像信息: %+v\n", image)

}
