package api

import (
	"OceanCloud-Image/database"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImagesRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// 获取所有 AppImages 的 API
	router.GET("/api/images", func(c *gin.Context) {
		images, err := database.GetAppImages(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, images)
	})

	// 新增 /api/info 路由
	router.GET("/api/info", func(c *gin.Context) {
		info := gin.H{
			"version":     "1.0.0",
			"description": "This is an API for managing AppImages.",
		}
		c.JSON(http.StatusOK, info)
	})

	return router
}

//func APITest() {
//	cfg, err := config.LoadConfig("config.yaml")
//	if err != nil {
//		log.Fatalf("error loading config: %v", err)
//	}
//	//初始化mysql
//	dbport := strconv.Itoa(cfg.Database.Port)
//	sqlink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, dbport, cfg.Database.Db_name)
//	fmt.Println("[INFO]连接至数据库：", sqlink)
//	//调自有函数的方式连数据库（弃用）：
//	//if err := database.InitDB(sqlink); err != nil {
//	//	log.Fatal(err)
//	//}
//	db, err := sql.Open("mysql", sqlink)
//	if err != nil {
//		log.Fatal("[ERROR]数据库连接错误：", err)
//	}
//	defer db.Close()
//	fmt.Println("[INFO]开始连接到 MySQL 数据库！")
//	router := gin.Default()
//
//	// 获取所有 AppImages 的 API
//	router.GET("/api/images", func(c *gin.Context) {
//		images, err := database.GetAppImages(db)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, images)
//	})
//
//	// 启动服务器
//	if err := router.Run(":8080"); err != nil {
//		log.Fatal(err)
//	}
//}
