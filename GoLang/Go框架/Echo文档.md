# [Echo 中文文档](go-echo.org)

## Guide

### 安装
`go get -u github.com/labstack/echo/...`

### Context
1. echo.Context表示当前HTTP请求的上下文。通过路径、路径参数、数据、注册处理器和相关API进行请求的读取与响应的输出。
2. 扩展Context
    * 定义一个自定义的context
    
    ```golang
    type CustomContext struct {
        echo.Context
    }
    func (c *CustomContext) Foo() {
        println("foo")
    }
    func (c *CustomContext) Bar() {
        println("bar")
    }
    ``` 
    
    * 创建一个中间件来扩展默认的context
        
    ```go
    e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
        return func (c echo.Context) error {
            cc:=&CustomContext{c}
            return h(cc)
        }
    })
    ```
    
    * 在处理器中使用
        ```
        e.Get("/",func(c echo.Context) error){
            cc:=c.(*CustomContext)
            cc.Foo()
            cc.Bar() 
        }
        ```

### Cookie
1. 设置Cookie     
    
    ```
    func writeCookie(c echo.Context) error{
        cookie:=new(http.Cookie)
        cookie.Name = "username"
        cookie.Value = "lucy"
        cookie.Expires = time.Now().Add(24 * time.Hour)
        c.SetCookie(cookie)
        c.String(http.StatusOK,"write a cookie")
    }
    ```
2. 读取单个Cookie
    
    ```
    func readCookie(c echo.Context) error {
        cookie,err:=c.Cookie("username")
        if err != nil {
            return err
        }
        fmt.Println(cookie.Name)
        fmt.Println(cookie.Value)
        c.String(http.StatusOK,"read a cookie")
    }
    ```
3. 读取所有的Cookie
    
    ```
    func readAllCookies(c echo.Context) error {
        for _,cookie := range c.Cookies() {
            fmt.Println(cookie.Name)
            fmt.Println(cookie.Value)
        }
        return c.String(http.StatusOK,"read all cookie")
    }
    ```
    
### 错误处理
1. 错误处理程序
    * echo 提倡通过中间件货处理程序(handler)返回 HTTP 错误集中处理。集中式错误处理程序允许我们从统一位置将错误记录到外部服务，并向客户端发送自定义 HTTP 响应。可以返回一个标准的 `error` 或者 `echo.*HTTPError`
    * 例如当基本身份验证中间件找到无效凭据时，会返回401未授权错误，并终止当前的 HTTP请求
        
        ```
        e.Use(func(next echo.HandlerFunc) echo.HandlerFunc){
            return func(c echo.Context) error {
                return echo.NewHTTPError(http.StatusUnauthorized)  
            }
        }
        ``` 
2. 默认 HTTP 错误处理程序
    * Echo 提供了默认的 HTTP 粗五处理程序，他用json格式发送错误
    * 标准错误 `error` 的响应时 `500 -Internal Server Error`，在调试(debug)模式下，原始的错误信息会被发送。如果错误是 *HTTPError ，则使用设置的状态代码和消息发送响应。如果启用了日志记录，则还会记录错误信息
3. 自定义 HTTP 错误处理程序
    * 通过 `e.HTTPErrorHandler` 可以设置自定义的HTTP错误处理程序(`error handler`)
    * 利用自定义 HTTP 错误处理程序，可以显示不同种类的错误页面的同时，记录错误日志
        
    ```
    func customHTTPErrorHandler(err error,c.echo.Context){
        code := http.StatusInternalServerError
        if he,ok := err.(*echo.HTTPError);ok {
            code = he.Code 
        }
        errPage := fmt.Sprintf("%d.html",code)
        if err := c.File(errPage);err != nil {
            c.Logger().Error(err)
        }
        c.Logger().Error(err)
    }
    e.HTTPErrorHandler = customHTTPErrorHandler
    ```
    
### 请求
1. 数据绑定
    
    ```
    User struct {
        Name string  `json:"name" from:"name" query:"name"`
        Email string `json:"email" from:"email" query:"email"`
    }
    
    func(c echo.Context) (err error) {
        u := new(User)
        if  err=c.Bind(u);err != nil {
            return
        }
        return c.JSON(http.StatusOK,u)
    }
    ```
2. JSON 数据
    
    ```
    curl 
    -X POST \
    http://localhost:1323/users \
    -H 'Content-Type: application/json'
    -d '{"name":"Joe","email":"joe@labstack.com"}'
    ```
3. FORM 表单数据
    
    ```
    curl 
    -X POST \
    http://localhost:1323/users \
    -d 'name=Joe‘ \
    -d 'email=joe@labstack.com'
    ```
4. 查询参数（Query Parameters）
    
    ```
    curl 
    -X GET \
    http://localhost:1323/users\?name\=Joe\&email\=joe@labstack.com
    ```
    
### 响应
1. 发送 String 数据
    
    ```
    func (c echo.Context) error {
        return c.String(http.StatusOK , "Hello world !")
    }
    ```
2. 发送 HTML 响应（参考模板）
    
    ```
    func (c echo.Context) error {
        return c.HTML(http.StatusOK , "<strong>Hello world !</strong>")
    }
    ```
3. 发送 JSON 数据
    
    ```
    func(c echo.Context) error {
        u := &User{
            Name:"Jon",
            Eamil:"jon@labstack.com",
        }
        c.Response().Header().Set(echo.HeaderContentType , echo.MIMEApplicationJSONCharsetUTF8)
        c.Response().WriteHeader(http.StatusOK)
        return json.NewEncoder(c.Response().Encode(u))
    }
    ```
    JSON 美化
    
    ```
    func(c echo.Context) error {
        u := &User{
            Name:"Jon",
            Eamil:"jon@labstack.com",
        }
        return json.Pretty(http.StatusOK,u," ")
    }
    ```
    
    `curl http://localhost:1323/users/1?pretty`
    
    JSON Blob (`Content#JSONBlob(code int,b []byte)`) 可以从外部源（例如数据库）直接发送预编码的JSON对象
    
    ```
    func (c echo.Context) error {
        encodedJSON :=[]byte{}
        return c.JSONBlob(http.StatusOK,encodedJSON)
    }
    ```
    
    发送JSONP数据（`Context#JSONP(code int, callback string, i interface{})`）可以将Golang的数据类型转换成JSON类型，并通过带有状态码的JSONNP结构发送    

4. 发送 XML 数据
    
    ```
    func(c echo.Context) error {
        u := &User{
            Name:"Jon",
            Email:"jon@labstack.com",
        }
        return c.XML(http.StatusOK,u)
    }
    ```
    
    XML 
2. 发送文件
3. 发送附件
4. 发送内嵌
5. 发送二进制文件
6. 发送流
     