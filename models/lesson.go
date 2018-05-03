package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"errors"
)

type Lesson struct {
	Id       int64   `form:"id"      json:"id"`
	Number   string  `form:"number"  json:"number"`
	Name     string  `form:"name"    json:"name"`
	Credit   int     `form:"credit"  json:"credit"`
	Type     string  `form:"type"    json:"type"`
}

const LessonTable = "lesson"

func (Lesson) TableName() string {
	return LessonTable
}

func init()  {
	orm.RegisterModel(new(Lesson))
}


func GetLessonList(pager Pager,number,name string) (lessons []Lesson, count int64) {
	lessons=[]Lesson{}
	o := orm.NewOrm()
	var r orm.RawSeter
	sql:="select * from "+LessonTable+" where 1=1 "
	if number!=""{
		sql+=" and number = '"+number+"'"
	}
	if name!=""{
		sql+=" and name like '%"+name+"%'"
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
	r=o.Raw(sql+" order by "+sort+" limit ? offset ?",pager.PageSize,offset)
	count,err:=r.QueryRows(&lessons)
	if err!=nil{
		fmt.Println(err)
		return []Lesson{},0
	}
	return lessons, count
}

func (t *Lesson) Save() (id int64, err error) {
	o := orm.NewOrm()
	var lessons []Lesson
	if t.Id > 0{
		num,err:=o.Raw("select id from "+LessonTable+" where id <> ? and number = ?",t.Id,t.Number).QueryRows(&lessons)
		if err != nil{
			return id,err
		}
		if num>0{
			return id,errors.New("已存在相同编号的课程")
		}
		_ , err = o.Update(t)
		if err == nil{
			id=t.Id
		}
	}else{
		num,err:=o.Raw("select id from "+LessonTable+" where number = ?",t.Number).QueryRows(&lessons)
		if err!=nil{
			return id,err
		}
		if num>0{
			return id,errors.New("已存在相同编号的课程")
		}
		id, err = o.Insert(t)
	}
	return
}

func (t *Lesson) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Lesson{Id: t.Id})
	return status, err
}


