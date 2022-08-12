package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("読み込みに失敗しました")
	}
	// DB接続情報
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_DATABASE")

	// DBと接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", db_user, db_password, db_host, db_name)

	dialector := mysql.Open(dsn)

	if DB, err = gorm.Open(dialector); err != nil {
		fmt.Println("DBとの接続に失敗しました", err)
	}
}
