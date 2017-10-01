package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 选择数据库
	orm.RegisterDataBase("default", "mysql", "root:jinzhuming@/work_board?charset=utf8")

	//注册 model
	orm.RegisterModel(new(Board))
	orm.RegisterModel(new(Users))

	// // 建表
	// CreateTables()

	// test
	// var board Board

	// Created(board)
}
