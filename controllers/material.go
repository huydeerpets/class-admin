package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
	path2 "path"
	"class-admin/utils"
)

type MatController struct {
	rbac.CommonController
}

func (c *MatController) GetMaterialList()  {
	if !c.IsAjax() {
		c.TplName = "material/material.tpl"
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
	lessonNo:=c.GetString("lessonNo")
	lessonName:=c.GetString("lessonName")

	userinfo := c.GetSession("userinfo")
	user:=userinfo.(m.User)

	list,count:=m.GetMaterialList(pager,start,end,name,lessonNo,lessonName,user.Username)
	c.Data["json"] = &map[string]interface{}{"itemsCount": count, "data": &list}
	c.ServeJSON()
	return
}

func (c *MatController) Save() {
	u := m.Material{}
	if err := c.ParseForm(&u); err != nil {
		c.Rsp(false, err.Error())
		return
	}
	_,vh,_:=c.GetFile("url")
	if vh!=nil{
		vSuffix:= path2.Ext(vh.Filename)
		u.Url=utils.UploadFile(vh)
		u.Extension=vSuffix
		u.FileName=vh.Filename
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

func (c *MatController) Del() {
	u := m.Material{}
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


