package file

import (
	"github.com/lizhennet/gorm-gen/pkg/log"
	"os"
)

func CreatIsNotExist(filepath string) error {
	exist, err := PathIsExist(filepath)
	if err != nil {
		return err
	}
	if !exist {
		log.Info("create dir:%s", filepath)
		return os.MkdirAll(filepath, os.ModePerm)
	}
	return nil
}

func PathIsExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
