#### api 练习项目一阶段

使用Gin框架

* 效果

  1. 访问对应的路由可以查看系统cpu,disk,ram占用情况
  2. 通过读取配置文件来配置运行情况
  3. 增加日志输出和保存到文件,并设置日志单个大小,备份文件个数(日志配置使用第三方日志包)
  4. 链接mysql数据库,从配置文件读取数据库相关链接信息
  5. 自定义业务错误信息
  6. 读取和返回HTTP请求
* 包依赖:

```ssh
git clone https://github.com/lexkong/vendor
```

* CURL测试命令

  ```shell
  -X/--request [GET|POST|PUT|DELETE|…]  指定请求的 HTTP 方法
  -H/--header                           指定请求的 HTTP Header
  -d/--data                             指定请求的 HTTP 消息体（Body）
  -v/--verbose                          输出详细的返回信息
  -u/--user                             指定账号、密码
  -b/--cookie                           读取 cookie...
  ```

* 本项目测试语句

  ```shell
  $ curl -XGET http://127.0.0.1:8080/sd/health

  $ curl -XGET http://127.0.0.1:8080/sd/disk

  $ curl -XGET http://127.0.0.1:8080/sd/cpu

  $ curl -XGET http://127.0.0.1:8080/sd/ram
  ```

* go语法记录

  ```shell
  message : =fmt.Sprintf("%s",text)
  格式化字符串,不输出
  ```

* 用到的库文件

  > viper

```shell
Viper 配置读取顺序：

viper.Set() 所设置的值
命令行 flag
环境变量
配置文件
配置中心：etcd/consul
默认值
```

> Go ORM


```go
import (
  "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/mysql"
)
db, err := gorm.Open("mysql", config) //链接数据库
defer db.Close() //关闭数据库链接,Close()数据库对象调用即可
.....
....

//设置数据库链接相关
func setupDB(db *gorm.DB){
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(2000) //用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) //用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

```

> Gin

```go
// Param()：返回 URL 的参数值，例如
 router.GET("/user/:id", func(c *gin.Context) {
  // a GET request to /user/john
     id := c.Param("id") // id == "john"
 })

//Query()：读取 URL 中的地址参数，例如
// GET /path?id=1234&name=Manu&value=
   c.Query("id") == "1234"
   c.Query("name") == "Manu"
   c.Query("value") == ""
   c.Query("wtf") == ""  

//DefaultQuery()：类似 Query()，但是如果 key 不存在，会返回默认值，例如
 //GET /?name=Manu&lastname=
 c.DefaultQuery("name", "unknown") == "Manu"
 c.DefaultQuery("id", "none") == "none"
 c.DefaultQuery("lastname", "none") == ""

//Bind()：检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。apiserver 采用的媒体类型是 JSON，所以 Bind() 是按 JSON 格式解析的。

//GetHeader()：获取 HTTP 头。
```



* 日志说明见配置文件