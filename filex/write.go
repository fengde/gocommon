package filex

import (
	"encoding/csv"
	"io"
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

	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	modePerm := os.FileMode(0666)
	if len(mode) > 0 {
		modePerm = mode[0]
	}

	return os.WriteFile(path, content, modePerm)
}

// WriteAppend 追加写文件
func WriteAppend(path string, reader io.Reader) error {
	return Write(path, reader, os.ModeAppend)
}

// WriteCsvFile 写数据到csv
func WriteCsvFile(filepath string, records [][]string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE

	if append {
		flag = flag | os.O_APPEND
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Comma = ','

	return writer.WriteAll(records)
}

// WriteStringToFile 写字符串到文件
func WriteStringToFile(filepath string, content string, append bool) error {
	var flag int
	if append {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

// WriteBytesToFile 写bytes到文件
func WriteBytesToFile(filepath string, content []byte) error {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)
	return err
}
