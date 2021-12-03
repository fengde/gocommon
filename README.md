## gocommon
gocmmon 对第三方包统一进行了封装，为工程提供通用的服务

### 推荐封装的意义


    *.统一不同工程使用体验最佳的三方库，防止相同功能的三方库在工程内部泛滥
    *.基础调用，方便从顶层视角做全局拦截，也可随时替换第三方方案
    *.基于第三方做二次封装，屏蔽复杂度，更便捷的使用
    *.基础代码的复用，是团队的一种沉淀，方便后续新人快速实现业务功能

### 包规范：

    为了避免和系统包名冲突，包名统一加x 作为suffix，x表示增强的意思

### 目前提供的子包如下（子包内容待不断完善）：

    errorx: 提供error创建，带stack的返回，方便调试
    hashx: 提供hash函数，如md5生成，hash数字生成
    jsonx: 提供json解析和反解析
    logx: 提供log打印方案
    mathx: 提供常用的math函数
    safex: 提供异常捕获的go方法，方便启动安全的协程
    storex: 提供存储相关的方案，如mysql,clickhouse,mongodb,redis等
    structx: 提供常用数据结构方案
    sysx: 提供通用系统参数查询的方案，如hostname, ip，mac等
    taskx: 提供并发任务执行，任务池
    timeoutx: 提供超时处理等函数


