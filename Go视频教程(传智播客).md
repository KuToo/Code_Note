### pointer(指针)
- 指针时地址，指针变量是存储地址的变量
- *p:解引用和间接引用
- 栈帧：
    1.用来给函数运行提供内存空间，取内存于stack（栈）上；
    2.当函数运行时，产生栈帧，函数调用结束时，释放栈帧
    3.栈帧存储：(1)局部变量；(2)形参；（形参与局部变量存储地位等同）(3)内存字段描述值
- 指针实用注意：
    1.空指针：未被初始化的指针
    2.野指针：被一片无效的地址空间初始化
- 变量存储:
    1.等号左边的变量，代表变量所指向的内存空间（写操作）；
    2.等号右边的变量，代表变量内存空间存储的数据值（读操作）。
- 指针的函数传参
    1.传地址（引用）：将形参的地址值作为函数参数
    2.传值（传输局）：将实参的值拷贝一份给形参
    3.传引用：在A栈帧内部，修改B栈帧的变量值
    函数传参：值传递，实参将自己的值拷贝一份给形参，可以间接的实现传地址（将碧昂凉的地址的值作为参数传递）

### slice(切片)
- 为什么用切片
    1.数组的容量固定，不能自动拓展
    2.值传递，数组作为函数的参数时，将整个数组值拷贝一份给形参
    3.在go语言中，我们几乎可以使用切片替换数组。
- 切片的本质
    1.不是一个数组的指针，是一种数据结构体，用来操作数组内部元素
- 切片的使用
    1.创建数组时，[]内需要指定数组长度，创建切片时，[]为空或者是...
    2.切片名称 [low:high:max] 其中，low:起始下标位置，high：结束下标位置len=high-low 容量：cap=max-low
    3.截取数组，初始化切片时，切片容量跟随源数组（切片）
    4.省略部分参数：
    - （1）s[:high:max],从0开始，到high结束（不包含）
    - （2）s[low:],从low开始到末尾
    - （3）s[:high],送0开始,容量跟随原先容量【常用】
- 创建切片
    1.自动推导类型创建：```slice :=[]int{1,2,4,6}```
    2.slice :=make([int],len,cap)
    3.slice:=make([]int,len) *此时，创建切片时，没有指定容量，容量跟随长度*
- 切片操作
    1.追加元素：append，语法：append(切片，待追加的元素)，向切片追加元素时，切片的容量会自动增长。1024以下时，以两倍方式增长。
    2.拷贝：copy,语法：copy(目标位置切片，源切片)，拷贝过程中，直接对应位置拷贝

### map(字典，映射)
- 定义：key ==> value 
- 注意点
    1.在一个map里面，所有的键都是唯一的，而且支持 **==** 和 **!=** 操作符
    2.key:唯一，无序。不能是引用类型数据（切片，函数以及包含切片的结构体）
- 创建map
    1.var m1 map[int]string,只是申明一个map，没有初始化，为空(nil)map，不能存储数据
    2.m2 := map[int]string{}，能存储数据
    3.m3 :=make(map[int]string)
    4.m4 :=make(map[int]string,cap)//cap容量(长度),不能使用cap()函数
- 初始化map
    1.```var m5 map[int] string = map[int]string{1:"luffy",250:"sanji",130:"zoro"}```
    2.```m5 := map[int]string{1:"luffy",250:"sanji",130:"zoro"}```
- 赋值map
    1.m2[1]="lucy",赋值过程中，新的map元素与原map元素key相比较，相同则覆盖，不同则添加
- 使用map
    1。遍历map：for range，如果是一个参数，则为key,如果只要value省略key,则可以使用*_,value*
    2.判断map中key是否存在：if v,of :=m9[1]{}，返回两个值，第一个是value,第二个是bool,代表key是否存在
    3.删除map的元素：delete(待删除元素的map,key),删除一个不存在的key不会报错
    4.拷贝：copy,语法：copy(目标位置切片，源切片)，拷贝过程中，直接对应位置拷贝

