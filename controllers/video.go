package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
	"time"
	"class-admin/lib"
	path2 "path"
	"strconv"
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
	list:=m.GetVideoList(pager,start,end,name)
	c.Data["json"] = &map[string]interface{}{"itemsCount": len(list), "data": &list}
	c.ServeJSON()
	return
}

func (c *VideoController) SaveVideo() {
	u := m.Video{}
	if err := c.ParseForm(&u); err != nil {
		c.Rsp(false, err.Error())
		return
	}
	_,vh,_:=c.GetFile("url")
	_,ih,_:=c.GetFile("poster")
	vSuffix:= path2.Ext(vh.Filename)
	iSuffix:= path2.Ext(ih.Filename)
	unix:=strconv.FormatInt(time.Now().Unix(), 10)
	vfilename:="video_"+unix+vSuffix
	ifilename:="image_"+unix+iSuffix
    vpath:=path2.Join("file","video",vfilename)
    ipath:=path2.Join("file","video",ifilename)
	c.SaveToFile("url",vpath)
	c.SaveToFile("poster",ipath)
	u.Url="/"+vpath
	u.Poster="/"+ipath
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
	Id, _ := c.GetInt64("Id")
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

