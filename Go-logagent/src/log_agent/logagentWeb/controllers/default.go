package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//c.Layout = "layout/layout.html"
	//c.TplName = "app/apply.html"
	c.TplName = "ljw.html"
}
