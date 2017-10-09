package controllers

import (
	"encoding/json"
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

type UserMeController struct {
	beego.Controller
}

func (this *UserCreateController) Post() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	userJson, errorJson := models.CreateUser(user)

	// 注册成功
	if errorJson.Message == "" {
		this.Data["json"] = userJson
	} else {
		this.Data["json"] = errorJson
		this.Ctx.Output.SetStatus(404)
	}

	this.ServeJSON()
}

func (this *UserLogInController) Post() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	userJson, errorJson := models.Login(user)

	// 如果登录成功
	if errorJson.Message == "" {
		this.Data["json"] = userJson
	} else {
		this.Data["json"] = errorJson
		this.Ctx.Output.SetStatus(403)
	}

	this.ServeJSON()
}

func (this *UserMeController) Get() {
	userSession := this.GetSession("user")

	if userSession == nil {
		// 如果没有 session，随机生成一个昵称
		userJson := models.GetMe()
		this.Data["json"] = userJson

		// 把生成的昵称，写入 session
		var user models.User
		user.Nickname = userJson.Nickname
		this.SetSession("user", user)
	} else {
		this.Data["json"] = userSession
	}

	this.ServeJSON()
}
