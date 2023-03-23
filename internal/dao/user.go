package dao

import (
	"github.com/eroneko/gin-service/internal/model"
	"github.com/eroneko/gin-service/internal/service"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar"`
	Status    uint8  `json:"status"`
	Deleted   int    `json:"deleted"`
}

func (d *Dao) GetHashedPassword(username string) string {
	user := model.User{
		Username: username,
	}
	return user.GetHashedPassword(d.engine)
}

func (d *Dao) GetUserByID(params User) (service.GetUserResponse, error) {
	user := model.User{
		Model: gorm.Model{ID: params.ID},
	}
	result, err := user.GetByID(d.engine)
	if err != nil {
		return service.GetUserResponse{}, err
	}
	resp := service.GetUserResponse{
		ID:        result.ID,
		Username:  result.Username,
		Nickname:  result.Nickname,
		AvatarURL: result.AvatarURL,
	}
	return resp, nil
}

func (d *Dao) CreateUser(params User) error {
	user := model.User{
		Username:  params.Username,
		Password:  params.Password,
		AvatarURL: params.AvatarURL,
	}
	return user.Create(d.engine)
}

func (d *Dao) UpdateUser(params User) error {
	user := model.User{
		Model:     gorm.Model{ID: params.ID},
		Nickname:  params.Nickname,
		AvatarURL: params.AvatarURL,
		Password:  params.Password,
	}
	return user.Update(d.engine, user)
}

func (d *Dao) DeleteUser(params User) error {
	user := model.User{
		Model: gorm.Model{ID: params.ID},
	}
	return user.Delete(d.engine)
}

func (d *Dao) GetUserID(params User) uint {
	user := model.User{
		Username: params.Username,
	}
	return user.GetID(d.engine)
}

func (d *Dao) IsUserExist(param User) bool {
	user := model.User{
		Model:    gorm.Model{ID: param.ID},
		Username: param.Username,
	}
	if param.ID > 0 {
		_, err := user.GetByID(d.engine)
		if err != nil {
			return false
		}
		return true
	} else if user.GetID(d.engine) == 0 {
		return false
	}
	return true
}
