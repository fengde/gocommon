## gocommon

    gocmmon 对第三方包统一进行了封装，为工程提供便捷的功能入口

### 包规范：

    为了避免和系统包名冲突，包名统一加x 作为suffix，x表示增强的意思

---
#### timex： 时间格式转换
使用： 
```
    import "github.com/fengde/gocommon/timex"
```

unix时间戳转字符串时间：
```
    s := timex.Unix2String(1667143429)
```

unix时间戳转时间对象： 
```
    t := timex.Unix2Time(1667143429)
```
    
字符串时间转unix时间戳：
```    
    unixSecond := timex.String2Unix("2022-10-30 00:00:00")
```
字符串时间转时间对象：
```
    t, err := timex.String2Time("2022-10-30 00:00:00")
```

字符串日期转时间对象： 
```
    t, err := timex.DateString2Time("2022-10-30")
```
    
时间对象转unix时间戳：
```
    unixSecond := timex.Time2Unix(t)
```
    
时间对象转字符串时间： 
```
s := timex.Time2String(t)
```

查询当前字符串时间(如2022-10-30 00:00:00)：
```
    s := timex.NowTimeString()
```

查询当前字符串日期(如2022-10-30)：
```
    s := timex.NowDateString()
```

查询当前unix时间戳：
```
    unixsecond := timex.NowUnix()
```

查询当前unix毫秒时间戳：
```
    unixMill := timex.NowUnixMilli()
```

查询当前unix微妙时间戳：
```
    unixMicro := timex.NowUnixMicro()
```

查询当前unix纳秒时间戳：
```
    unixNano := timex.NowUnixNano()
```

查询当前时间对象：
```
    t := timex.Now()
```
    
使用时间宏定义(2006-01-02 15:04:05)：
```
    timex.DATETIME_LAYOUT
```

使用日期宏定义(2006-01-02)：
```
    timex.DATE_LAYOUT
```

---
#### jsonx：json序列化便捷函数
使用： 
```
    import "github.com/fengde/gocommon/jsonx"
```

对象序列化成字符串： 
```
    s := jsonx.MarshalToStringNoErr(obj)
```
字符串反序列化成对象：
```
    err := jsonx.UnmarshalString(s, &obj)
```

---
#### flagx：命令行参数读取
    
使用： 
```
    import "github.com/fengde/gocommon/flagx"
    
    var input struct {
        Age   int     `flag:"age" default:"1" help:"年龄"`
        User  string  `flag:"user" default:"fengde"  help:"用户名称"`
        Money float64 `flag:"money" help:"金钱"`
        IsBoy bool  `flag:"old" help:"是男孩子吗"`
    }
    
    err := flagx.Parse(&input)
```

struct成员变量类型目前支持int、int64、string、float64、bool类型，其他类型暂时不支持

---
#### safex：安全执行协程、函数、捕获异常等
使用：
```
    import "github.com/fengde/gocommon/safex"
```
    
带异常捕获的协程执行： 
```
    safex.Go(fn)
```

带异常捕获的函数执行： 
```
    safex.Func(fn)
```

---
#### slicex：数组相关操作
使用：
```
    import "github.com/fengde/gocommon/slicex"  
```

字符串数组是否包含某元素： 
```
    if slicex.StrContains(arr, "good") {}
```
int数组是否包含某元素：
```
    if slicex.IntContains(arr, intNum) {}
```

int64数组是否包含某元素：
```
    if slicex.Int64Contains(arr, int64Num) {}
```

字符串数组移除重复的元素：
```
    arr = slicex.StrRemoveRepeat(arr)
```

int数组移除重复的元素：
```
    arr = slicex.IntRemoveRepeat(arr)
```

int64数组移除重复的元素：
```
    arr = slicex.Int64RemoveRepeat(arr)
```
---

#### httpx：http相关操作

使用：
```
    import "github.com/fengde/gocommon/httpx"
```

get请求：
```
    httpx.Get(url， headers, urlValues)
```

post请求（json数据）：
```
    httpx.PostJSON(url， headers, obj)
```

put请求（json数据）：
```
    httpx.PutJSON(url， headers, obj)
```

delete请求：
```
    httpx.DeleteJSON(url， headers, obj)
```


headers即请求头键值对，obj可以传string，[]byte，struct，map等格式，最终都会转换成json字符串

---
#### confx：工程配置，从yaml文件读取配置
使用：
```
    查看gocommon/confx/example
```

---
#### taskx：任务执行相关
并发控制使用：
```
    import "github.com/fengde/gocommon/taskx"
    func main() {
        // 设置最大支持10个协程同时运行，不传入参数则不限制协程数
        g := taskx.NewTaskGroup(10)
        g.Run(func() {
            fmt.Println("hello world")
        })
        g.Wait()
    }
```
函数组串行执行：
```    
    import "github.com/fengde/gocommon/taskx"
    func main() {
        sg := taskx.NewSerialTaskGroup(func1, func2, func3)
        // 顺序执行函数，遇到返回err的函数停止执行
        if err := sg.Run(); err != nil {
            fmt.Println(err)
        }
    }
```

---
#### funcx: 函数执行相关
使用：
```
    import "github.com/fengde/gocommon/funcx"
```
函数重试：
```
    // 最多重试3次，成功退出，每次重试中间sleep 1秒
    funcx.Retry(3, time.Second, func(loop int) error {
        fmt.Println("hello world")
        return errors.New("test retry")
    })
```
函数重复执行：
```
    // 重复执行函数3次
    funcx.Repeat(3, func() {
        fmt.Println("hello world")
    })
```
函数上锁：
```
    locker := funcx.NewFuncLocker()

    index := 0
    for index < 100 {
        // Exec执行的函数，会竞争锁，得到锁才执行
        go locker.Exec(func() {
            fmt.Println("hello world")
        })
        index++
    }

```
---
#### base64x：base64编解码
使用
```
    import "github.com/fengde/gocommon/base64x"
```
Url base64编解码
```
    t := base64x.UrlEncode([]byte("your url"))
    t := base64x.UrlDecode("base64 string")
```
通用的base64编解码
```
    t := base64x.Encode([]byte("your url"))
    t := base64x.Decode("base64 string")
```

---
#### hashx: 哈希算法集合
使用
```
    import "github.com/fengde/gocommon/hashx"
```
获取hash值：
```
    number := hashx.Hash([]byte(""))
```
获取md5值：
```
    md5Bytes := hashx.Md5([]byte(""))
```
获取md5 16进制字符串：
```
    md5HexStr := hashx.Md5Hex([]byte(""))
```
获取Sha256值：
```
    sha256Bytes := hashx.Sha256([]byte(""))
```
获取Sha256 16进制字符串：
```
    sha256Str := hashx.Sha256Hex([]byte(""))
```

---
#### sendx/emailx：普通邮件发送
使用
```
    import "github.com/fengde/gocommon/sendx/emailx"
    client := emailx.NewEmailClient(host, port, user, password)
    client.SendText(...)
    client.SendHTML(...)
```

---
#### sysx: 系统相关
使用
```
    import "github.com/fengde/gocommon/sysx"
```
获取主机名：
```
    hostname := sysx.Hostname()
```


---
#### ratelimitx: 限流算法
使用
```
    import "github.com/fengde/gocommon/ratelimitx"
```
令牌桶算法限流：
```
    limiter := ratelimitx.NewTokenBucketRatelimit(...)
    limiter.Run(...)
    limiter.RunWithTimeout(...)
```
漏桶算法限流：
```
    limiter := NewLeakyBucketRatelimit(...)
    limiter.Run(...)
```