package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
)

type InfoController struct {
	rbac.CommonController
}


func (c *InfoController) GetTeaList() {

	number:=c.GetString("number")
	name:=c.GetString("name")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}
	teachers, count := m.GetTeacherList(pager,number,name)
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"itemsCount": count, "data": &teachers}
		c.ServeJSON()
		return
	} else {
		c.TplName = "info/teacher.tpl"
	}
}

func (c *InfoController) SaveTea() {
	u := m.Teacher{}
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

func (c *InfoController) DelTea() {
	u := m.Teacher{}
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


func (c *InfoController) GetStuList() {

	stuId:=c.GetString("stuId")
	name:=c.GetString("name")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}

	students, count := m.GetStudentList(pager,stuId,name)
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"itemsCount": count, "data": &students}
		c.ServeJSON()
		return
	} else {
		c.TplName = "info/student.tpl"
	}
}


func (c *InfoController) SaveStu() {
	u := m.Student{}
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

func (c *InfoController) DelStu() {
	u := m.Student{}
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


func (c *InfoController) GetLesList() {

	number:=c.GetString("number")
	name:=c.GetString("name")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}

	lessons, count := m.GetLessonList(pager,number,name)
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"itemsCount": count, "data": &lessons}
		c.ServeJSON()
		return
	} else {
		c.TplName = "info/lesson.tpl"
	}
}


func (c *InfoController) SaveLes() {
	u := m.Lesson{}
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

func (c *InfoController) DelLes() {
	u := m.Lesson{}
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

func (c *InfoController) GetLecList() {
	teacherNo:=c.GetString("teacherNo")
	lessonNo:=c.GetString("lessonNo")
	term:=c.GetString("term")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}

	lessons := m.GetLectureList(pager,teacherNo,lessonNo,term)
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"itemsCount": len(lessons), "data": &lessons}
		c.ServeJSON()
		return
	} else {
		c.TplName = "info/lecture.tpl"
	}
}

func (c *InfoController) GetLecByTea()  {
	userinfo := c.GetSession("userinfo")
	user:=userinfo.(m.User)
	lessons := m.GetLecByTea(user.Username)
	c.Data["json"]=lessons
	c.ServeJSON()
}


func (c *InfoController) SaveLec() {
	u := m.Lecture{}
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

func (c *InfoController) DelLec() {
	u := m.Lecture{}
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

func (c *InfoController) GetClassList() {
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
		c.TplName = "info/class.tpl"
	}
}


func (c *InfoController) SaveClass() {
	u := m.Class{}
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

func (c *InfoController) DelClass() {
	u := m.Class{}
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