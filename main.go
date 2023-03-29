package main

import (
	_ "home-service/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

