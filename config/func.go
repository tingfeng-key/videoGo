package config

import (
	"github.com/go-martini/martini"
	"github.com/xormplus/xorm"
	"log"
	"net/http"
)
/**
返回martini对象
 */
func MartiniMap() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(martini.Static("resources/public"))
	return m
}
/**
返回xorm对象
 */
func XormMap()  *xorm.Engine{
	x, err := xorm.NewEngine(Db().Type, Db().Dns)
	if(err != nil){
		log.Fatalln(err)
	}
	return x
}
/**
是否为ajax请求
 */
func Isajax(req *http.Request) bool {
	return req.Header.Get("X-Requested-With") == "XMLHttpRequest"
}
/**
判断http和ajax
返回模板名
 */
func PajxTemplate(req *http.Request, httpName string, pajxName string) string {
	if(Isajax(req)){
		return pajxName
	}
	return httpName;
}