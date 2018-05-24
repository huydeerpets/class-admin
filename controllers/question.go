package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
)

type QuesController struct {
	rbac.CommonController
}

func (c *QuesController) GetQuestionList()  {
	if !c.IsAjax() {
		c.TplName = "question/question.tpl"
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
	lessonNo:=c.GetString("lessonNo")
	lessonName:=c.GetString("lessonName")
	stuName:=c.GetString("stuName")

	userinfo := c.GetSession("userinfo")
	user:=userinfo.(m.User)

	list:=m.GetQuestionList(pager,start,end,lessonNo,lessonName,stuName,user.Username)
	c.Data["json"] = &map[string]interface{}{"itemsCount": len(list), "data": &list}
	c.ServeJSON()
	return
}

func (c *QuesController) Save() {
	u := m.Question{}
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

func (c *QuesController) AnsIndex()  {
	quesId,_:=c.GetInt("id")
	ques:=m.GetQuestionById(quesId)
	c.Data["ques"]=ques
	c.TplName = "question/answer.tpl"
}






