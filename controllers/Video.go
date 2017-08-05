package controllers

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"videoGo/config"
	"videoGo/models"
)

func VideoIndex(r render.Render) {
	var Data struct{
		Result []map[string]interface{}
		Category []map[string]interface{}
	}
	Data.Result = models.GetVideoIndexDefault()
	Data.Category = models.GetVideoTypeList()
	r.HTML(200, "homeIndex", Data)
}

func VideoCategory(params martini.Params,r render.Render, req *http.Request){
	var Data struct{
		Item []map[string]interface{}
	}
	Data.Item = models.GetVideoCategoryData(params["id"], 1)
	r.HTML(200, config.PajxTemplate(req, "homeCategory", "homePajxCategory"), Data);
}

func VideoInfo(params martini.Params, r render.Render, req *http.Request)  {
	Data := models.GetVideoInfo(params["id"])
	r.HTML(200, config.PajxTemplate(req, "homeInfo", "homePajxInfo"), Data);
}