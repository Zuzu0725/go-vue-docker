package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Books []Book

// 一覧取得
func (b *Books) GetBooks() *gorm.DB {
	return DB.Find(&b)
}

// idに該当する書籍取得
func (b *Book) GetBook(id int) *gorm.DB {
	return DB.First(&b, id)
}

// 登録
func (b *Book) CreateBook() *gorm.DB {
	return DB.Create(&b)
}

// 更新
func (b *Book) UpdateBook() *gorm.DB {
	return DB.Model(&b).Updates(b)
}

// 削除
func (b *Book) DeleteBook(id int) *gorm.DB {
	return DB.Delete(&b, id)
}
