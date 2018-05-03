package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)

type Class struct {
	Id         int64   `form:"id"         json:"id"`
	LectureId  int64   `form:"lectureId"  json:"lecture_id"`
	StuNo      string  `form:"stuNo"      json:"stu_no"`
	Score      float64 `form:"score"      json:"score"`
}

const ClassTable = "class"

func (Class) TableName() string {
	return ClassTable
}

func init()  {
	orm.RegisterModel(new(Class))
}

func GetClassList(pager Pager,stuNo string,lecture int) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select c.*,les.name as les_name,s.name as stu_name,t.name as tea_name,lec.term as term from "+ClassTable+" as c "
	sql+=" left join "+StudentTable+" as s on s.stu_id=c.stu_no "+
		" left join "+LectureTable+" as lec on c.lecture_id=lec.id "+
			" left join "+LessonTable+" as les on les.number= lec.lesson_no "+
				" left join "+TeacherTable+" as t on t.number=lec.teacher_no where 1=1 "
	if stuNo!=""{
		sql+=" and stu_no = '"+stuNo+"'"
	}
	if lecture!=0{
		sql+=" and lecture_id = "+strconv.Itoa(lecture)+""
	}
	var sort string
	if len(pager.SortField) > 0 {
		sort = pager.SortField + " "+pager.SortOrder
	} else {
		sort = "id"
	}
	var offset int
	if pager.PageIndex <= 1 {
		offset = 0
	} else {
		offset = (pager.PageIndex - 1) * pager.PageSize
	}
	_,err:=o.Raw(sql+" order by "+sort+" limit ? offset ?",pager.PageSize,offset).Values(&maps)
	if err!=nil{
		fmt.Println(err)
		return []orm.Params{}
	}
	return maps
}

func (t *Class) Save() (id int64, err error) {
	o := orm.NewOrm()
	if t.Id > 0{
		_ , err = o.Update(t)
		if err == nil{
			id=t.Id
		}
	}else{
		id, err = o.Insert(t)
	}
	return
}

func (t *Class) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Class{Id: t.Id})
	return status, err
}


