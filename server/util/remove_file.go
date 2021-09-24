package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

func RemoveFile(filePath string) {
	if filePath == "" {
		return
	}
	err := os.Remove(filePath)
	if err != nil {
		logrus.Error(err)
		return
	}
}
