# 书籍 Docker实战

## 第1章

## 第2章

## 第3章

### 3.3 保存和还原工作成果
* **技巧14**:在开发中“保存游戏”的方式
    1. 问题：想要保存开发环境的状态
    2. 解决方案：docker commit 来保存状态
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        
* **技巧15**:给Docker镜像打标签
    1. 问题：想要方便的引用并且保存一次Docker提交
    2. 解决方案：使用docker tag 给这次提交命名
        `docker tag CONTAINER 标签名`
    
    | 术语 | 含义 |
    | --- | --- |
    | 镜像 | 一个只读层 |
    | 名称 | 镜像的名称 |
    | 标签 | 作为动词的话，是指给一个镜像取名字，作为名词，它是镜像名的一个修饰词 |
    | 仓库 | 一组托管的打好标签的镜像，它们一起为容器创建相应的文件系统 | 
    
* **技巧16**:在Docker Hub上分享镜像
    1. 问题：想要公开分享一个Docker镜像
    2. 解决方案：使用Docker Hub注册中心(registry)来分享镜像

    Docker 注册中心的术语
    
    | 术语 | 含义 |
    | --- | --- |
    | 用户名(username) | Docker注册中心的用户名 |
    | 注册中心(registry) | 注册中心持有镜像，一个注册中心就是一个可以上传或者下载镜像的存储。可以是公开的，也可以是私有的。 |
    | 注册中心宿主机(registry host) | 运行Docker注册中心的宿主机 |
    | Docker Hub | 托管在https://hub.docker.com 上默认的注册中心 |
    | 索引(index) | 与注册中心宿主机含义相同 |

* **技巧17**:在构建时指定特定的镜像
    1. 问题：想要确保时从一个特定的未做修改的镜像构建
    2. 解决方案：从一个特定的镜像ID构建
    
    ```
    # Dockerfile 文件
    
    #从一个指定的镜像ID构建
    FROM 8eaa4ff06b53
    #在这个镜像里运行一个命令，把构建时引用的镜像记录到新镜像的一个文件里
    RUN echo "built from image id:" > /etc/buildinfo 
    RUN echo "8eaa4ff06b53" >> /etc/buildinfo
    RUN echo "an ubuntu 14.04.01 image" >> /etc/buildinfo
    #构建的镜像默认会输出记录到 /etc/buildinfo文件里的信息
    CMD ["echo","/etc/buildinfo"]
    ```
### 3.4进程即环境
* **技巧18**:在开发中“保存游戏”的方式
    1. 问题：为了能够在需要的时候恢复到一个已知的状态，想要定期保存容器的状态
    2. 解决方案：当不确定是否可以活下来时使用`docker commit` 来“保存游戏”

## 第4章 Docker日常

### 4.1 卷--持久化问题
* **技巧19**:Docker卷--持久化问题
    1. 问题：想要在容器里访问宿主机上的文件
    2. 解决方案：使用Docker的卷表纸，在容器里访问宿主机上的文件
        
        `docker run -v /var/db/tables:/var/data -it debian bash`
        
        -v(--volume)参数表示为容器指定一个外部的卷，将外部的/var/db/tables目录映射到容器的/var/data目录

