package exceptions

import (
	"fmt"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"log"
)

func detectError(err interface{}) string {
	switch e := err.(type) {
	case error:
		return e.Error()
	case string:
		return e
	default:
		return fmt.Sprintf("%v", err)
	}
}

func genErrorLog(msg error, err ...interface{}) {
	id := helpers.RandomString(8)
	logContent := fmt.Sprintf("Message: %s - Error: %s\n", msg, err)

	helpers.WriteLog(constants.ErrorLogPath, fmt.Sprintf("[Id: %s] - Date: ", id), logContent)
}

func genSystemLog(msg string) {
	id := helpers.RandomString(8)
	logContent := fmt.Sprintf("Message: %s\n", msg)

	helpers.WriteLog(constants.SystemLogPath, fmt.Sprintf("[Id: %s] - Date: ", id), logContent)
}

func Panic(msg error, err ...interface{}) {
	if err != nil {
		er := detectError(err)
		go genErrorLog(msg, er)
		panic(fmt.Sprintf("%s: %s\n", msg, er))
	}
}

func Fatal(msg error, err ...interface{}) {
	if err != nil {
		er := detectError(err)
		go genErrorLog(msg, er)
		log.Fatalf("%s: %s\n", msg, er)
	}
}

func Log(msg error, err ...interface{}) {
	if err != nil {
		er := detectError(err)
		go genErrorLog(msg, er)
		log.Printf("%s: %s\n", msg, er)
	}
}

func SystemLog(msg string) {
	go genSystemLog(msg)
	log.Printf("%s\n", msg)
}
