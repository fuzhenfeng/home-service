package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"home-service/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/:userId([\\w]+)", &controllers.UserController{}, "get:Get")
	beego.Router("/:userId([\\w]+)", &controllers.UserController{}, "put:Put")
	beego.Router("/:userId([\\w]+)", &controllers.UserController{}, "post:Login")

	beego.Router("/:userId([\\w]+)/album/:pictureId([\\w]+)", &controllers.AlbumController{}, "get:Get")
	beego.Router("/:userId([\\w]+)/album/:pictureId([\\w]+)", &controllers.AlbumController{}, "put:Put")
	beego.Router("/:userId([\\w]+)/album/:pictureId([\\w]+)", &controllers.AlbumController{}, "post:Post")

	beego.Router("/:userId([\\w]+)/albums/:pageNo([\\w]+)", &controllers.AlbumControllers{}, "get:GetByPageNo")
}
