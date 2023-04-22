package utils

import (
	"os"
)

func GetFilePath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd + path
}

// func FileExists(path string) bool {
// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		return false
// 	}
// 	return true
// }

// func FileDelete(path string) error {
// 	if err := os.Remove(path); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func FileCreate(path string) error {
// 	if _, err := os.Create(path); err != nil {
// 		return err
// 	}
// 	return nil
// }
