package models

import (
	"time"

	"github.com/mhdiiilham/gournal/db"
)

// Journal type
type Journal struct {
	ID               uint `json:"id" gorm:"primary_key;unique;AUTO_INCREMENT"`
	Title          string `json:"tile" gorm:"type:varchar(10)"`
	ImageURL       string `json:"image_url" gorm:"type:varchar(20)"`
	Description    string `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
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
