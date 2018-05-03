package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
)

type GradeController struct {
	rbac.CommonController
}

func (c *GradeController) GetClassList() {
	stuNo:=c.GetString("stuNo")
	lectureId,_:=c.GetInt("lectureId")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}

	class := m.GetClassList(pager,stuNo,lectureId)
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"itemsCount": len(class), "data": &class}
		c.ServeJSON()
		return
	} else {
		c.TplName = "grade/class.tpl"
	}
}