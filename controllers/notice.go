package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
)

type NoticeController struct {
	rbac.CommonController
}

func (c *NoticeController) GetNoticeList()  {
	if !c.IsAjax() {
		c.TplName = "notice/notice.tpl"
		return
	}
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}
	lessonNo:=c.GetString("lessonNo")
	lessonName:=c.GetString("lessonName")
	list:=m.GetNoticeList(pager,lessonNo,lessonName)
	c.Data["json"] = &map[string]interface{}{"itemsCount": len(list), "data": &list}
	c.ServeJSON()
	return
}

func (c *NoticeController) Save() {
	u := m.Notice{}
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

func (c *NoticeController) Del() {
	u := m.Notice{}
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

