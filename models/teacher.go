package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
	"fmt"
)

type Teacher struct {
	Id     int64   `form:"id" json:"id"`
	Number string  `orm:"size(100)" form:"number"  valid:"Required" json:"number"`
	Name   string  `orm:"size(100)" form:"name"  valid:"Required" json:"name"`
	Gender int     `orm:"default(1)" form:"gender"  valid:"Required" json:"gender"`
	Office string  `orm:"null;size(255)" form:"office" valid:"MaxSize(200)" json:"office"`
	Email  string  `orm:"size(255)" form:"email" json:"email"`
	Tel    string  `orm:"size(100)" form:"tel" json:"tel"`

}
const TeacherTable="teacher"
func (n *Teacher) TableName() string {
	return TeacherTable
}

func init() {
	orm.RegisterModel(new(Teacher))
}

func GetTeacherList(pager Pager,number,name string) (teachers []Teacher, count int64) {
	teachers=[]Teacher{}
	o := orm.NewOrm()
	var r orm.RawSeter
	sql:="select * from "+TeacherTable+" where 1=1 "
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
	count,err:=r.QueryRows(&teachers)
	if err!=nil{
		fmt.Println(err)
		return []Teacher{},0
	}
	return teachers, count
}

func (t *Teacher) Save() (id int64, err error) {
	o := orm.NewOrm()
	var teachers []Teacher
	if t.Id > 0{
        num,err:=o.Raw("select id from "+TeacherTable+" where id <> ? and number = ?",t.Id,t.Number).QueryRows(&teachers)
        if err != nil{
        	return id,err
		}
		if num>0{
			return id,errors.New("已存在相同编号的教师")
		}
        _ , err = o.Update(t)
        if err == nil{
        	id=t.Id
        }
	}else{
		num,err:=o.Raw("select id from "+TeacherTable+" where number = ?",t.Number).QueryRows(&teachers)
		if err!=nil{
			return id,err
		}
		if num>0{
			return id,errors.New("已存在相同编号的教师")
		}
		id, err = o.Insert(t)
	}
	return
}

func (t *Teacher) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Teacher{Id: t.Id})
	return status, err
}

