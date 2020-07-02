package models

import (
	"time"

	"github.com/mhdiiilham/gournal/db"
)

// Journal type
type Journal struct {
	ID          uint64    `json:"id" gorm:"primary_key;unique;AUTO_INCREMENT"`
	Title       string    `json:"title" gorm:"type:varchar(25)"`
	Image       Image     `json:"image" gorm:"type:varchar(50)"`
	ImageID     uint64	  `json:"image_id"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}


// JournalInput from user
type JournalInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgurID     string `json:"imgur_id"`
}

// Image saver
type Image struct {
	ID         uint64    `json:"id" gorm:"primary_key;unique;AUTO_INCREMENT"`
	ImgurID    string    `json:"imgur_id" gorm:"type:varchar(10)"`
	Link       string    `json:"link" gorm:"type:varchar(50)"`
	DeleteHash string    `json:"deletehash" gorm:"type:varchar(25)"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

/*
|--------------------------------------------------------------------------
| Jorunal's Methods
|--------------------------------------------------------------------------
|
| Here's methods that usually used.
|
*/

// Save ...
func (j *Journal) Save() {
	db.DB().Save(j)
}
// First ...
func (j *Journal) First(id string) {
	db.DB().Where("id = ?", id).Preload("Image").First(j)
}

// Save Image
func (i *Image) Save() {
	db.DB().Save(&i)
}

// Find Image
func (i *Image) Find(id string) {
	db.DB().Where("id = ?", id).First(&i)
}
