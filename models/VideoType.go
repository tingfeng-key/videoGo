package models

import (
	"videoGo/config"
)

type VideoTypes struct {
	Id int64
	Name string
	Sort int64
	Status int64
}

func GetVideoTypeList() []map[string]interface{}  {
	sql := "SELECT * FROM video_types"
	res,_ := config.XormMap().Sql(sql).Query().List()
	return res
}