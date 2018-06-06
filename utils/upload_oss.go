package utils

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"time"
	"bytes"
	"github.com/revel/revel"
	"mime/multipart"
	"log"
	"path"
	"class-admin/lib"
)

func UploadNews(url string) (code string,err error) {
	code=""
	doc,err:=goquery.NewDocument(url)
	if err!=nil{
		revel.WARN.Println("newDocument Error: ", err)
		return
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		var uploadChan = make(chan string)
		src,exists:=s.Attr("data-src")
		if !exists{
			return
		}
		dataType,exists:=s.Attr("data-type")
		if !exists{
			revel.WARN.Println("get img data-type Error: ", err)
			return
		}
		go uploadImg(src,dataType,uploadChan)
		imageUrl:=<-uploadChan
		if imageUrl!=""{
			s.SetAttr("src",imageUrl)
		}
	})
	newsHtml,_:=doc.Html()
	filename := "article_html_"+strconv.FormatInt(time.Now().Unix(), 10)+".html"
	err=lib.Bucket.PutObject(filename,bytes.NewReader([]byte(newsHtml)))
	if err!=nil{
		revel.WARN.Println("PutObject Error: ", err)
		return
	}
	code=lib.GetCDNUrl(filename)
	return
}

func uploadImg(src string,dataType string,uploadChan chan string) {
	resp,err:=http.Get(src)
	if err!=nil{
		uploadChan<-""
		revel.WARN.Println("get img src Error: ", err)
		return
	}
	filenames := []string{
		"article_img",
		strconv.FormatInt(time.Now().Unix(), 10),
		RandStringNum(2),
	}
	filename := strings.Join(filenames, "_") + "." +dataType
	err=lib.Bucket.PutObject(filename,resp.Body)
	if err!=nil{
		uploadChan<-""
		return
	}
	uploadChan<-lib.GetCDNUrl(filename)
}

func UploadFile(fileHeader *multipart.FileHeader) string {
	file,_:=fileHeader.Open()
	fileSuffix:= path.Ext(fileHeader.Filename)
	filename:="file_"+strconv.FormatInt(time.Now().Unix(), 10)+fileSuffix
	err := lib.Bucket.PutObject(filename, file)
	if err != nil {
		log.Println("UploadFileToOss error", err)
	}
	return lib.GetCDNUrl(filename)
}

func UploadHtml(html string) string {
	filename:="spider_html_"+strconv.FormatInt(time.Now().Unix(), 10)+".html"
	err := lib.Bucket.PutObject(filename, bytes.NewReader([]byte(html)))
	if err != nil {
		log.Println("UploadHtmlToOss error", err)
	}
	return lib.GetCDNUrl(filename)
}