package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"fmt"
)

var langTypes []string

func init() {
	langTypes = strings.Split(beego.AppConfig.String("lang_types"),"|")
		for _,lang:=range langTypes{
			beego.Trace("Loading language: "+lang)
			fmt.Println(lang)
			if err:=i18n.SetMessage(lang,"conf/"+"locale_"+lang+".ini");err !=nil{
				beego.Error("Fail to set message file:",err)
				return
			}
		}

}

type BaseController struct {
	beego.Controller
	i18n.Locale
}

func (this *BaseController)Prepare()  {
	this.Lang = ""
	al:=this.Ctx.Request.Header.Get("Accept-Language")//获取浏览器的请求头来判断语言
	fmt.Println("al=", al)
	if len(al)>4{
		al=al[:5]
		if i18n.IsExist(al){
			this.Lang=al
		}
	}
	if len(this.Lang)==0{
		this.Lang="en-US"
	}
	fmt.Println("lang=", this.Lang)
	this.Data["Lang"] = this.Lang
}

type AppController struct {
	BaseController
}

func (this *AppController) Get()  {
	this.TplName="welcome.html"
}
func (this *AppController) Join()  {
	uname:=this.GetString("uname")
	tech:=this.GetString("tech")
	if len(uname)==0{
		this.Redirect("/",302)
		return
	}
	switch tech {
	case "longpolling":
		this.Redirect("/lp?uname="+uname,302)
	default:
		this.Redirect("/",302)

	}
	return //通常从定向之后都需要返回
}