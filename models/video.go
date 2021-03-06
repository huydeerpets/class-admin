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
	CreatedAt  time.Time  `json:"created_at" orm:"type(datetime);auto_now_add"`
	UpdatedAt  time.Time  `json:"updated_at" orm:"type(datetime);auto_now"`
	Uploader   string     `json:"uploader"`
}

const VideoTable = "video"
func (n *Video) TableName() string {
	return VideoTable
}

func init() {
	orm.RegisterModel(new(Video))
}

func GetVideoList(pager Pager,start,end,name string) ([]orm.Params,int) {
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
		sort = "created_at desc"
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
		return []orm.Params{},0
	}
	if len(maps)==0{
		return []orm.Params{},0
	}
	var allMaps []orm.Params
	o.Raw(sql).Values(&allMaps)
	return maps,len(allMaps)
}

func (t *Video) Save() (id int64, err error) {
	o := orm.NewOrm()
	if t.Id > 0{
		temp:=Video{Id:t.Id}
		if o.Read(&temp) == nil {
			temp.Name=t.Name
			temp.Brief=t.Brief
			temp.Type=t.Type
			temp.LessonNo=t.LessonNo
			if t.Url!=""{
				temp.Url=t.Url
			}
			if t.Poster!=""{
				temp.Poster=t.Poster
			}
			_ , err = o.Update(&temp)
			if err == nil{
				id=t.Id
			}
		}

	}else{
		id, err = o.Insert(t)
	}
	return
}

func (t *Video) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Video{Id: t.Id})
	return status, err
}


