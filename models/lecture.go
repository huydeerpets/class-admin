package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Lecture struct {
	Id         int64   `form:"id"         json:"id"`
	TeacherNo  string  `form:"teacherNo"  json:"teacher_no"`
	LessonNo   string  `form:"lessonNo"   json:"lesson_no"`
	Place      string  `form:"place"      json:"place"`
	Term       string  `form:"term"       json:"term"`
	ClassTime  string  `form:"classTime"  json:"class_time"`
	People     int     `form:"people"     json:"people"`
}

const LectureTable = "lecture"

func (Lecture) TableName() string {
	return LectureTable
}

func init()  {
	orm.RegisterModel(new(Lecture))
}

func GetLectureList(pager Pager,teacherNo,lessonNo,term string) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select lec.*,les.name as les_name,t.name as tea_name from "+LectureTable+" as lec "
	sql+=" left join "+LessonTable+" as les on les.number=lec.lesson_no "+
		" left join "+TeacherTable+" as t on t.number=lec.teacher_no where 1=1 "
	if teacherNo!=""{
		sql+=" and teacher_no = '"+teacherNo+"'"
	}
	if lessonNo!=""{
		sql+=" and lesson_no = '"+lessonNo+"'"
	}
	if term!=""{
		sql+=" and term = '"+term+"'"
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

func GetLecByTea(teacherNo string) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select lec.*,les.name as les_name from "+LectureTable+" as lec "
	sql+=" left join "+LessonTable+" as les on les.number=lec.lesson_no "+
		" where teacher_no = '"+teacherNo+"'"
	_,err:=o.Raw(sql).Values(&maps)
	if err!=nil{
		fmt.Println(err)
		return []orm.Params{}
	}
	return maps
}

func (t *Lecture) Save() (id int64, err error) {
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

func (t *Lecture) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Lecture{Id: t.Id})
	return status, err
}


