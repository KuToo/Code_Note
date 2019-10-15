#### Map学习

#### Map介绍
Hash表是一种巧妙并且实用的数据结构，是一个无序的key/value对的集合，其中所有的key都是不同的，通过给定的key可以在常数时间复杂度内检索、更新或删除对应的 value 。Map其实是一个 Hash 表的引用，能够基于键快速检索出数据，键就像索引一样指向与该键关联的值。以后有机会再给大家讲Map底层的东西，教会大家如何使用Map才是这一节的重点，记住一点：Map存储的是无序的键值对集合。

##### 创建与初始化
- 使用Make函数创建，make可以创建切片，也可以用来创建Map。规则是这样的：
    `
        //m := make(map[keyType]valueType)
        month := make(map[string]int)
        //方式1
        month["January"] = 1
        month["February"] = 2
        month["March"] = 3
        //方式2
        month := map[string]int{"January":1,"February":2,"March":3}
        // 方式3
        month := map[string]int{
            "January":1,
            "February":2,
            "March":3,
        }
    `
- 使用字面量也可以创建空Map，大括号里面不赋值就可以了：
    `
        month := map[string]int{}
        fmt.Println(month)        // 输出：map[]
    `
- 有空Map，是不是有nil Map？当然是有为nil的Map:
    `
        var month map[string]int
        fmt.Println(month == nil)    // 输出：true
    `
- 对于nil的map是不能存取键值对的，否则就会报错panic: assignment to entry in nil map。可以使用提到的make函数，为其初始化：
    `
        var month map[string]int
        month = make(map[string]int)
        month["January"] = 1
        fmt.Println(month)   // 输出：map[January:1]
    `
- 自然能想到，Map的零值就是nil，Map就是底层Hash表的引用。
Map的key可以是内置类型，也可以是结构类型，只要可以使用 == 运算符做比较，都可以作为key。**切片**、**函数**以及包含切片的结构类型，这些类型由于具有引用语义，不能作为key，使用这些类型会造成编译错误：
    `
        month := map[[]string]int{}
        // 编译错误：invalid map key type []string
    `
- 对于Map的value来说，就没有类型限制，当然也没有任何理由阻止用户使用切片作为Map的值：
    `
        m := map[string][]int{}
        slice := []int{1,2,3}
        m["slice"] = slice
        fmt.Println(m["slice"])
        // 或者
        slice := []int{1,2,3}
        m := map[string][]int{"slice":slice}
        fmt.Println(m["slice"])
    `

#### 如何使用Map
- Map的使用就很简单了，类似于数组，数组是使用索引，Map使用key获取或修改value。
    `
        m := map[string]int{}
        m["January"] = 1        // 赋值
        fmt.Println(m)            // 输出：map[January:1]
        m["January"] = 10       //修改
        fmt.Println(m)          // 输出：map[January:10]
        january := m["January"]   // 获取value
        fmt.Println(january)     // 输出：10
    `
- 执行修改操作的时候，如果key已经存在，则新值会覆盖旧值，上面代码已经体现出来了，所以key是不允许重复的。
  获取一个不存在的key的value的时候，会返回值类型对应的零值，这个时候，我们就不知道是存在一个值为零值的键值对还是键值对就根本不存在。好在，Map给我们提供了方法：
    `
        february,exists := m["February"]
        fmt.Println(february,exists)   // 输出：0 false
    `
  获取值的时候多了一个返回值，第一个返回值是value，第二个返回值是boolean类型变量，表示value是否存在。这给我们判断一个key是否存在就提供了很大便利。

#### Delete-- 删除键值对
- 不像Slice，Go为我们提供了删除键值对的功能 -- delete函数。
    函数原型：
    `
        func delete(m map[Type]Type1, key Type)
    `
- 第一个参数是Map，第二个参数是key。
    `
        m := map[string]int{
            "January":1,
            "February":2,
            "March":3,
        }
        fmt.Println(m)     // 输出：map[March:3 January:1 February:2]
        delete(m,"January")
        fmt.Println(m)     // 输出：map[February:2 March:3]
    `
- 删除一个不存在的键值对时，delete函数不会报错，没任何作用。

#### 遍历Map
- Map没法使用for循环遍历，跟数组、切片一样，可以使用range遍历。
    `
        for key, value := range m {
            fmt.Println(key, "=>", value)
        }
        输出：
        February => 2
        March => 3
        January => 1
    `
  可以使用空白操作符_忽略返回的key或value。多次执行代码的时候，你会发现，返回值的顺序有可能是不一样的，也就是说Map的遍历是无序的。

#### len函数
- 可以使用len函数返回 Map 中键值对的数量：
    `
        fmt.Println("len(m) =",len(m))
    `

#### Map是一种引用类型
- Map是对底层数据的引用。编写代码的过程中，会涉及到Map拷贝、函数间传递Map等。跟Slice类似，Map指向的底层数据是不会发生copy的。
    `
        m := map[string]int{
            "January":1,
            "February":2,
            "March":3,
        }
        month := m    
        delete(month,"February")
        fmt.Println(m)
        fmt.Println(month)
        输出：
        map[January:1 March:3]
        map[January:1 March:3]
    `
  上面的代码，将Map m赋值给month，删除了month中的一个键值对，m也发生了改变，说明Map拷贝时，m和month是共享底层数据的，改变其中一方数据，另一方也会随之改变。类似，在函数间传递Map时，其实传递的是Map的引用，不会涉及底层数据的拷贝，如果在被调用函数中修改了Map，在调用函数中也会感知到Map的变化。
- 那如果我真想拷贝一个Map怎么办？
    `
        month := map[string]int{}
        m := map[string]int{
            "January":1,
            "February":2,
            "March":3,
        }
        for key,value := range m{
            month[key] = value
        }
        delete(month,"February")
        fmt.Println(m)
        fmt.Println(month)
        输出：
        map[January:1 February:2 March:3]
        map[January:1 March:3]
    `
  上面的代码，我们使用range将m的键值对循环赋值给了month，然后删除month其中一个键值对，通过打印的结果可以看出，m没有改变。这就实现了真正的拷贝。