### 结构体(struct)
- 定义一种数据类型,定义以后等价于int,string等基础数据类型
- 定义和初始化
    1.定义结构体
    ```
        type Person struct{
            name string
            sex byte
            age int
        }
    ```
    2.顺序初始化：依次讲结构体内部所有成员变量初始化
    ```
        var man Persion=Person{"andy",'m',20}
    ```
    3.指定成员初始化：未初始化的变量取该数据类型的默认初始值
    ```
        man := Person{name:"andy",age:20}
    ```
- 结构体的使用
    1.访问成员变量使用**.**参数，例如 man.name
    2.比较结构体（只能== 和 !=,不能>,<,>=,<=）
    ```
        p1 :=Person{"andy","m",20}
        p2 :=Person{"andy","m",20}
        p3 :=Person{"andy","m",30}
        p1==p2 //true
        p2==p3 //false

        unsafe.Sizeof(p1)//获取结构体大小
    ```
    3.**相同类型**结构体整体赋值
    ```
        var tmp Person
        tmp=p1 
    ```
    4.函数内部使用结构体传参(值传递,几乎不用，内存消耗过大)
    ```
        func test(man Person){
            man.name="test"
            man.age=100
        }
    ```
    5.结构体地址：结构体地址==结构体首个元素的地址
- 结构体的指针使用
    1.顺序初始化：依次讲结构体内部所有成员变量初始化
    ```
        var man *Persion = &Person{"andy",'m',20}
    ```
    2.指定成员初始化：未初始化的变量取该数据类型的默认初始值
    ```
        man := &Person{name:"andy",age:20}
    ```
    3.new(Person)
- 指针作为函数的返回值
    1.通过函数返回值，初始化结构体
    2.不能返回局部变量的地址值----局部变量保存在栈帧上，函数调用结束后栈帧释放，局部变量的地址不再受系统保护，随时可能分配给其他程序

