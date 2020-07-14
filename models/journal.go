package models

import (
	"strconv"
	"time"
)

// Journal type
type Journal struct {
	ID          uint64    `json:"id" gorm:"primary_key;unique;AUTO_INCREMENT"`
	Title       string    `json:"title" gorm:"type:varchar(25)"`
	Image       Image     `json:"image" gorm:"type:varchar(50)"`
	ImageID     uint64    `json:"image_id"`
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

// ListJournal ...
type ListJournal []Journal

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
| Journal's Methods
|--------------------------------------------------------------------------
|
| Here's methods that usually used.
|
*/

// Save ...
func (j *Journal) Save() {
	DB.Save(j)
}

// First ...
func (j *Journal) First(id string) {
	DB.Where("id = ?", id).Preload("Image").First(j)
}

// Find function
/*
	This is a variadic function.
	1. If the args is empty when
		invoking this Function,
		it will return 5 Journal's
	1. If the args is given:
		the first one is field, and the
		second one is it value.
		Return list of journal that match.
		eg:
			var journal Journal
			journal.Find("title", "Hello")

	ps:
		FILTERING IS NOT WORKING, YET!
		1. Always ordered by "created_at"
*/
func (j *ListJournal) Find(page string, args ...string) {
	var limit uint64
	limit = 5
	pages, _ := strconv.ParseUint(page, 10, 64)
	skip := pages - 1
	offset := skip * limit

	switch len(args) {
	case 2:
		DB.Offset(offset).Limit(limit).Preload("Image").Find(j)
		break
	case 0:
		DB.Offset(offset).Limit(limit).Preload("Image").Find(j)
		break
	}
}

/*
|--------------------------------------------------------------------------
| Image's Methods
|--------------------------------------------------------------------------
|
| Here's methods that usually used.
|
*/

// Save Image
func (i *Image) Save() {
	DB.Save(&i)
}

// Find Image
func (i *Image) Find(id string) {
	DB.Where("id = ?", id).First(&i)
}
