package model

import (
	"gogin/pkg/auth"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

type User struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
	Username  string     `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password  string     `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// Validate the fields.
func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}

// Encrypt the user password.
func (user *User) Encrypt() (err error) {
	user.Password, err = auth.Encrypt(user.Password)
	return
}

func (user *User) Create() error {
	return DB.Create(&user).Error
}

func List() []string {
	users := []User{}
	DB.Find(&users)
	var userlist []string
	for _, v := range users {
		userlist = append(userlist, v.Username)
	}
	//strings.Replace(strings.Trim(fmt.Sprint(userlist), "[]"), " ", ",", -1)
	return userlist
}
