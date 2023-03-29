package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"home-service/models/album"
	"strconv"
)

type AlbumController struct {
	beego.Controller
}

// Get 根据ID查询图片
func (c *AlbumController) Get() {
	pictureId := c.Ctx.Input.Param(":pictureId")
	id, _ := strconv.Atoi(pictureId)
	album := album.Get(id)
	c.Data["album"] = album
	c.ServeJSON()
}

// Post 创建图片
func (c *AlbumController) Post() {
	body := c.Ctx.Input.RequestBody
	temp := album.Album{}
	json.Unmarshal(body, &temp)
	album.Set(temp)
	c.ServeJSON()
}

// Put 修改图片
func (c *AlbumController) Put() {
	body := c.Ctx.Input.RequestBody
	temp := album.Album{}
	json.Unmarshal(body, &temp)
	album.Set(temp)
	c.ServeJSON()
}
