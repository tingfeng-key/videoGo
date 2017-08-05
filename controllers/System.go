package controllers

import (
	"videoGo/config"
	_ "github.com/go-sql-driver/mysql"
	"videoGo/structs"
)

func SystemInit() string {
	enign := config.XormMap()
	enign.CreateTables(&structs.Videos{}, &structs.Types{})
	return "初始化数据库完成"
}