* **技巧20**:通过BitTorrent Sync的分布式卷
    1. 问题：想要在互联网上跨宿主机共享卷
    2. 解决方案：使用BitTorrent Sync镜像来共享卷
        (1)在主服务器上配置第一台宿主机上的容器
            
        ```
        # 运行ctlc/btsync镜像作为一个守护容器，映射两个端口
        docker run -d -p 8888:8888 -p 5555:5555 --name btsync ctlc/btsync
        # 获取btsync容器的输出，并记录键的内容
        docker logs btsync
        # 启动一个交互式容器，挂载btsync服务器的卷
        docker run -it --volumes-from btsync ubuntu /bin/bash
        # 添加一个文件到/data卷
        touch /data/shared_from_server_one
        ```
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          
        (2）在次服务器上同步卷
        ```
        # 启动一个btsync客户端容器作为守护进程，并且传递运行在host1上的守护进程生成的键
        docker run -d --name btsync-client -p 8888:8888 -p 5555:5555 ctlc/btsync SECRET值
        # 启动一个交互式容器，通过客户端守护进程挂载卷
        docker run -it --volume-from btsync-client ubuntu bash
        # 在host2上创建第二个文件到/data卷
        touch /data/shared_from_server_two
        ```
        
* **技巧21**:保留容器的 bash 历史
    1. 问题：想要将容器的bash历史和宿主机的命令历史共享
    2. 解决方案：给 `dcoker run` 命令起一个别名来共享容器和宿主机的bash历史 
        为了在宿主机上共享bash历史，可以运行Docker镜像时挂载一个卷
        
        ```
        # 设置 bash拾取的环境变量，这可以确保所用的bash历史文件就是挂载的那个文件
        docker run -e HIST_FILE=/root/.bash_history \ 
        # 将容器里root目录下的bash历史文件映射到宿主机上
        -v=$HOME/.bash_history:root/.bash_history \
        -it ubuntu /bin/bash
        ```
        
        可以将命令写进～/.bashrc文件，之后每次运行`docker run`命令时都会加上 `-e -e HIST_FILE=/root/.bash_history -v=$HOME/.bash_history:root/.bash_history`这来两个参数，达到共享bash的目的。

* **技巧22**:数据容器
    1. 问题：想要在容器里使用一个外部卷，但是只想让Docker访问这些数据
    2. 解决方案：启动一个数据卷容器，然后在运行其他容器时使用 --volume-from选项
        
        `docker run -v /shared-data --name dc busybox touch /shared-data/somefile`
    
        -v 参数并没有将卷映射到宿主机目录，因此它将会在容器的管辖范围内创建一个目录。这个目录通过`touch`填充一个文件，然后容器立刻退出了---一个数据容器在被使用的时候并不需要处于运行状态。
        
        `docker run -it --volumes-from dc busybox /bin/sh`
        
        --volumes-from 标志允许通过挂载它们到当前容器的形式，来引用数据容器里的文件，只需要传入卷定义的ID即可。
        (1)**纯数据容器不一定需要运行：**一个常常让人困惑的问题便是纯数据容器是否需要运行，答案是不需要！它只需要存在，在宿主机上运行过并且没有被删除。
        (2)**卷会持久化：**需要明白的重要一点是，使用这种模式会造成大量的磁盘消耗，这可能使得调试变得相当困难。由于Docker在纯数据容器里管理卷，而且不会在最后一个引用他的容器已经退出的前提下删除该卷。因此在该卷上的数据均会保留下来，这是为了防止意外的数据丢失。
        (3)**文件路径之争：**如果应用程序时从多个容器写日志到统一数据容器的话，很重要的一点便是得确保每个容器日志文件写入的是一个唯一的文件路径。如果无法保证这一点，不同的容器便有可能覆盖或者截断该文件，从而造成数据的丢失或者可能写入的数据时交错混杂的，这就很难分析文件的内容。类似的，如果对数据容器调用--volume-from，就是允许该容器潜在的覆盖自己的目录，因此也要小心这里的命名冲突。

* **技巧23**:使用 SSHFS 挂载远程卷
    1. 问题：想要挂载一个远程文件系统而无须任何服务器端的配置
    2. 解决方案：使用 SSHFS 来挂载远程文件系统
                
        ```
        # 在宿主机上启动一个容器并附加--privileged参数（成为特权容器）
        docker run -it --privileged debian /bin/bash
        # 进入容器之后更新并安装 SSHFS
        apt-get update -y && apt-get install -y sshfs
        # 按一下步骤登陆到远程宿主机
         LOCALPATH=/path/to/local/directory //选择一个远程位置对应的挂载目录
         mkdir $LOCALPATH  //创建本地挂载目录
         sshfs user@host:/path/to/remote/directory $LOCALPATH //登录
        ```
        (1)**需要root权限：**用户需要有root权限才能使用这一技巧，而且需要安装FUSE(用户空间级的文件系统内核模块)。可以在终端运行`ls /dev/fuse`查看文件是否存在来确认当前系统是否支持。
        (2)**变更不会持久化到容器：**虽然这一技巧没有用到卷功能，并且这些文件在文件系统上也是可见的，但是这一技巧并不会提供任何容器级别的持久化，任何变更都将制作用到严惩服务器的文件系统。
        
* **技巧24**:通过NFS共享数据
    1. 问题：想要通过NFS无缝访问远程文件系统
    2. 解决方案：使用一个基础设施数据容器来中转访问
        (1)NFS服务器将一个内部目录公开为一个/export文件夹,它会呗绑定挂载到NFS服务器宿主机上
        (2)Docker宿主机随后会使用NFS协议将文件夹挂载到它的/mnt文件夹
        (3)创建一个所谓的基础设施容器来绑定挂载的文件夹，该容器充当数据卷容器
        (4)其他容器启动时`volumes-from`这个基础设施容器

* **技巧25**:开发工具容器
    1. 问题：想在其他人的机器上访问自己定制的开发环境
    2. 解决方案：用自己的配置创建一个容器，然后把它放在注册中心上。
    
        ```
        # 下载工具镜像
        docker pull dockerinpractice/docker-dev-tools-image
        # 基于该镜像启动容器
        docker run -it dockerinpractice/docker-dev-tools-image
        # 挂载Docker套接字以访问宿主机上的Docker守护进程
        docker run -it -v /var/run/docker.sock:/var/run/docker.sock \
        -v /tmp/.X11-unix:/tmp/X11-unix \
        -e DISPLAY=$DISPLAY \
        --net=host \
        --ipc=host \
        -v /opt/workspace:/home/dockerinpractice
        dockerinpractice/docker-dev-tools-image
        ```
    
### 4.2 运行容器
* **技巧26**:在Docker里运行GUI
    1. 问题：想要在一个容器中运行GUI，就像其他正常的桌面应用一样
    2. 解决方案：使用自己的用户凭证和程序创建镜像，然后将用户的X服务器绑定挂载到它上面
    
        ```
        # 创建一个新目录并进入到新目录
        mkdir dockergui
        cd dockergui
        # 收集Dockerfile所需的用户id
        id
        
        # 创建Dockerfile
        FROM ubuntu
        RUN apt-get update
        RUN apt-get install -y firfox
        RUN groupadd -g GID USERNAME
        RUN useradd -d /home/USERNAME -s /bin/bash -m USERNAME -u UID -g GID
        USER USERNAME
        ENV HOME /home/USERNAME
        CMD /usr/bin/firefox
        
        # 构建GUI镜像
        docker build -t gui .
        #基于镜像启动容器
        docker run -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY+$DISPLAY gui \
        -h HOSTNAME -v $HOME/..Xauthority:/home/$USER/.Xauthority
        ```
        
* **技巧27**:检查容器
    1. 问题：想要找出一个容器的Ip地址
    2. 解决方案：使用`docker inspect`查询和过滤容器的元数据
        ```
        # 查看所有的信息
        docker inspect CONTAINER_ID/IMAGE/NETWORK ...
        # 查看镜像的原始输出
        docker inspect ubuntu | head  
        # 查看一个容器的IP地址
        docker inspect --format '{{.NetworkSetting.IPAddress}}' CONTAINER_ID
        # 查看所有正在运行中的容器的IP地址，然后逐个ping
        docker ps -q | xargs docker inspect --format '{{.NetworkSetting.IPAddress}}' xargs -11 ping c1
        ```
        
* **技巧28**:干净地杀掉容器
    1. 问题：想要干净地终止一个容器
    2. 解决方案：使用`docker stop`而不是`docker kill`
        停止和杀死容器
    | 命令 | 默认信号量 | 默认信号量的值 |
    | --- | --- | --- |
    | kill | TERM | 15 |
    | docker kill | KILL | 9 |
    | docker stop | TERM | 15 |
    
* **技巧29**:使用Docker Machine来置备Docker宿主机
    1. 问题：想要在与自己机器独立的Docker宿主机上启动容器
    2. 解决方案：使用 Docker Machine
        (1)Docker Machine：是一个便利程序，它将大量配置外部宿主机繁琐的指令包装起来，变成易于上手的命令。它不是一种Docker的集群解决方案
        (2)安装：安装程序是一个二进制文件，针对不同架构的[下载链接](https://github.com/docker/machine/release)
        (3)使用：
        
            ```
            # 创建一个VirtualBox虚拟机
            docker-machine create --driver virtualbox host1
            # 使用docker-machine env设置默认宿主机上的Docker命令
            eval $(docker-machine env host1)
            # 环境变量均以 DOCKER_ 开头
            env | grep DOCKER
            # 设置DOCKER_HOST变量，即虚拟机上Docker守护进程的端口
            DOCKER_HOST=tcp://192.168.99.101:2376
            # 设置用来处理客户端和新宿主机之间连接的安全性
            DOCKER_TLS_VERIFY=yes
            DOCKER_CERT_PATH=/home/USERNAME/.docker/machines/host1
            DOCKER_MACHINE_NAME=host1
            # docke命令现在被指向创建好的虚拟机上，而不再是之前的宿主机
            docker ps-a
            # ssh子命令将会直接连接到新的虚拟机上
            docker-machine ssh host1
            ```
        
        (4)管理宿主机
        
        | 子命令 | 行为 |
        | --- | --- |
        | Create | 创建一台新的机器 |
        | ls | 列出Docker宿主机 |
        | start | 启动机器 |
        | stop | 停止机器 |
        | restart | 重启机器  |
        | rm | 销毁一台机器 |
        | kill | 杀掉一台机器 |
        | inspect | 一JSON的格式返回机器的元数据 |
        | Config | 返回链接机器所需的配置信息 |
        | Ip | 返回一台机器的Ip地址 |
        | url | 返回一台机器上Docker守护进程的URL |
        | upgrade | 将宿主机上的Docker升级到最新版本 |

### 构建镜像
* **技巧30**:使用ADD将文件注入镜像
    1. 问题：想要以一种简便的方式下载和解压一个压缩包到镜像
    2. 解决方案：打包并压缩，随后在Dockerfile里使用ADD命令
    
* **技巧31*:重新构建时不使用缓存
    1. 问题：想要不使用缓存重新构建Dockerfile
    2. 解决方案：构建镜像时加上 --no-cache 标志

* **技巧32**:拆分缓存
    1. 问题：想要在Dockerfile构建中从一个指定的点开始失效Docker构建缓存
    2. 解决方案：在命令后面增加一条无害的注释，从而让缓存失效

### 保持阵型
* **技巧33**:运行Docker时不加sudo
    1. 问题：想要无须sudo便可以运行docker命令
    2. 解决方案：将自身加入到docker组
        
        `sudo addgroup -a username docker`
        
* **技巧34**:清理容器
    1. 问题：想要整理系统上的容器
    2. 解决方案：设置一个别名来运行清理旧容器的命令         
        (1)删除所有容器
    
        ```
        # 获取所有容器ID的列表，包括正在运行的以及已经停止的，然后将他们传给后面的命令
        docker ps -a -q \
        # docker rm -f命令，被传入的仁义容器将会被删除，即使它们还处于运行状态
        xargs --no-run-if-empty docker rm -f
        # xargs 命令会获取输入的每一行内容，并将它们全部作为参数传递给后续命令，--no-run-if-empty参数可以避免在前面的命令完全没有输出的情况下执行该命令。
        ```
        
        (2)删除已经退出的容器，保留正在运行的容器
        
        ```
        docker ps -a -q --fliter status=exited | xargs --no-run-if-empty docker rm 
        ```
        
        (3)删除异常退出的容器
        
        ```
        # comm命令用来比较两个文件的差异，-3参数将不会显示同时出现在两个文件里的行内容(这些容器的退出码都是0)，然后输出其他不同的部分
        comm -3
        # 找出退出的容器ID，给它们排序，然后以文件形式给comm
        <(docker ps -a -q --filter=status=exited | sort) \
        # 找出退出码为0的容器，给它们排序，然后以文件形式传给comm
        <(docker ps -a -q --filter=status=0 | sort) | \
        # 对非0退出码(comm命令管道的输出)的容器执行docker inspect,并将输出结果保存到error_containers文件中
        xargs --no-runif-empty docker inspect > error_containers
        ```
        
        **bash中的进程替换：**bash里的“<(命令)”语法被称为进程替换，它允许把一个命令的输出结果作为一个文件，传给其他命令，这在无法使用管道输出的时候非常有用。 
        
    
        **将单行命令设置为命令：**可以给命令设置别名，以便在登录在宿主机上后更容易执行，需要在～/.bashrc文件里加上如下代码
        
        
        `alias dockernuke='docker ps -a -q | xargs --no-run-if-empty docker rm -f' `
        然后可以使用dockernuke命令达到以上目的。（注意：可能会删除没有运行的数据容器）
    
* **技巧35**:清理卷
    1. 问题：孤立的Docker卷挂载在本地用掉了大量的磁盘空间
    2. 解决方案：在调用`docker rm`命令时加上 -v 标志(只会删除没有被其他容器挂载的卷)，或者如果忘记的话可以用一个脚本来销毁它们
        使用脚本来处理
        
        ```
        docker run \
        # 挂载Docker服务器的套接字，这样就可以在容器中调用Docker
        -v /var/run/docker.sock:/var/run/docker.sock \
        # 挂载Docker目录，这样便可以删除孤立的卷
        -v /var/lib/docker:/var/lib/docker \
        # 升级权限，这样才有权限删除孤立的卷
        --privileged dockerinpactice/docker-cleanup-volumes
        ```
        
        **恢复数据：**如果想要从一个不再被任何容器引用的未删除卷中恢复数据，需要以root身份查看/var/lib/docker/volumes文件夹里的内容 
    
* **技巧36**:解绑容器的时候不停掉它
    1. 问题：想要解绑一个容器的交互会话，同时不停掉它
    2. 解决方案：按下 CTRL+P，在按下CTRL+Q来解绑

* **技巧37**:使用DockerUI来管理Docker守护进程
    1. 问题：想要在宿主机上不通过命令形式管理容器和镜像
    2. 解决方案：使用DockerUI
    [DockerUI源码链接](https://github.com/crosbymichael/dockerui)
    使用：`docker run -d -p 9000:9000 --privileged -v /var/run/docker.sock:var/run/docker.sock dcokerui/dockerui`
    
    访问localhost:9000,就可以看到展示docker概览信息的面板
    
* **技巧36**:为Docker镜像生成一个依赖图
    1. 问题：想要以树的形式将存放在宿主机上的镜像可视化
    2. 解决方案：使用一个我们之前为此任务创建的镜像作为一条单行命令来输出一个PNG或者获取一个Web视图
    
    ```
    # 在生成镜像后删除容器
    docker run --rm \
    # 挂载Docker服务器的Unix域套接字，以便可以在容器里访问Docker服务器。如果已经改了Docker守护进程的默认配置，这将不会凑效
    -v /var/run/docker.sock:/var/run/docker.sock \
    # 指定一个镜像然后生成一个png作为产品
    dockerinpractice/docker-image-graph > docker_images.png
    ```
    
* **技巧39**:直接操作--对容器执行命令
    1. 问题：想要在一个正在运行的容器里执行一些命令
    2. 解决方案：使用`docker exec`命令
        
    Dcoker exec 模式
    
    | 模式 | 描述 |
    | --- | --- |
    | 基本 | 在命令行上对容器同步的执行命令 |
    | 守护进程 | 在容器的后台运行命令 |
    | 交互 | 运行命令并允许用户与其交互 |
    
    启动一个容器并在后台一直运行:
    
    `docker run -d --name sleeper debian sleep infinty`
    
    ```
    # 以基本模式进入容器执行命令
    docker exec sleeper echo "hello host from container"
    ```
    
    ```
    # 以守护进程模式执行命令
    docker exec -d sleeper \
    # 删除所有在最近7天没有做过更改并且以log结尾的文件
    find / -ctime7 -name '*log' -exec rm {}
    ```
    
    ```
    # 以交互模式执行命令
    docker exec -it sleeper  /bin/bash
    ```
    
## 第5章 配置管理--让一切井然有序

### 5.1 配置管理和Dockerfile

* **技巧40**:使用ENTRYPOINT创建可靠的定制工具
    1. 问题：想要在定义容器将会执行的命令，但是将命令的具体参数留给用户
    2. 解决方案：使用Dockerfile的ENTRYPOINT指令
    
    clean_log脚本：会删除超过特定参数的日志，具体参数作为一个命令行选项传入
    
    ```
    #!/bin/bash
    echo "Cleaning logs over $1 days old"
    find /log_dir -ctime "$1" --name '*.log' -exec rm {} \;
    ```
    
    删除旧的日志文件的Dockerfile
    
    ```
    FROM ubuntu:14.04
    ADD clean_log /usr/bin/clean_log
    RUN chmod +x /usr/bin/clean_log
    ENTRYPOINT {"/usr/bin/clean_log"}
    CMD ["7"]
    ```
     
     基于Dockerfile构建镜像
     
     `docker build -t log_cleaner`
     
     基于镜像启动容器
     `docker run -v /var/log/myapplogs:/log_dir log_cleaner 365`

* **技巧41**:在构建中指定版本来避免软件包的漂移
    1. 问题：想要确保deb包是自己期望的版本
    2. 解决方案：在一个已经验证安装的系统上运行一个脚本来拉取所有依赖软件包的版本，并且获取依赖包版本。在 Dockerfile里安装指定的版本
    针对版本方面的基本检查可以通过在一个已经验证过的系统上调用apt-cache来完成
    `apt-cache show nginx | grep ^Version`
    
    可以通过`--source`参数所有依赖的信息
    `apt-cache --resource depends nginx`
    
    然后在Dockerfile里指定版本
    `RUN apt-get -y install nginx=1.4.6ubuntu3`
    
    在某些时候，用户的构建会因为版本不再可用而失败，这种情况下可以看看哪些包做了改动，重新检查一下这些改动，确认是否满足特定镜像的需求。
    
* **技巧42**:用perl-p-i-e替换文本
    1. 问题：想要在构建期间修改多个文件里的特定行
    2. 解决方案：perl -p -i -e
    典型示例：perl -p -i -e 's/127\.0\.0\.1/0.0.0.0/g' *
    -p:要求perl循环处理看到的所有的行
    -i:要求perl即时更新匹配的行内容
    -e:要求perl把传入的字符串当作一个perl程序处理

* **技巧43**:镜像的扁平化
    1. 问题：想要从镜像的分层历史中移除私密信息
    2. 解决方案：基于该镜像创建一个容器，将它导出再导入，然后给它打上最初镜像ID的标签

* **技巧44**:用alien管理外来软件包
    1. 问题：想要安装一个外来的发行版的软件包
    2. 解决方案：使用一个基于alien的Docker镜像转换软件包
    alien是一款命令行工具，是专为转换不同格式的软件包文件设计的，alien支持俄包个事如下表
    
    | 扩展名 | 描述 |
    | --- | --- |
    | .deb | Debian包 |
    | .rpm | Red Hat 包管理 |
    | .tgz | Slackware Gzip压缩的TAR包 |
    | .pkg | Solaris pkg包 |
    | .slp | Stampede包 |
    
    示例：eatmydata包
    
    ```
    # 创建一个空的工作目录
    mkdir tmp && cd tpm
    # 获取想要转换的包文件
    wget http://mirrors.kernel.org/ubuntu/pool/main/libe/libeatmydata/eatmydata_26-2_i386.deb
    
    # 运行dockerinpractice/alienate镜像，将当前目录挂载到容器的/io路径下。容器会检查该目录，尝试转换找到任意有效文件 
    docker run -v $(pwd):/io dockerinpractice/alienate
    # 在自己的容器中运行alien
    docker run -it --entrypoint /bin/bash dockerinpractice/alienate
    ```

* **技巧45**:把镜像逆向工程得到Dockerfile
    1. 问题：有一个镜像，想要逆向工程得到最初的Dockerfile
    2. 解决方案：使用 `docker history` 命令并通过查看分层的方式确定更改过的地方
    示例Dockerfile
    
    ```
    FROM busybox
    MAINTAINER iblue_me@163.com
    ENV myenvname myenvvalue
    LABEL myenvname myenvvalue
    WORKDIR /opt
    RUN mkdir -p copied
    COPY Dockerfile copied/Dockerfile
    RUN mkdir -p added
    ADD Dockerfile added/dockerfile
    RUN touch /tmp/afile
    ADD Dockerfile /
    EXPOSE 80
    VOLUME /data
    ONBUILD touch /tmp/built
    ENTRYPOINT /bin/bash
    CMD -r
    ```
    
    通过以上Dockerfile构建镜像
    `docker build -t reverseme`
    
    **什么是构建上下文**:Docker构建上下文是指`docker build`命令里传入目录的位置及目录里的一组文件，该上下文也是Docker在 Dockerfile里用来查找ADD或者COPY对应文件的地方
    
### 5.2 传统配置工具和Docker



     


    
