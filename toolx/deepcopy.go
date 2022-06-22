package toolx

import "github.com/mohae/deepcopy"

// DeepCopy 对象深拷贝, 对于结构体对象，未公开的变量不会copy
func DeepCopy(src interface{}) interface{} {
	return deepcopy.Copy(src)
}
