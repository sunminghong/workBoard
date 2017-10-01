package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id       int
	Username string `orm:"size(50)"`
	Password string `orm:"size(50)"`
}

func CreateUser(user Users) string {
	fmt.Println(user)
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("users")
	fmt.Println(1)
	num, err := qs.Filter("username", user.Username).Count()
	message := ""

	// 如果没有重复用户
	if err == nil {
		if num == 0 {
			id, err := o.Insert(&user)
			if err == nil {
				fmt.Println("注册成功", id)
				message = "注册成功"
				return message
			}
		}

		fmt.Println("已被注册")
		message = "已被注册"
		return message
	}

	return "出现错误"
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
