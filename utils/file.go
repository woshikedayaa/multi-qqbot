package utils

import (
	"io/ioutil"
	"multi-qqbot/log"
	"os"
)

func ReadDir(path string) []os.FileInfo {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error(err)
	}
	return infos
}
