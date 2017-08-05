package route

import (
	"github.com/martini-contrib/render"
	"videoGo/controllers"
	"videoGo/config"
)
func RouteInit()  {
	m := config.MartiniMap()
	m.Use(render.Renderer(render.Options{
		Directory: "resources/views",
		Extensions: []string{".tmpl", ".html"},
	}))
	m.Get("/", controllers.VideoIndex)
	m.Get("/category/:id", controllers.VideoCategory)
	m.Get("/info/:id", controllers.VideoInfo)
	m.RunOnAddr(":8080")
}
