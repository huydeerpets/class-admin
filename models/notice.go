package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Notice struct {
	Id         int64   `form:"id"         json:"id"`
	Title      string  `form:"teacherNo"  json:"teacher_no"`
	LectureId  string  `form:"lessonNo"   json:"lesson_no"`
	Content    string  `form:"place"      json:"place"`
	CreatedAt  string  `json:"created_at" orm:"auto_now_add"`
}

const NoticeTable = "notice"

func (Notice) TableName() string {
	return NoticeTable
}

func init()  {
	orm.RegisterModel(new(Notice))
}


func GetNoticeList(pager Pager,lessonNo string) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select n.*,les.name as les_name,lec.class_time as class_time from "+NoticeTable+" as n "
	sql+=" left join "+LectureTable+" as lec on n.lecture_id=lec.id "+
		" left join "+LessonTable+" as les on les.number=lec.lesson_no where 1=1 "
	if lessonNo!=""{
		sql+=" and lec.lesson_no = '"+lessonNo+"'"
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

func (t *Notice) Save() (id int64, err error) {
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

func (t *Notice) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Notice{Id: t.Id})
	return status, err
}


