package controllers

import (
	"class-admin/controllers/rbac"
	m "class-admin/models"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type GradeController struct {
	rbac.CommonController
}

func (c *GradeController) GetClassList() {
	if !c.IsAjax(){
		c.TplName = "grade/grade.tpl"
		return
	}
	stuNo:=c.GetString("stuNo")
	lectureId,_:=c.GetInt("lectureId")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	sortField := c.GetString("sortField")
	sortOrder := c.GetString("sortOrder")
	pager:=m.Pager{pageIndex,pageSize,sortField,sortOrder}

	userinfo := c.GetSession("userinfo")
	user:=userinfo.(m.User)

	class := m.GetClassList(pager,stuNo,user.Username,lectureId)
	c.Data["json"] = &map[string]interface{}{"itemsCount": len(class), "data": &class}
	c.ServeJSON()

}

func (c *GradeController) Import()  {
	lectureId,err:=c.GetInt("lectureId",0)
	if lectureId==0||err!=nil{
		c.Rsp(false, "lectureId数据有误")
		return
	}
	file,_,err:=c.GetFile("grade")
	if file==nil{
		c.Rsp(false, err.Error())
		return
	}
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		c.Rsp(false, err.Error())
		return
	}
	rows := xlsx.GetRows("Sheet1")
	no:=-1
	grade:=-1
	if len(rows)<1||len(rows[0])<2{
		c.Rsp(false, "Excel文件格式有误")
		return
	}
	var class []*m.Class
	for i, row := range rows {
		if i==0 {
			for j, colCell := range row {
				if colCell=="学生学号"{
					no=j
				}
				if colCell=="成绩"{
					grade=j
				}
			}
			if no==-1||grade==-1{
				c.Rsp(false, "Excel缺少学号或成绩列")
				return
			}
			continue
		}
		score,err:=strconv.ParseFloat(row[grade],64)
		if err!=nil{
			c.Rsp(false, "成绩数据格式有误")
			return
		}
		class=append(class,&m.Class{StuNo:row[no],Score:score,LectureId:int64(lectureId)})
	}
	err=m.UpdateScore(class)
	if err!=nil{
		c.Rsp(false, "数据库错误")
	}else{
		c.Rsp(true, "Success")
	}
}