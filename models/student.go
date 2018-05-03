package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"errors"
)

type Student struct {
	Id     int64   `form:"id" json:"id"`
	StuId  string  `orm:"size(100)" form:"stuId"  valid:"Required" json:"stu_id"`
	Name   string  `orm:"size(100)" form:"name"  valid:"Required" json:"name"`
	Gender int     `orm:"default(1)" form:"gender"  valid:"Required" json:"gender"`
	Email  string  `orm:"size(255)" form:"email" json:"email"`
	Tel    string  `orm:"size(100)" form:"tel" json:"tel"`

}
const StudentTable="student"
func (n *Student) TableName() string {
	return StudentTable
}

func init() {
	orm.RegisterModel(new(Student))
}

func GetStudentList(pager Pager,stuId,name string) (students []Student, count int64) {
	students=[]Student{}
	o := orm.NewOrm()
	var r orm.RawSeter
	sql:="select * from "+StudentTable+" where 1=1 "
	if stuId!=""{
		sql+=" and stu_id = '"+stuId+"'"
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
	count,err:=r.QueryRows(&students)
	if err!=nil{
		fmt.Println(err)
		return []Student{},0
	}
	return students, count
}

func (t *Student) Save() (id int64, err error) {
	o := orm.NewOrm()
	var students []Student
	if t.Id > 0{
		num,err:=o.Raw("select id from "+StudentTable+" where id <> ? and stu_id = ?",t.Id,t.StuId).QueryRows(&students)
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
		num,err:=o.Raw("select id from "+StudentTable+" where stu_id = ?",t.StuId).QueryRows(&students)
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

func (t *Student) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Student{Id: t.Id})
	return status, err
}

