package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Marshal 对象json化成字节序列
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalNoErr 对象json化成字节序列返回，无错误
func MarshalNoErr(v interface{}) []byte {
	bs, _ := Marshal(v)
	return bs
}

// MarshalToStringNoErr 对象json化成字符串返回，无错误
func MarshalToStringNoErr(v interface{}) string {
	bs, _ := Marshal(v)
	return string(bs)
}

// Unmarshal 字节序列反序列化成对象
func Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(string(data), err)
	}

	return nil
}

// UnmarshalString 字符串反序列化成对象
func UnmarshalString(str string, v interface{}) error {
	decoder := json.NewDecoder(strings.NewReader(str))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(str, err)
	}

	return nil
}

// UnmarshalReader reader反序列化成对象
func UnmarshalReader(reader io.Reader, v interface{}) error {
	var buf strings.Builder
	teeReader := io.TeeReader(reader, &buf)
	decoder := json.NewDecoder(teeReader)
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(buf.String(), err)
	}

	return nil
}

func unmarshalUseNumber(decoder *json.Decoder, v interface{}) error {
	decoder.UseNumber()
	return decoder.Decode(v)
}

func formatError(v string, err error) error {
	return fmt.Errorf("string: `%s`, error: `%s`", v, err.Error())
}
