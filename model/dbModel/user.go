package dbModel

import (
	"errors"
	"gin-admin-back/initialize"
	"gin-admin-back/tools"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"uuid"`
	UserName   string    `json:"userName"`
	PassWord   string    `json:"passWord"`
	NickName   string    `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg  string    `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	//Propertie                //	多余属性自行添加
	//PropertieId uint  // 自动关联 Propertie 的Id 附加属性过多 建议创建一对一关系
}

func NewUser(user User) *User {
	return &User{UserName: user.UserName, PassWord: user.PassWord, NickName: user.NickName, HeaderImg: user.HeaderImg}
}

func (u *User) Create() (err error, userInter *User) {
	var user User
	dbErr := initialize.DEFAULTDB.Where("user_name = ?", u.UserName).First(&user).Error
	if dbErr != nil {
		return errors.New("用户名已注册"), nil
	} else {
		u.PassWord = tools.MD5V(u.PassWord)
		u.UUID = uuid.NewV4()
		err = initialize.DEFAULTDB.Create(u).Error
	}
	return err, u
}

func (u *User) Update() (err error, userInter *User) {
	//var user User
	//err = initialize.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.UserName, u.PassWord).First(&user).Update(u).Error
	err = initialize.DEFAULTDB.Create(u).Error
	return err, u
}

//func (u *User) Read() (err error, userInter *User) {
//	err = initialize.DEFAULTDB.Create(u).Error
//	return err, u
//}
//func (u *User) Delete() (err error, userInter *User) {
//	err = initialize.DEFAULTDB.Create(u).Error
//	return err, u
//}

func (u *User) Login() (err error, userInter *User) {
	var user User
	u.PassWord = tools.MD5V(u.PassWord)
	err = initialize.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.UserName, u.PassWord).First(&user).Error
	return err, &user
}

func (u *User) ChangePassword(newPassWord string) (err error, userInter *User) {
	var user User
	u.PassWord = tools.MD5V(u.PassWord)
	err = initialize.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.UserName, u.PassWord).First(&user).Update("pass_word", tools.MD5V(newPassWord)).Error
	return err, &user
}
