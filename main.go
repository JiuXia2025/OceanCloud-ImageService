package main
//Main
//Create by JiuXia2025 <2025226181@qq.com> <Blog.InekoXia.COM>
import (
	"OceanCloud-Image/api"
	"OceanCloud-Image/cli"
	"OceanCloud-Image/config"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	// 检查参数数量
	if len(os.Args) < 2 {
		// 默认直接启动就是启动Web的APIServer服务
		_start()
		return
	}

	// 获取第一个参数
	command := os.Args[1]

	// 根据参数执行不同的逻辑
	switch command {
	case "listImage":
		cli.ListImages()
	case "setup":
		cli.Setup()
	case "--help":
		cli.PrintHelp()
	case "start":
		_start()
	case "createImage":
		if len(os.Args) != 5 {
			fmt.Println("用法: ./ocimage createImage <镜像类型:{centos,ubuntu,debian,windows}> <镜像路径> <镜像显示名称>")
			return
		}
		cli.CreateImage(os.Args[2], os.Args[3], os.Args[4])
	case "deleteImage":
		if len(os.Args) != 3 {
			fmt.Println("用法: ./ocimage deleteImage <镜像ID>")
			return
		}
		cli.DeleteImage(os.Args[2])
	case "infoImage":
		if len(os.Args) != 3 {
			fmt.Println("用法: ./ocimage infoImage <镜像ID>")
			return
		}
		cli.InfoImage(os.Args[2])
	default:
		fmt.Println("未知命令：'", command, "'，请使用 --help 查看命令帮助列表")
	}
}

func _start() {
	// 加载配置
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("加载配置时出错: %v", err)
	}
	config.UpdateCheck()

	var db *sql.DB
	if cfg.Database.Db_type == "mysql" {
		// 初始化 MySQL
		dbport := strconv.Itoa(cfg.Database.Port)
		sqlink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, dbport, cfg.Database.Db_name)

		if cfg.Debug == true {
			fmt.Println("[Warning] 调试模式已开启，某些用于调试的数据将会输出导致安全性降低，如需关闭请前往 'config.yaml'")
			fmt.Println("[INFO] 正在连接到MySQL数据库:", sqlink)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		db, err = sql.Open("mysql", sqlink)
		if err != nil {
			log.Fatalf("[ERROR] 数据库连接错误: %v", err)
		}
		defer db.Close()

	} else if cfg.Database.Db_type == "sqlite" {
		// 初始化 SQLite
		fmt.Println("[INFO] 正在连接到SQlite3数据库：", cfg.Database.Sqlite_file)
		db, err = sql.Open("sqlite3", cfg.Database.Sqlite_file)
		if err != nil {
			log.Fatalf("[ERROR] 数据库连接错误: %v", err)
		}
		defer db.Close()
		fmt.Println("[INFO] 成功连接到 SQLite3 数据库！")

	} else {
		log.Fatalf("[ERROR] 不支持的数据库类型: %s。请检查配置文件。", cfg.Database.Db_type)
	}
	fmt.Println("[INFO] 正在准备启动APIServer")
	// 初始化路由
	router := api.ImagesRouter(db)

	// 启动服务器
	go func() {
		fmt.Println("[INFO] APIServer已启动于localhost:8080，如需停止请按下Ctrl+C")
		fmt.Println("[INFO] 当前运行于WebAPI模式，如需使用命令模式请在终端执行 ./ocimage --help 查看帮助")
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("[ERROR] 服务器启动失败: %v", err)
		}
	}()
	// 优雅的关闭APIServer
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("[STOP] 正在关闭服务器...")
	if err := db.Close(); err != nil {
		log.Fatalf("[ERROR] 关闭数据库连接失败: %v", err)
	}
	fmt.Println("[STOP] 服务器成功关闭...")
}
