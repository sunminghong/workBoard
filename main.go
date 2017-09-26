package main

import (
	_ "workBoard/models"
	_ "workBoard/routers"

	"github.com/astaxie/beego"
)

func init() {
	// beego bug 必须 gob.Register 才能注册 session 相关 struct
	// gob.Register(models.Users{})
}

func main() {
	beego.SetStaticPath("/static", "frontend/build/static/")
	beego.SetViewsPath("frontend/build")
	// re, _ := regexp.Compile("-")
	// uuid := re.ReplaceAllString(uuid.NewV4().String(), "")
	// fmt.Println(uuid)
	beego.Run()
	// 创建
	// u1 := uuid.NewV4()
	// fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	// u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")

	// if err != nil {
	// 	fmt.Printf("Something gone wrong: %s", err, 1)
	// 	return
	// }
}
