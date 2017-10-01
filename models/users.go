package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id       int64
	Username string `orm:"size(50)"`
	Password string `orm:"size(50)"`
}

type UserSuccessJson struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func CreateUser(user Users) (UserSuccessJson, ErrorMessage) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("users")
	num, err := qs.Filter("username", user.Username).Count()

	var errorJson ErrorMessage
	var userJson UserSuccessJson

	// 如果没有重复用户
	if err == nil {
		if num == 0 {
			id, err := o.Insert(&user)
			if err == nil {
				userJson.Username = user.Username
				userJson.Id = id

				fmt.Println("注册成功", userJson)
				return userJson, errorJson
			}
		}

		errorJson.Message = "用户名已被注册"
		return userJson, errorJson
	}

	errorJson.Message = "注册错误，稍后重试"
	return userJson, errorJson
}

func Login(user Users) string {
	o := orm.NewOrm()
	o.Using("default")

	var ouser Users

	err := o.QueryTable("users").Filter("username", user.Username).One(&ouser)

	if err == nil {
		if ouser.Password != user.Password {
			return "密码错误"
		}

		return "登录成功"
	}
	return "用户不存在"
}
