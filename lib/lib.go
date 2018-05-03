package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//password hash function
func Pwdhash(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}


// TimeParse ..
func TimeParse(t string) time.Time {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return tm
}

// TimeDate ..
func TimeDate(t string) time.Time {
	tm, _ := time.ParseInLocation("2006-01-02", t, time.Local)
	return tm
}

// FormatTime ..
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 03:04:05")
}

// FormatDate ..
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}