package filex

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

// ReadFileToString 读文件到字符串
func ReadFileToString(path string) (string, error) {
	_bytes, err := ReadFileToBytes(path)
	if err != nil {
		return "", err
	}
	return string(_bytes), nil
}

// ReadFileToBytes 读文件到bytes
func ReadFileToBytes(path string) ([]byte, error) {
	_bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return _bytes, nil
}

// ReadFileToLines 读文件到行
func ReadFileToLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, l)
	}

	return result, nil
}

// ReadCsvFile 读csv文件
func ReadCsvFile(filepath string) ([][]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
