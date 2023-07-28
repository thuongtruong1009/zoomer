package exceptions

import (
	"fmt"
	"log"
)

func detectError(err interface{}) string {
	switch e := err.(type) {
	case error:
		return e.Error()
	case string:
		return e
	default:
		return err.(string)
	}
}

func Panic(err interface{}, msg ...string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s\n", msg, detectError(err)))
	}
}

func Fatal(err error, msg ...string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, detectError(err))
	}
}

func Log(err error, msg ...string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, detectError(err))
	}
}
