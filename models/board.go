package models

import (
	"fmt"
	"regexp"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

type Board struct {
	Token    string `orm:"PK"`
	Describe string `orm:"size(500)"`
	Users    string
	Tasks    string
	Deadline string `orm:"size(100)"`
}

func CreatedUuid() string {
	re, _ := regexp.Compile("-")
	uuid := re.ReplaceAllString(uuid.NewV4().String(), "")
	return uuid
}

func CreatedBoard(board Board) {
	o := orm.NewOrm()
	o.Using("default")

	board.Token = CreatedUuid()

	token, err := o.Insert(&board)
	fmt.Println(board)

	if err == nil {
		fmt.Println("注册成功", token)
	}

	// qs := o.QueryTable("boards")
}
