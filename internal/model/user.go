package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;uniqueIndex:uk_user_username" json:"username"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	Salt      string `gorm:"type:varchar(40);not null" json:"salt"`
	Nickname  string `gorm:"type:varchar(20);not null;index:idx_user_nickname" json:"nickname"`
	AvatarURL string `gorm:"type:varchar(255);not null" json:"avatar"`
	Status    uint8  `gorm:"type:tinyint(1);not null" json:"status"`
	Deleted   int    `gorm:"type:bigint;not null;uniqueIndex:uk_user_username" json:"deleted"`
}

func (a User) GetHashedPassword(db *gorm.DB) string {
	var user User
	db.Where("username = ?", a.Username).First(&user)
	return user.Password
}

func (a *User) Get(db *gorm.DB) (int64, error) {
	result := db.Where("ID = ?", a.ID).First(&a)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (a User) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a User) Update(db *gorm.DB) error {
	return db.Where("id = ?", a.ID).Update(&a).Error
}

func (a User) Delete(db *gorm.DB) error {
	return db.Delete(&a).Error
}

func (a User) GetUserID(db *gorm.DB) uint {
	var user User
	db.Where("username = ?", a.Username).First(&user)
	if user.ID > 0 {
		return user.ID
	}
	return 0
}
