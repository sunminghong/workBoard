package routers

import (
	"workBoard/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// models.GetWeather("北京")
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/users/create", &controllers.UserCreateController{})
	beego.Router("/api/users/login", &controllers.UserLogInController{})
}
