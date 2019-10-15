#### ORM
`
    /*
    //数据库插入操作
    //1.有orm对象
    o := orm.NewOrm()
    //2.有一个要插入数据的结构体对象
    user := models.User{}
    //3.对结构体对象赋值
    user.Name = "111"
    user.Pwd = "222"
    //4.插入数据操作
    o.Insert(&user) //返回影响行数和错误
    //查询数据操作
    //1.有orm对象
    o1 := orm.NewOrm()
    //2.查询的对象
    user1 := models.User{}
    //3.指定查询对象的字段值
    user1.Id = 1
    //4.查询数据操作
    o1.Read(&user1) //返回数据和错误
    user1.Name = "111"
    o1.Read(&user1, "Name") //返回数据和错误
    //数据库更新操作
    //1.有orm对象
    o2 := orm.NewOrm()
    //2.有一个要更新数据的结构体对象
    user2 := models.User{}
    //3.查到需要更新的数据
    user2.Id = 1
    err := o2.Read(&user2) //返回数据和错误
    if err==nil {
        //4.给数据重新赋值
        user.Name="333"
        user.Pwd="333"
        //5.更新
        o2.Update(&user)//返回影响行数和错误
    }
    //数据库删除操作
    //1.有orm对象
    o3 := orm.NewOrm()
    //2.删除的对象
    user3 := models.User{}
    //3.指定删除哪一条
    user3.Id = 1
    //4.删除
    o3.Delete(&user)//返回影响行数和错误
    */
`