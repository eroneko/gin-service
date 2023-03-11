package api

import (
	"github.com/eroneko/gin-service/global"
	"github.com/eroneko/gin-service/internal/dao"
	"github.com/eroneko/gin-service/internal/model"
	"github.com/eroneko/gin-service/internal/service"
	"github.com/eroneko/gin-service/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (a User) Login(c *gin.Context) {
	d := dao.New(global.DBEngine)
	var req service.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid request body",
		})
		return
	}
	if !d.IsUserExist(dao.User{Username: req.Username}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not exist",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(d.GetHashedPassword(req.Username)), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password error",
		})
		return
	}
	token, err := app.ReleaseToken(model.User{
		Model: gorm.Model{
			ID: d.GetUserID(dao.User{Username: req.Username}),
		}})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "login success",
	})
	return
}

func (a User) Register(c *gin.Context) {
	d := dao.New(global.DBEngine)
	var req service.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid request body",
		})
		return
	}
	if d.IsUserExist(dao.User{Username: req.Username}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exist",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if err = d.CreateUser(dao.User{Username: req.Username, Password: string(hashedPassword)}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "create user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
	return
}

func (a User) Update(c *gin.Context) {
	d := dao.New(global.DBEngine)
	var req service.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid request body",
		})
		return
	}
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}
	currentUser, ok := c.Get("user")
	if currentUser.(service.GetUserResponse).ID != uint(userID) || !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You can only update your own information",
		})
		return
	}
	user := dao.User{
		ID:        currentUser.(service.GetUserResponse).ID,
		Password:  req.Password,
		Nickname:  req.Nickname,
		AvatarURL: req.AvatarURL,
	}
	if err := d.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "update user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user success",
	})
}

func (a User) Delete(c *gin.Context) {
	d := dao.New(global.DBEngine)
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}
	currentUser, ok := c.Get("user")
	if currentUser.(service.GetUserResponse).ID != uint(userID) || !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You can't destroy other's account",
		})
		return
	}
	user := dao.User{
		ID: uint(userID),
	}
	if !d.IsUserExist(user) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not exist",
		})
		return
	}
	if err := d.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
	return
}

func (a User) Info(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
	return
}
