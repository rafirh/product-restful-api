package helper

import "time"

func StrToDateTime(strTime string) time.Time {
	result, err := time.Parse("2006-01-02 15:04:05", strTime)
	PanicIfError(err)
	return result
}