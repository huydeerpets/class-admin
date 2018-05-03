package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Video struct {
	Id         int64      `json:"id"           form:"id"`
	Name       string     `json:"name"         form:"name"`
	Type       string     `json:"type"         form:"type"`
	LessonNo   string     `json:"lesson_no"    form:"lesson_no"`
	Url        string     `json:"url"`
	Poster     string     `json:"poster"`
	Brief      string     `json:"brief"        form:"brief"`
	ViewNum    int        `json:"view_num"`
	CreatedAt  time.Time  `json:"created_at" orm:"auto_now_add"`
	Uploader   string     `json:"uploader"`
}

const VideoTable = "video"
func (n *Video) TableName() string {
	return VideoTable
}

func init() {
	orm.RegisterModel(new(Video))
}

func GetVideoList(pager Pager,start,end,name string) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select * from "+VideoTable+" where 1=1 "
	if start!=""{
		sql+=" and created_at >= '"+start+"'"
	}
	if end!=""{
		sql+=" and created_at <= '"+end+"'"
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

func (t *Video) Save() (id int64, err error) {
	o := orm.NewOrm()
	if t.Id > 0{
		_ , err = o.Update(t)
		if err == nil{
			id=t.Id
		}
	}else{
		t.CreatedAt=time.Now()
		id, err = o.Insert(t)
	}
	return
}

func (t *Video) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Lecture{Id: t.Id})
	return status, err
}


