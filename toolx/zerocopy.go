package toolx

import "unsafe"

// String2bytes 零拷贝实现string转bytes
func String2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Bytes2string 零拷贝实现bytes转string
func Bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
