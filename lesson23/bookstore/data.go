package main

import (
	"context"
	"errors"
	"log"

	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Shelf{}, &Book{})
	return db, err
}

type Shelf struct {
	gorm.Model
	ID    int64  `gorm:"primaryKey"`
	Theme string `json:"theme" gorm:"type:varchar(100)"`
	Size  int64  `json:"size"`

	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct {
	ID      int64 `gorm:"primaryKey"`
	Author  string
	Title   string
	ShelfID int64

	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

// CreateShelf 创建书架
func (bs *bookstore) CreateShelf(ctx context.Context, shelf *Shelf) (*Shelf, error) {
	if err := bs.db.WithContext(ctx).Create(shelf).Error; err != nil {
		log.Printf("Failed to create shelf: %v", err)
		return nil, err
	}

	return shelf, nil
}

// GetShelf 获取书架
func (bs *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	var shelf Shelf
	if err := bs.db.WithContext(ctx).First(&shelf, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("shelf not found")
		}
		log.Printf("Failed to get shelf: %v", err)
		return nil, err
	}

	return &shelf, nil
}

// ListShelves 列出所有书架
func (bs *bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {
	var shelves []*Shelf
	if err := bs.db.WithContext(ctx).Find(&shelves).Error; err != nil {
		log.Printf("Failed to list shelves: %v", err)
		return nil, err
	}

	return shelves, nil
}

// DeleteShelf 删除书架
func (bs *bookstore) DeleteShelf(ctx context.Context, id int64) error {
	result := bs.db.WithContext(ctx).Delete(&Shelf{}, id)
	if result.Error != nil {
		log.Printf("Failed to delete shelf: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("shelf not found")
	}

	return nil
}

// GetBookListByShelfID 根据书架id查询图书
func (b *bookstore) GetBookListByShelfID(ctx context.Context, shelfID int64, cursor int64, pageSize int) ([]*Book, error) {
	var vl []*Book
	err := b.db.Debug().WithContext(ctx).Where("shelf_id = ? and id > ?", shelfID, cursor).Order("id asc").Limit(pageSize).Find(&vl).Error
	return vl, err
}