### 字符串处理函数
- strings包[strings包](http://docscn.studygolang.com/pkg/strings/)
- Split函数:字符串按指定分隔符拆分，返回切片
    ```
        str := " I love work and I love my family"
        ret := strings.Split(str," ")
    ```
- Fields函数：字符串按空格拆分，返回切片
    ```
        str := " I love work and I love my family"
        ret := strings.Fields(str)
    ```
- HasSuffix：判断字符串结束标记
    ```
        str := "test.abc"
        ret := strings.HasSubffix(str,'.abc')
    ```
- HasPrefix：判断字符串起始标记
    ```
        str := "test.abc"
        ret := strings.HasPrefix(str,"tes")
    ```

### 文件处理（os包）
- 创建文件：create ，其中，文件不存在则创建，文件存在则将文件清空（不能创建路径目录）
    1.参数：name ，打开文件的路径，相对路径和绝对文件
    2.返回值：打开文件的指针
    3.例子
    ```
        f,err :=os.Create("D:/test/test.txt")
        if err !=nil {
            fmt.Println("create err:",err)
            return 
        }
        defer f.close
        fmt.Println("success")
    ```
- 打开文件：Open ，其中，以只读方式打开文件，不能写入文件，文件不存在则打开失败
    1.参数：name ，打开文件的路径，相对路径和绝对文件
    2.例子
    ```
        f, err := os.Open("D:/test/test.txt");
        if err !=nil {
            fmt.Println("Open failed, error:",err)
        }
        defer f.Close()
        fmt.Println("Opend successed")

        //测试写入内容
        _,err1 :=f.WriteString("你好")
        if err1 != nil {
            fmt.Println("Write failed, error:",err1) //Write failed, error: write D:/test/test.txt: Access is denied.
        }
        fmt.Println("write successed")
    ```
- 打开文件：OpenFile ，其中，以只读,只写，读写方式打开文件
    1.参数：name ，打开文件的路径，相对路径和绝对文件，
            打开文件权限：O_RDONLY,O_WRONLY,O_RDWR
            权限级别：1-7，一般传6
    2.例子
    ```
        f, err := os.Open("D:/test/test.txt",);
        if err !=nil {
            fmt.Println("Open failed, error:",err)
        }
        defer f.Close()
        fmt.Println("Opend successed")

        //测试写入内容
        _,err1 :=f.WriteString("你好")
        if err1 != nil {
            fmt.Println("Write failed, error:",err1) //Write failed, error: write D:/test/test.txt: Access is denied.
        }
        fmt.Println("write successed")
    ```
- 写文件
    1.按字符串写：WriteString(str) ，返回写入的字符个数
    2.按位置写：
        - Seek(),修改文件的读写指针位置
            参数1：偏移量，正数向后，负数向前
            参数2：偏移起始位置，io.SeekStart文件起始位置，io.SeekEnd文件结束位置，io.SeekCurrent文件当前位置
            返回值：从起始位置开始，到当前文件读写指针位置的偏移量
        - WriteAt(),表示在文件指定偏移位置，写入[]byte,通常搭配Seek()
            参数1：表示要写入的数据
            参数2：偏移位置
            返回值：实际写入的字节数
    3.按字节写：
- 读文件
    reader :=bufio.NewReader(f)
    1.按行读：buf,err :=reader.ReadBytes('\n')
        - 创建一个带有缓冲区的reader(读写器)：reader= reader :=bufio.NewReader(打开的文件指针)
        - 从reader的缓冲区，读取指定长度的数据，数据长度取决于参数dlime(分隔符)：buf,err :=reader.ReadBytes('\n')
        - 判断达到文件结尾：if err!=nil && err==io.EOF 到文件结尾
    2.按字节读取Read()
- 缓冲区
- 虚拟内存映射
- 打开目录:OpenFile()
    参数1：name,打开文件的路径（绝对路径或者相对路径）
    参数2：打开文件的权限 *O_RDONLY* *O_WRONLY* *O_RDWR*
    参数3：os.ModeDir
    返回值：返回一个文件指针，可以用来读写目录
- 读目录：ReadDir(n int) Fileinfo,error
    参数：n ：读取目录项的个数，-1 代表所有
    返回值：文件信息和错误

### Go并发编程
- 并发与并行
    1.并行（parallel）：在同一时刻，有多条指令在多个处理器上同时执行
        (1).借助多核CPU实现
    2.并发（concurrency）：在同一时只能有一条指令执行，但多个进程指令被快速的轮换，使得在宏观上具有多个进程同时执行的效果
        (1).宏观：用户体验上，程序在并行执行
        (2).微观：多个计划任务，顺序执行，在飞快的切换，轮换使用CPU的时间轮片
- 程序和进程
    1.程序：编译成功的二进制文件，占用磁盘空间
    2.进程：运行起来的程序，占用内存，最小的资源分配单位
    比拟：程序->剧本(占用纸) 进程->戏(舞台、灯光、演员、道具...)
- 进程状态（5种）
    1.初始态 2.就绪态 3.运行态 4.挂起态 5.终止态
- 线程和进程的关系
    1.线程LWP：轻量级的进程，本质依然是进程,最小的执行单位
    2.线程同步：协同步调，规划先后顺序
- 协程：轻量级线程，目的是提高程序执行效率
- 进程，线程，协程总结
    1.都可以完成并发
    2.进程稳定性小，线程节约资源，协程效率高

#### Go并发（goroutine和channel）
- Goroutine
    1.创建于进程中，直接使用关键字go，放置于函数调用前面，产生一个Go程，并发。
- Goroutine的使用
    1.Goroutine特性：主Go程结束，子Go程也会自动退出。
- runtime包
    1.runtime.Gosched()函数：出让当前Go程所占用的cpu时间片，调度器安排其他等待的任务运行，并在下次在获得cpu时间片的时候，从该出让cpu的位置恢复执行
    2.runtime.Goexit()函数：退出Go程
        return:返回当前函数调用到调用者，return之前注册的defer都生效 
        Goexit:结束调用该函数的当前Go程，Goexit之前注册的defer都生效
    3.runtime.GOMAXPROCS()函数，用来设置当前进程可以并行计算的CPU核数的最大值，并返回之前的值

- 补充知识点 
    1.每当一个进程启动时，系统会自动打开三个文件：标准输入、标准输出、标准错误。--对应三个文件：stdin(0),stdout(1),stderr(2),进程结束，系统自动关闭这三个文件

#### 通道(channel) 
- 定义：
    1.是一种数据类型，对应一个“管道”（通道），主要用来解决Go程的同步问题以及Go程之间数据共享（数据传递）的问题
    2.Goroutine运行在相同的地址空间，因此访问共享内存必须做好同步，Goroutine奉行通过通信来共享内存，而不是共享内存来通信
    3.和map类似，channel也是一个对应make创建的底层数据结构的引用，```make(chan type,cap)```,其中cap=>缓冲区，cap=0无缓冲channel，cap>0有缓冲的channel,eg：make(chan int)
    4.channel有两个端，
        - 一端：写端（传入端）channel<-,
        - 另一端：读端（传出端）<-channel
    5.要求读端与写端必须同时满足条件，才在chan上进行数据流动，否则阻塞
- channel数据同步
    1.len(channel):channel中剩余未读取数据个数;cap(channel):通道的容量
- 无缓存channel和有缓存channel
    1.无缓存channel:
        - 通道容量为0，len=0;
        - channel应用于两个Go程中，一个读一个写;
        - 具备同步的能力，读写同步。（打电话）
    2.有缓存channel:
        - 通道容量大于0;
        - channel应用于两个Go程中，一个读一个写;
        - 缓冲区可以进行数据的存储，存储至容量上限，阻塞，具备一部能力，不需要同事操作channel缓冲区（发短信）。
    3.比较
- 关闭channel
    1.使用close(ch)关闭channel      
    2.确定不向对端发送数据，关闭channel
    3.对端可以判断channel是否关闭
        ```
            if num,ok := <-ch;ok==true{
                //如果对端没有关闭,ok-->true,numa保存独到的数据
            }else{
                //如果对端已经关闭,ok-->false,num无数据
            }
        ```
    4.可以使用range替代ok
        ```
            for num:=range ch{//ch不能替换为<-ch}
        ```
    5.数据不发送玩，不应该关闭
    6.已经关闭的channel不能再向其写数据。报错:panic:send on closed channel
    7.已经关闭的channel可以向其读数据。先读缓冲的数据，读完之后可以继续读取数据，值为0 
- 单向channel
    1.默认的channel是双向的 var ch chan int / ch :=make(chan int)
    2.单向channel分为：
        - 单向写channel：var sendCh chan <- int / sendCh := make(chan<- int)
        - 单向读channel：var recvCh <- chan int / recvCh := make(<-chan int)
    3.转换：
        - 双向channel可以隐式转换为任意一种单向channel sendCh = ch  
        - 单向channel不能转换为双向channel
    4.传参（传引用）  
- 生产者消费者模型
    1.生产者：发送数据端
    2.公共区(缓冲区):
        - 解耦，降低生产者和消费者之间的耦合度
        - 处理并发，生产者和消费者数量不对等时，能保持正常通信
        - 缓存，生产者和消费者数据处理速度不一致时，暂存数据
    3.消费者：接受数据端
- 定时器Timer
    
    > 1.创建定时器，指定定时时长，定时到达后，系统会自动向定时器的成员C写系统当前时间（对chan的写事件）    
    > 2.读取Timer.C得到定时后的系统时间，并且完成一次chan的读操作    
    > 3.类型解析：
    ```
        type Timer struct {
            c <-chan Time
            r tuntimeTimer
        }
    ```
