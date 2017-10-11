package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户信息
type User struct {
	Id       int64
	Username string `orm:"size(50)"`
	Nickname string `orm:"size(50)"`
	Password string `orm:"size(50)"`
}

// 用户成功信息 json
type UserSuccessJson struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

// 错误信息 json
type ErrorMessage struct {
	Message string `json:"message"`
}

// 随机昵称
type Nickname struct {
	Adjective []string `json:"adjective"`
	Noum      []string `json:"noum"`
}

func CreateUser(user User) (UserSuccessJson, ErrorMessage) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("user")
	num, err := qs.Filter("username", user.Username).Count()

	var errorJson ErrorMessage
	var userJson UserSuccessJson

	if err == nil {
		// 如果没有重复用户
		if num == 0 {
			user.Nickname = GetMe()
			id, err := o.Insert(&user)

			fmt.Println(err)
			// 注册成功
			if err == nil {
				userJson.Username = user.Username
				userJson.Nickname = user.Nickname
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

func Login(user User) (UserSuccessJson, ErrorMessage) {
	o := orm.NewOrm()
	o.Using("default")

	// 存放数据库取出的 user
	var ouser User
	var errorJson ErrorMessage
	var userJson UserSuccessJson

	err := o.QueryTable("user").Filter("username", user.Username).One(&ouser)
	fmt.Println(err)

	if err == nil {
		// 如果密码不相等
		if ouser.Password != user.Password {
			errorJson.Message = "密码错误"
			return userJson, errorJson
		}

		userJson.Username = ouser.Username
		userJson.Nickname = ouser.Nickname
		userJson.Id = ouser.Id
		return userJson, errorJson
	}

	// 如果不存在用户，或者出现错误
	errorJson.Message = "用户不存在"
	return userJson, errorJson
}

func GetMe() string {
	// 读取 json
	bytes, err := ioutil.ReadFile("models/nickname.json")

	if err != nil {
		// 读取失败，生成 json，使用默认名称
		return "默认的昵称"
	}

	var nickname Nickname

	json.Unmarshal(bytes, &nickname)

	// 随机取出字符串
	var adjectiveLen int = len(nickname.Adjective) - 1
	var noumLen int = len(nickname.Noum) - 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var adjective string = nickname.Adjective[r.Intn(adjectiveLen)]
	var noum string = nickname.Noum[r.Intn(noumLen)]

	return fmt.Sprintf("%s的%s", adjective, noum)
}
