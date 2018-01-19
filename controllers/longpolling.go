package controllers

import (
	"order/models"
	"fmt"
)

type LongPollingController struct {
	baseController
}

func (this *LongPollingController) Join()  {
	uanme:=this.GetString("uname")
	if len(uanme)==0{
		this.Redirect("/",302)
		return
	}
   Join(uanme,nil)
fmt.Println("***************")
   this.TplName="longpolling.html"
   this.Data["IsLongPolling"]=true
   this.Data["UserName"]=uanme
}
func (this *LongPollingController) Post()  {
	this.TplName = "longpolling.html"
	uname:=this.GetString("uname")
	content := this.GetString("content")
	if len(uname) == 0||len(content)==0{
		return
	}
	publish <- newEvent(models.EVENT_MESSAGE,uname,content)
}
func (this *LongPollingController)Fetch()  {
	lastReceived,err:=this.GetInt("lastReceived")
	if err!=nil{
		return
	}
	events:=models.GetEvents(int(lastReceived))
	if len(events)>0{
		this.Data["json"]=events
		this.ServeJSON()
		return
	}
	//等待一个新信息
	ch:=make(chan bool)
	waitingList.PushBack(ch)
	<-ch
	this.Data["json"]= models.GetEvents(int(lastReceived))
 this.ServeJSON()
	}