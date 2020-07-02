package models

import (
	"time"

	"github.com/mhdiiilham/gournal/db"
)

// Admin models
type Admin struct {
	ID             uint64 `json:"id" gorm:"primary_key;unique;AUTO_INCREMENT"`
	Fullname       string `json:"fullname" gorm:"type:varchar(50)"`
	Email          string `json:"email" gorm:"type:varchar(100);unique_index"`
	PasswordHashed string `json:"password_hashed" gorm:"type:varchar(60)"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// AdminSignUp ...
type AdminSignUp struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminSignIn ...
type AdminSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*
|--------------------------------------------------------------------------
| Admin's Methods
|--------------------------------------------------------------------------
|
| Here's methods that usually used.
|
*/

// Save ...
func (a *Admin) Save() {
	db.DB().Save(a)
}

// First ...
func First(email string) (Admin, error) {
	var admin Admin
	err := db.DB().Where("email = ?", email).First(&admin)
	if err != nil {
		return admin, err.Error
	}
	return admin, nil
}
