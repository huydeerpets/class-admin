package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
)

type WorkController struct {
	rbac.CommonController
}

func (c *WorkController) GetWorkList()  {
	if !c.IsAjax() {
		c.TplName = "work/work.tpl"
		return
	}
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}
	lessonNo:=c.GetString("lessonNo")
	lessonName:=c.GetString("lessonName")

	userinfo := c.GetSession("userinfo")
	user:=userinfo.(m.User)

	list,count:=m.GetWorkList(pager,lessonNo,lessonName,user.Username)
	c.Data["json"] = &map[string]interface{}{"itemsCount": count, "data": &list}
	c.ServeJSON()
	return
}

func (c *WorkController) Save() {
	u := m.Work{}
	if err := c.ParseForm(&u); err != nil {
		c.Rsp(false, err.Error())
		return
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

func (c *WorkController) Del() {
	u := m.Work{}
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

