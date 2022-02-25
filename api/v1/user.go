package v1

import (
	"fmt"
	"gin-admin-back/model/dbModel"
	"github.com/gin-gonic/gin"
)

type RegistAndLoginStuct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

func Register(c *gin.Context) {
	var R RegistAndLoginStuct
	_ = c.BindJSON(&R)
	U := dbModel.User{UserName: R.UserName, PassWord: R.PassWord}
	fmt.Println(U)
	err, user := U.Create()
	fmt.Println(err, user)
}
