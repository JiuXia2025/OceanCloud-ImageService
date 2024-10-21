package cli

import "fmt"

func PrintHelp() {
	fmt.Println()
	fmt.Println("----- OceanCloud-Glance_CLI 帮助列表 -----")
	fmt.Println("启动主服务：./ocimage start")
	fmt.Println("初始化数据库：./ocimage setup")
	fmt.Println("镜像列表：./ocimage listImage")
	fmt.Println("创建镜像：./ocimage createImage <镜像类型:{centos,ubuntu,debian,windows}> <镜像路径> <镜像显示名称>")
	fmt.Println("删除镜像：./ocimage deleteImage <镜像ID>")
	fmt.Println("查询镜像：./ocimage infoImage <镜像ID>")
	fmt.Println("-------------------------------------")
	fmt.Println()
}
