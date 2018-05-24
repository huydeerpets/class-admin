package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Work struct {
	Id         int64   `form:"id"         json:"id"`
	Title      string  `form:"title"      json:"title"`
	LectureId  string  `form:"lectureId"  json:"lecture_id"`
	Content    string  `form:"content"    json:"content"`
	CreatedAt  time.Time  `json:"created_at" orm:"auto_now_add"`
	UpdatedAt  time.Time  `json:"updated_at" orm:"type(datetime);auto_now"`
}

const WorkTable = "homework"

func (Work) TableName() string {
	return WorkTable
}

func init() {
	orm.RegisterModel(new(Work))
}

func GetWorkList(pager Pager,lessonNo,lessonName,teaNo string) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select n.*,lec.lesson_no as les_no,les.name as les_name,lec.class_time as class_time from "+WorkTable+" as n "
	sql+=" left join "+LectureTable+" as lec on n.lecture_id=lec.id "+
		" left join "+LessonTable+" as les on les.number=lec.lesson_no where lec.teacher_no='"+teaNo+"' "
	if lessonNo!=""{
		sql+=" and lec.lesson_no = '"+lessonNo+"'"
	}
	if lessonName!=""{
		sql+=" and les.name like '%"+lessonName+"%' "
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
	if len(maps)==0{
		return []orm.Params{}
	}
	return maps
}

func (t *Work) Save() (id int64, err error) {
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

func (t *Work) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Work{Id: t.Id})
	return status, err
}


