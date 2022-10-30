## gocommon

    gocmmon 对第三方包统一进行了封装，为工程提供便捷的功能入口

### 包规范：

    为了避免和系统包名冲突，包名统一加x 作为suffix，x表示增强的意思


### gocommon/timex： 时间格式转换
    使用： import "github.com/fengde/gocommon/timex"
    
    unix时间戳转字符串时间： s := timex.Unix2String(1667143429)
    unix时间戳转时间对象： t := timex.Unix2Time(1667143429)
    
    字符串时间转unix时间戳：unixSecond := timex.String2Unix("2022-10-30 00:00:00")
    字符串时间转时间对象： t, err := timex.String2Time("2022-10-30 00:00:00")
    字符串日期转时间对象： t, err := timex.DateString2Time("2022-10-30")
    
    时间对象转unix时间戳：unixSecond := timex.Time2Unix(t)
    时间对象转字符串时间： s := timex.Time2String(t)

    查询当前字符串时间(如2022-10-30 00:00:00) s := timex.NowTimeString()
    查询当前字符串日期(如2022-10-30) s := timex.NowDateString()
    查询当前unix时间戳：unixsecond := timex.NowUnix()
    查询当前unix毫秒时间戳：unixMill := timex.NowUnixMilli()
    查询当前unix微妙时间戳：unixMicro := timex.NowUnixMicro()
    查询当前unix纳秒时间戳：unixNano := timex.NowUnixNano()
    查询当前时间对象：t := timex.Now()
    
    使用时间宏定义(2006-01-02 15:04:05)：timex.DATETIME_LAYOUT
    使用日期宏定义(2006-01-02)：timex.DATE_LAYOUT


### gocommon/jsonx：jsonx序列化便捷函数
    使用： import "github.com/fengde/gocommon/jsonx"

    对象序列化成字符串： s := jsonx.MarshalToStringNoErr(obj)
    字符串反序列化成对象：err := jsonx.UnmarshalString(s, &obj)


### gocommon/flagx：命令行参数读取
    使用： import "github.com/fengde/gocommon/flagx"

    var input struct {
		Age   int     `flag:"age" default:"1" help:"年龄"`
		User  string  `flag:"user" default:"fengde"  help:"用户名称"`
		Money float64 `flag:"money" help:"金钱"`
		IsBoy bool  `flag:"old" help:"是男孩子吗"`
	}

	err := flagx.Parse(&input)

    struct成员变量类型目前支持int、int64、string、float64、bool类型，其他类型暂时不支持


### gocommon/safex：安全执行协程、函数、捕获异常等
    使用：import "github.com/fengde/gocommon/safex"
    带异常捕获的协程执行： safex.Go(fn)
    带异常捕获的函数执行： safex.Func(fn)


### gocommon/slicex：数组相关操作
    使用：import "github.com/fengde/gocommon/slicex"  
    字符串数组是否包含某元素： if slicex.StrContains(arr, "good") {}
    int数组是否包含某元素：if slicex.IntContains(arr, intNum) {}
    int64数组是否包含某元素：if slicex.Int64Contains(arr, int64Num) {}

    字符串数组移除重复的元素：arr = slicex.StrRemoveRepeat(arr)
    int数组移除重复的元素：arr = slicex.IntRemoveRepeat(arr)
    int64数组移除重复的元素：arr = slicex.Int64RemoveRepeat(arr)


### gocommon/httpx：http相关操作
    使用：import "github.com/fengde/gocommon/httpx"
    Get请求：httpx.Get(url， headers, urlValues)
    Post请求（json数据）：httpx.PostJSON(url， headers, obj)
    Put请求（json数据）：httpx.PutJSON(url， headers, obj)
    Delete请求：httpx.DeleteJSON(url， headers, obj)

    注意：
        headers即请求头键值对
        obj可以传string，[]byte，struct，map等格式，最终都会转换成json字符串

### gocommon/confx：工程配置，从yaml文件读取配置
    使用：查看gocommon/confx/example

    
    


