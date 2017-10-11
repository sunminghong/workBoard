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

		this.SetSession("user", user)
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

		var user models.User
		user.Id = userJson.Id
		user.Nickname = userJson.Nickname
		user.Username = userJson.Username
		this.SetSession("user", user)
	} else {
		this.Data["json"] = errorJson
		this.Ctx.Output.SetStatus(403)
	}

	this.ServeJSON()
}

func (this *UserMeController) Get() {
	userSession := this.GetSession("user")

	if userSession == nil {
		var errorMessage models.ErrorMessage
		errorMessage.Message = "用户未登录"
		this.Ctx.Output.SetStatus(404)
	} else {
		this.Data["json"] = userSession
	}

	this.ServeJSON()
}
