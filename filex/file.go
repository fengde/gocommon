package filex

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Write 写文件到本地
func Write(path string, reader io.Reader, mode ...os.FileMode) error {
	dir := filepath.Dir(path)

	if !IsDirExist(dir) {
		if err := CreateDir(dir); err != nil {
			return err
		}
	}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	modePerm := os.FileMode(0666)
	if len(mode) > 0 {
		modePerm = mode[0]
	}

	return ioutil.WriteFile(path, content, modePerm)
}

// IsDirExist 文件夹是否存在
func IsDirExist(dir string) bool {
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateDir 创建文件夹
func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// RemoveDir 删除文件夹
func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}