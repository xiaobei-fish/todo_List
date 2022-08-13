package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"unique"`
	PasswordStr string
}

const (
	SecretLevel = 12
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), SecretLevel)
	if err != nil {
		return err
	}
	user.PasswordStr = string(fromPassword)

	return nil
}

// CheckPassword 校验密码
func (user User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordStr), []byte(password))
	if err != nil {
		return false
	}
	return true
}
