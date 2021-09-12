package week2

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetUser(db *gorm.DB, id string) (*User, error) {
	var user = &User{}
	err := db.Table("users").Where("id = ?", id).First(user).Error
	return user, errors.Wrap(err, fmt.Sprintf("user(id:%s not found", id))
}
