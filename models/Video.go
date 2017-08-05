package models

import (
	"videoGo/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type Video struct {
	Id int64
	Title string
	Thumb string
}
func GetVideoIndexDefault() []map[string]interface{} {
	sql := "select * from videos limit 10"
	res, _ := config.XormMap().Sql(sql).Query().List()
	return res
}

func GetVideoCategoryData(id string, page int) []map[string]interface{} {
	limit := 10
	pageRow := (page -1) * page
	sql := "SELECT * FROM videos WHERE `type_id`="+id+" LIMIT "+strconv.Itoa(pageRow)+","+strconv.Itoa(limit)
	res, _ := config.XormMap().Sql(sql).Query().List()
	return res
}

func GetVideoInfo(id string) Video {
	video := Video{}
	err := config.XormMap().Where("id = ?", id).Find(&video)
	if(err != nil){
		log.Fatal(err)
	}
	return video
}
