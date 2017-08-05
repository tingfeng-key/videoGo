package config

import (
	"github.com/go-martini/martini"
	"github.com/xormplus/xorm"
	"github.com/xormplus/core"
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
	//SnakeMapper 支持struct为驼峰式命名映射
	x.SetMapper(core.SnakeMapper{})
	//采用sql/xormcfg.ini配置文件中的配置，可直接使用InitSqlMap()初始化
	x.InitSqlMap()
	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	//该监控模式下，如删除配置文件，内存中不会删除相关配置
	x.StartFSWatcher()
	x.ShowSQL(true)
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