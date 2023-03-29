package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"home-service/models/album"
	"strconv"
)

type AlbumControllers struct {
	beego.Controller
}

// GetByPageNo 分页查询
func (c *AlbumControllers) GetByPageNo() {
	pageNo, _ := strconv.Atoi(c.Ctx.Input.Param(":pageNo"))
	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":userId"))
	album := album.GetByPageNo(userId, pageNo)
	c.Data["album"] = album
	c.ServeJSON()
}
