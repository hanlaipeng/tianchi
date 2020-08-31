package util

import (
	"os"
)

type JsonOptions struct {
	InputPath string
	OutPath   string
	Data      []byte
}


func CheckFileIsExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func (j *JsonOptions)WriteJson() error {

	fp, err := os.OpenFile(j.OutPath, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer fp.Close()
	_, err = fp.Write(j.Data)
	if err != nil {
		return err
	}

	return nil
}