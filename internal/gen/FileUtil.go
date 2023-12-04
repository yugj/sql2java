package gen

import (
	"log"
	"os"
)

func createDirIfNeed(basePath string, name string) {

	var path = basePath + "/" + name
	autoCreateDir(path)
}

func autoCreateDir(path string) {

	if !existsPath(path) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
}

func existsPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func openFile(destFile string) *os.File {
	file, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Open file err =", err)
		panic(err)
	}
	//defer file.Close()
	return file
}
