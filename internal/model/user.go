package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;unique_index:uk_user_username" json:"username"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	Salt      string `gorm:"type:varchar(40);not null" json:"salt"`
	Nickname  string `gorm:"type:varchar(20);not null;index:idx_user_nickname" json:"nickname"`
	AvatarURL string `gorm:"type:varchar(255);not null;default:''" json:"avatar"`
	Status    uint8  `gorm:"type:tinyint(1);not null;default:'0'" json:"status"`
	Deleted   int    `gorm:"type:bigint;not null;unique_index:uk_user_username;default:'0'" json:"deleted"`
}

func (a User) GetHashedPassword(db *gorm.DB) string {
	var user User
	db.Where("username = ? and deleted = ?", a.Username, 0).First(&user)
	return user.Password
}

func (a User) GetByID(db *gorm.DB) (*User, error) {
	result := db.Where("ID = ? and deleted = ?", a.ID, 0).First(&a)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("user does not exist")
	}
	return &a, nil
}

func (a User) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a User) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("deleted = ?", 0).Updates(values).Error
}

func (a User) Delete(db *gorm.DB) error {
	return db.Model(&a).Where("deleted = ?", 0).Update("deleted", a.ID).Error
}

func (a User) GetID(db *gorm.DB) uint {
	var user User
	db.Where("username = ?  and deleted = ?", a.Username, 0).First(&user)
	if user.ID > 0 {
		return user.ID
	}
	return 0
}
