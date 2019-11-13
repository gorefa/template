/**
* @file   : user
* @descrip: 
* @author : ch-yk
* @create : 2018-09-05 下午1:05
* @email  : commonheart.yk@gmail.com
**/

package model

import (
	"fmt"
	"sync"
	"time"

	"api_gateway/internal/auth"
	"api_gateway/internal/constvar"

	"gopkg.in/go-playground/validator.v9"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

/* 查看 UserList 列表的时候:
type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
*/
type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

/**/

// Token represents a JSON web token.
// 只有登录成功才有 token 返回
type Token struct {
	Token string `json:"token"`
}


// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "api_users"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
//如果模型有DeletedAt字段，它将自动获得软删除功能！
//那么在调用Delete时不会从数据库中永久删除，而是只将字段DeletedAt的值设置为当前时间。
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser gets an user by the username.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

/* ListUser List all users，需要传递的信息 username 以及分页选项 (offset, limit)
默认传递 {"offset": 0, "limit": 20} 即可
如果有传递条件，那么根据条件查询
*/
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	//查询 count
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}


	err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id asc").Find(&users).Error
	if err != nil {
		return users, count, err
	}
	//log.Infof("offset = %d ; ListUser 本次查询到了%d", offset, len(users))
	return users, count, nil
}

// Compare with the plain text password.
// Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
