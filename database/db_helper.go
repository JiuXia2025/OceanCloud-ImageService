package database
//DB_Helper
//Create by JiuXia2025 <2025226181@qq.com> <Blog.InekoXia.COM>
import (
	"OceanCloud-Image/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 连接数据库（弃用）
func InitDB(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

// 获取数据库中的 AppImages
func GetAppImages(db *sql.DB) ([]config.AppImage, error) {
	//func GetAppImages(db *sql.DB) ([]config.AppImage, error) {
	query := "SELECT id, type, path, name FROM app_image"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []config.AppImage
	for rows.Next() {
		var img config.AppImage
		if err := rows.Scan(&img.ID, &img.Type, &img.Path, &img.Name); err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}

func AddAppImage(db *sql.DB, img config.AppImage) (int64, error) {
	// 更新查询以插入 ID
	query := "INSERT INTO app_image (id, type, path, name) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, img.ID, img.Type, img.Path, img.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId() // 这里可以返回插入的 ID
}

// 更新现有的 AppImage
func UpdateAppImage(db *sql.DB, img config.AppImage) error {
	query := "UPDATE app_image SET type = ?, path = ?, name = ? WHERE id = ?"
	_, err := db.Exec(query, img.Type, img.Path, img.Name, img.ID)
	return err
}

// 删除 AppImage
func DeleteAppImage(db *sql.DB, id int64) error {
	query := "DELETE FROM app_image WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

type AppImage struct {
	ID   int
	Type string
	Path string
	Name string
}

func InfoAppImage(db *sql.DB, id int64) (AppImage, error) {
	query := "SELECT id, type, path, name FROM app_image WHERE id = ?"
	var image AppImage

	// 使用 QueryRow 执行查询并将结果扫描到 image 结构体
	err := db.QueryRow(query, id).Scan(&image.ID, &image.Type, &image.Path, &image.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return AppImage{}, fmt.Errorf("未找到 ID 为 %d 的镜像", id)
		}
		return AppImage{}, err
	}

	return image, nil
}

//
//// 查询用户
//func GetUsers() ([]User, error) {
//	rows, err := db.Query("SELECT id, name FROM app_user")
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var users []User
//	for rows.Next() {
//		var user User
//		if err := rows.Scan(&user.ID, &user.Name); err != nil {
//			return nil, err
//		}
//		users = append(users, user)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//	return users, nil
//}
//
//// 用户结构体
//type User struct {
//	ID   int
//	Name string
//}
