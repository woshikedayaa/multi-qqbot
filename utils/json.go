package utils

import (
	"encoding/hex"
	"io"
	"io/ioutil"
	"multi-qqbot/log"
)

func ReadJsonFromReaderByString(r io.Reader) string {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		log.Error(err)
	}
	return hex.EncodeToString(bs)
}

func ReadJsonFromReaderByBytes(r io.Reader) []byte {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		log.Error(err)
	}
	return bs
}

func ReadJsonFromFileDir(path string) (ss [][]byte) {

	for _, v := range ReadDir(path) {
		bs, err := ioutil.ReadFile(path + "/" + v.Name())
		if err != nil {
			log.Error(err)
		}
		ss = append(ss, bs)
	}
	return
}
