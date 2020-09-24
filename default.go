package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "2757634617@qq.com"
	c.TplName = "index.tpl"
	name:=c.Ctx.Input.Query("name")
	age:=c.Ctx.Input.Query("age")
	sex:=c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	if name !="liuhong" ||age != "20" {
		c.Ctx.ResponseWriter.Write([]byte("数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据正确"))
}
func (c *MainController) Post() {
	fmt.Println("Post类型的请求")

	user:=c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为:",user)
	password:=c.Ctx.Request.FormValue("password")
	fmt.Println("用户名为:",password)

	if user !="admir"|| password!="123456"{
		c.Ctx.ResponseWriter.Write([]byte("对不起数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据正确"))
}