package controllers

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"videoGo/config"
	"videoGo/structs"
	"log"
)
var Data struct{
	Items []map[string]interface{}
	Categorys []map[string]interface{}
	Info map[string]structs.Videos
}
//首页
func VideoIndex(r render.Render) {
	items, _ := config.XormMap().SqlMapClient("videoLimit10").Query().List()
	Data.Items = items
	Category,_ := config.XormMap().SqlMapClient("videoTypeAll").Query().List()
	Data.Categorys = Category
	r.HTML(200, "homeIndex", Data)
}
//分类
func VideoCategory(params martini.Params,r render.Render, req *http.Request){
	page,limit := 1, 10
	pageRow := (page -1) * page
	sql := map[string]interface{}{"typeId": params["id"], "start": pageRow, "limit": limit}
	res, _ := config.XormMap().SqlMapClient("videoAllByTypeId", &sql).Query().List()
	Data.Items = res
	r.HTML(200, config.PajxTemplate(req, "homeCategory", "homePajxCategory"), Data);
}
//详情页
func VideoInfo(params martini.Params, r render.Render, req *http.Request)  {
	info := make(map[string]structs.Videos, 0)
	err := config.XormMap().Where("id = ?", params["id"]).Find(&info)
	if(err != nil){
		log.Fatal(err)
	}
	Data.Info = info
	r.HTML(200, config.PajxTemplate(req, "homeInfo", "homePajxInfo"), info);
}