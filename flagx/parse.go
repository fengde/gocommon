package flagx

import (
	"flag"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// 解析并获取命令行入参，传入结构体指针
// 结构体格式示例：
//
//		var input struct {
//			Age   int     `flag:"age" default:"1" help:"input your age"`
//			User  string  `flag:"user" default:"fedel" help:"input your user name"`
//			Money float64 `flag:"money" help:"input your money"`
//			Old   bool    `flag:"old" help:"are you old man"`
//	}
//
// 命令行入参帮助：go run main.go --help
// 命令行入参格式：go run main.go --age=32 --user=fedel --money=1000000 --old=true
func Parse(v any) error {
	ti := reflect.TypeOf(v)
	vi := reflect.ValueOf(v)
	realti := ti.Elem()
	realvi := vi.Elem()

	if ti.Kind() != reflect.Ptr || realti.Kind() != reflect.Struct {
		return errors.New("入参需要结构体指针")
	}

	t := map[string]interface{}{}

	for i := 0; i < realti.NumField(); i++ {
		filedti := realti.Field(i)
		filedFlag, ok := filedti.Tag.Lookup("flag")
		if !ok {
			return errors.New(filedti.Name + "缺失必要的tag: flag")
		}
		filedDefault := filedti.Tag.Get("default")
		filedHelp := filedti.Tag.Get("help")

		var v interface{}
		switch filedti.Type.Kind() {
		case reflect.String:
			v = flag.String(filedFlag, filedDefault, filedHelp)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var int64Default int64 = 0
			var err error
			if filedDefault != "" {
				int64Default, err = strconv.ParseInt(filedDefault, 10, 64)
				if err != nil {
					return errors.New(filedti.Name + "default默认值类型错误")
				}
			}
			v = flag.Int64(filedFlag, int64Default, filedHelp)
		case reflect.Float32, reflect.Float64:
			var float64Default float64 = 0.0
			var err error
			if filedDefault != "" {
				float64Default, err = strconv.ParseFloat(filedDefault, 64)
				if err != nil {
					return errors.New(filedti.Name + "default默认值类型错误")
				}
			}
			v = flag.Float64(filedFlag, float64Default, filedHelp)
		case reflect.Bool:
			var boolDefault bool
			var err error
			if filedDefault != "" {
				boolDefault, err = strconv.ParseBool(filedDefault)
				if err != nil {
					return errors.New(filedti.Name + "default默认值类型错误")
				}
			}

			v = flag.Bool(filedFlag, boolDefault, filedHelp)
		default:
			return errors.New(filedti.Type.Kind().String() + "，不支持类型不支持")
		}
		t[filedFlag] = v
	}

	flag.Parse()

	for i := 0; i < realti.NumField(); i++ {
		filedti := realti.Field(i)
		filedFlag := filedti.Tag.Get("flag")
		switch filedti.Type.Kind() {
		case reflect.String:
			v := t[filedFlag].(*string)
			realvi.Field(i).Set(reflect.ValueOf(*v))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := t[filedFlag].(*int64)
			if filedti.Type.Kind() == reflect.Int {
				realvi.Field(i).Set(reflect.ValueOf(int(*v)))
			} else if filedti.Type.Kind() == reflect.Int8 {
				realvi.Field(i).Set(reflect.ValueOf(int8(*v)))
			} else if filedti.Type.Kind() == reflect.Int16 {
				realvi.Field(i).Set(reflect.ValueOf(int16(*v)))
			} else if filedti.Type.Kind() == reflect.Int32 {
				realvi.Field(i).Set(reflect.ValueOf(int32(*v)))
			} else if filedti.Type.Kind() == reflect.Int64 {
				realvi.Field(i).Set(reflect.ValueOf(*v))
			}
		case reflect.Float32, reflect.Float64:
			v := t[filedFlag].(*float64)
			if filedti.Type.Kind() == reflect.Float32 {
				realvi.Field(i).Set(reflect.ValueOf(float32(*v)))
			} else if filedti.Type.Kind() == reflect.Float64 {
				realvi.Field(i).Set(reflect.ValueOf(*v))
			}
		case reflect.Bool:
			v := t[filedFlag].(*bool)
			realvi.Field(i).Set(reflect.ValueOf(*v))
		}
	}

	return nil
}
