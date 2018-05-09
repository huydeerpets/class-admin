package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Material struct {
	Id         int64      `json:"id"           form:"id"`
	LectureId  string     `form:"lectureId"    json:"lecture_id"`
	Name       string     `json:"name"         form:"name"`
	Type       string     `json:"type"         form:"type"`
	Url        string     `json:"url"`
	CreatedAt  time.Time  `json:"created_at" orm:"type(datetime);auto_now_add"`
	UpdatedAt  time.Time  `json:"updated_at" orm:"type(datetime);auto_now"`
	Extension  string     `json:"extension"`
	LoadCount  int        `json:"load_count" orm:"default:0"`
	FileName   string     `json:"file_name"`
}

const MaterialTable = "material"
func (n *Material) TableName() string {
	return MaterialTable
}

func init() {
	orm.RegisterModel(new(Material))
}

func GetMaterialList(pager Pager,start,end,name,lessonNo,lessonName string) ([]orm.Params) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select n.*,lec.lesson_no as les_no,les.name as les_name,lec.class_time as class_time from "+MaterialTable+" as n "
	sql+=" left join "+LectureTable+" as lec on n.lecture_id=lec.id "+
		" left join "+LessonTable+" as les on les.number=lec.lesson_no where 1=1 "
	if lessonNo!=""{
		sql+=" and lec.lesson_no = '"+lessonNo+"'"
	}
	if lessonName!=""{
		sql+=" and les.name like '%"+lessonName+"%' "
	}
	if start!=""{
		sql+=" and updated_at >= '"+start+"'"
	}
	if end!=""{
		sql+=" and updated_at <= '"+end+"'"
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

func (t *Material) Save() (id int64, err error) {
	o := orm.NewOrm()
	if t.Id > 0{
		temp:=Material{Id:t.Id}
		if o.Read(&temp) == nil {
			temp.Name=t.Name
			temp.LectureId=t.LectureId
			temp.Type=t.Type
			if t.Url!=""{
				temp.Url=t.Url
				temp.Extension=t.Extension
				temp.FileName=t.FileName
				temp.LoadCount=0
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

func (t *Material) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Material{Id: t.Id})
	return status, err
}


