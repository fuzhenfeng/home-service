package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"home-service/models/user"
	"log"
	"strconv"
)

type UserController struct {
	beego.Controller
}

// Get 根据用户ID查询用户信息
func (c *UserController) Get() {
	userId := c.Ctx.Input.Param(":userId")
	id, _ := strconv.Atoi(userId)
	user := user.Get(id)
	c.Data["user"] = user
	c.ServeJSON()
}

// Put 用户修改信息
func (c *UserController) Put() {
	body := c.Ctx.Input.RequestBody
	temp := user.User{}
	json.Unmarshal(body, &temp)
	user.Set(temp)
	c.ServeJSON()
}

func (c *UserController) Login() {
	body := c.Ctx.Input.RequestBody
	temp := user.User{}
	json.Unmarshal(body, &temp)
	log.Println(temp)
	c.Data["json"] = map[string]interface{}{
		"success": true,
	}
	c.ServeJSON()
}
