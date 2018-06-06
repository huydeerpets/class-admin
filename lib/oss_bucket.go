package lib

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/revel/revel"
	"github.com/astaxie/beego"
)

var Bucket *oss.Bucket

func InitOssBucket() {
	oss_endpoint := beego.AppConfig.String("oss_endpoint")
	oss_access_id := beego.AppConfig.String("oss_access_id")
	oss_access_pass := beego.AppConfig.String("oss_access_pass")
	oss_bucket := beego.AppConfig.String("oss_bucket")
	client, err := oss.New(oss_endpoint, oss_access_id, oss_access_pass)
	if err!=nil{
		revel.WARN.Printf("oss错误: %v", err)
	}
	Bucket, err = client.Bucket(oss_bucket)
	if err!=nil{
		revel.WARN.Printf("oss错误: %v", err)
	}
}

func GetCDNUrl(filename string) string {
	return beego.AppConfig.String("oss_website")+"/"+filename
}
