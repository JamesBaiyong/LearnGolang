###### api 练习项目一阶段

使用Gin框架

* 效果

  访问对应的路由可以查看系统cpu,disk,ram占用情况


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

  ​