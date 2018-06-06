package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
	"time"
	"class-admin/lib"
	"class-admin/utils"
)

type VideoController struct {
	rbac.CommonController
}

//日期类型转换string->Time，若空值默认最近30天
func checkDate(startString string,endString string) (start,end string) {
	if startString==""{
		startString=lib.FormatDate(time.Now().AddDate(0,0,-60))
	}
	startDate:=lib.TimeDate(startString)
	if endString==""{
		endString=lib.FormatDate(time.Now())
	}
	endDate:=lib.TimeDate(endString)
	endDate=endDate.AddDate(0,0,1)
	return lib.FormatTime(startDate),lib.FormatTime(endDate)
}

func (c *VideoController) GetVideoList()  {
	if !c.IsAjax() {
		c.TplName = "video/video.tpl"
		return
	}
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}
	startString:=c.GetString("startDate")
	endString:=c.GetString("endDate")
	start,end:=checkDate(startString,endString)
	name:=c.GetString("name")
	list,count:=m.GetVideoList(pager,start,end,name)
	c.Data["json"] = &map[string]interface{}{"itemsCount": count, "data": &list}
	c.ServeJSON()
	return
}

func (c *VideoController) SaveVideo() {
	u := m.Video{}
	if err := c.ParseForm(&u); err != nil {
		c.Rsp(false, err.Error())
		return
	}
	_,ih,_:=c.GetFile("poster")
	if ih!=nil{
		u.Poster=utils.UploadFile(ih)
	}
	_,vh,_:=c.GetFile("url")
	if vh!=nil{
		u.Url=utils.UploadFile(vh)
		if ih==nil{
			u.Poster=u.Url+"?x-oss-process=video/snapshot,t_1,f_jpg,h_0,w_0"
		}
	}

	id, err := u.Save()
	if err == nil && id > 0 {
		c.Rsp(true, "Success")
		return
	} else {
		c.Rsp(false, err.Error())
		return
	}
}

func (c *VideoController) DelVideo() {
	u := m.Video{}
	Id, _ := c.GetInt64("id")
	if Id == 0 {
		c.Rsp(false, "缺少id")
		return
	}
	u.Id = Id
	status, err := u.Delete()
	if err == nil && status > 0 {
		c.Rsp(true, "Success")
		return
	} else {
		c.Rsp(false, err.Error())
		return
	}
}

