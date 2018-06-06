package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)

type Question struct {
	Id         int64      `json:"id"           form:"id"`
	LectureId  string     `form:"lectureId"`
	Title      string     `json:"title"`
	Question   string     `json:"question"`
	Number     int        `json:"number"` //第几节课
	StuNo      int        `json:"stu_no"`
	Answer     string     `json:"answer"       form:"answer"`
	CreatedAt  time.Time  `json:"created_at"   orm:"type(datetime);auto_now_add"`
	UpdatedAt  time.Time  `json:"updated_at"   orm:"type(datetime);auto_now"`
	Status     int        `json:"status"`
}


const QuestionTable = "question"
func (n *Question) TableName() string {
	return QuestionTable
}

func init() {
	orm.RegisterModel(new(Question))
}

func GetQuestionList(pager Pager,start,end,lessonNo,lessonName,stuName,teaNo string) ([]orm.Params,int) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select n.id,n.title,n.question,n.number,n.updated_at"+
		",stu.name as stu_name,lec.lesson_no as les_no,les.name as les_name,lec.class_time as class_time from "+QuestionTable+" as n "
	sql+=" left join "+LectureTable+" as lec on n.lecture_id=lec.id "+
		" left join "+LessonTable+" as les on les.number=lec.lesson_no "+
			"left join "+StudentTable+" as stu on n.stu_no=stu.stu_id where lec.teacher_no='"+teaNo+"' "
	if lessonNo!=""{
		sql+=" and lec.lesson_no = '"+lessonNo+"'"
	}
	if lessonName!=""{
		sql+=" and les.name like '%"+lessonName+"%' "
	}
	if stuName!=""{
		sql+=" and stu.name like '%"+stuName+"%' "
	}
	if start!=""{
		sql+=" and updated_at >= '"+start+"'"
	}
	if end!=""{
		sql+=" and updated_at <= '"+end+"'"
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
		return []orm.Params{},0
	}
	if len(maps)==0{
		return []orm.Params{},0
	}
	var allMaps []orm.Params
	o.Raw(sql).Values(&allMaps)
	return maps,len(allMaps)
}

func (t *Question) Save() (id int64, err error) {
	o := orm.NewOrm()
	if t.Id > 0{
		temp:=Question{Id:t.Id}
		if o.Read(&temp) == nil {
			temp.Answer=t.Answer
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

func (t *Question) Delete() (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Question{Id: t.Id})
	return status, err
}

func GetQuestionById(id int) orm.Params {
	var maps []orm.Params
	o := orm.NewOrm()
	sql:="select n.*,stu.name as stu_name,lec.lesson_no as les_no,les.name as les_name,lec.class_time as class_time from "+QuestionTable+" as n "
	sql+=" left join "+LectureTable+" as lec on n.lecture_id=lec.id "+
		" left join "+LessonTable+" as les on les.number=lec.lesson_no "+
		"left join "+StudentTable+" as stu on n.stu_no=stu.stu_id where n.id="+strconv.Itoa(id)+" limit 1 "
	_,err:=o.Raw(sql).Values(&maps)
	if err!=nil || len(maps)==0 {
		fmt.Println(err)
		return orm.Params{}
	}
	return maps[0]
}

