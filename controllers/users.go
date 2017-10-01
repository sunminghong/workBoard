package controllers

import (
	"encoding/json"
	"fmt"
	"workBoard/models"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type UserCreateController struct {
	beego.Controller
}

type UserLogInController struct {
	beego.Controller
}

func (this *UserCreateController) Post() {
	var user models.Users
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	this.ServeJSON()
	models.CreateUser(user)
}

func (this *UserLogInController) Post() {
	var ob models.Users
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	mes := models.Login(ob)
	u := this.GetSession("user")
	fmt.Println(u)
	if mes == "登录成功" {
		this.SetSession("user", ob)
	}

	this.ServeJSON()
}
