package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// models.GetWeather("北京")
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.MainController{})
}
