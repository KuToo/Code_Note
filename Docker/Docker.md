#### Docker基础
##### Docker的基本组成
- Docker Client 客户端
- Docker Daemon 守护进程
- Docker Image 镜像
- Docker Container 容器
- Docker Registry 仓库

##### 容器基本操作
 - 1.启动容器
    (1)执行单次命令：docker run IMAGE [COMMAND] [ARG...]
    eg:docker run ubuntu echo "hello world"
    (2)交互式：docker run -i -t IMAGE /bin/bash 
    eg:docker run -i -t ubuntu /bin/bash
- 2.查看容器
    docker ps  查看正在运行的容器
    docker ps -a 查看所有的容器
    docker ps -l 查看最近的容器
    docker inspect container_id(名称)
- 3.自定义容器名
    启动时使用-name参数：docker run --name=container IMAGE /bin/bash
- 4.重新启动停止的容器
    docker start [i(交互式)] 容器名
- 5.删除容器
    docker rm 容器名（只能删除已经停止的容器）

##### 守护式容器
- 1.以守护形式运行容器
    启动：docker run -it IMAGE /bin/bash CTRL+P + CTRL+Q 退出后台运行
    重新进入：docker attach 容器名
- 2.启动时使用 -d参数运行
    docker run -itd IMAGE /bin/bash 
- 3.查看容器日志
    docker logs [-f][-t] [--tail] 容器名
        -f --follows =true|false 默认false 一直跟踪
        -t --timestamp =true|false 默认false 加上时间戳
        --tail 显示几行 ，默认全部
- 4.查看容器内进程
    docker top 容器名
- 5.在运行中的容器内启动新进程
    docker exec [-d][-i][-t] 容器名 [COMMAND][ARG...]
- 6.停止守护式容器
    docker stop 容器名 （发送信号等待停止）
    docker kill 容器名（直接停止）
- 7.查看容器端口映射情况
    docker port 容器名

##### 在容器中部署静态网站
- 1.设置容器的端口映射
    启动容器时使用-p参数
        -p 将容器暴露的所有端口进行映射 eg:docker run -P -i -t ubuntu /bin/bash
        -p 将容器指定端口进行映射 eg:docker run -p 80 -i -t ubuntu /bin/bash
    4种映射形式
    （1）docker run -p 80 -i -t ubuntu /bin/bash 只指定容器的端口（宿主机端口随记映射）
    （2）docker run -p 8080:80 -i -t ubuntu /bin/bash 同时指定容器和宿主机的端口
    （3）docker run -p 0.0.0.0:80 -i -t ubuntu /bin/bash 只指定容器的IP和端口
    （4）docker run -p 0.0.0.0:8080:80 -i -t ubuntu /bin/bash 同时指定容器和宿主机的IP和端口
- 2.Nginx部署流程
    （1）创建映射端口的交互式容器
        docker -p 80 -it --name=web ubuntu /bin/bash
    （2）安装Nginx
        apt-get update && apt-get install -y nginx
    （3）安装Vim
        apt-get install -y vim
    （4）创建静态页面
        vim /var/www/html/index.html
    （5）修改Nginx配置文件
        vim /etc/nginx/site-enabled/default
        root   /var/www/html
    （6）运行Nginx
        nginx
    （7）验证网站访问

##### 镜像操作
- 列出镜像
    docker images [OPTIONS] [REPOSITORY]
    -a,--all=false 所有的镜像（默认不显示中间层）
    -f,--filter=[]
    --no--trunc=false 是否截断镜像id形式显示
    -q,--quiet=false 
- 查看镜像详细信息
    docker inspect [OPTIONS] IMAGE
- 删除镜像
    docker rmi [OPTIONS] IMAGE [IMAGE...]
    -f,--force=false 强制删除镜像
    --no-prune=false 不删除未打标签的父镜像
- 获取和推送镜像
    （1）查找镜像 docker search [OPTIONS] TERM
        --automated=false 只显示自动化构建的镜像
        --no-trunc=false
        --s,--stars 按star数排序显示
    （2）拉取镜像
        docker pull IMAGE:TAG
        加快拉取速度： --registry-mirror
        <1>.修改：/etc/default/docker
        <2>.修改：DOCKER_OPTS="--registry-mirror=http://MIRROE_ADDR"
    （3）推送镜像
        docker push 用户名/镜像名
- 构建镜像
    （1）docker commit [OPTIONS] 容器名称 [REPOSITORY[:TAG]]  通过容器构建
        -a,--author="" 指定镜像的作者
        -m,--message="" 指定镜像的信息
        -p,--pause=true 在构建过程中中断容器运行
    （2）docker build  [OPTIONS] PATH | URL | - 通过Dockerfile文件构建
        --force-rm=false
        --no-cache=false
        --pull=false
        -q,--quiet=false
        --rm=true
        -t,--tag=""
- 查看镜像构建过程
    docker history IMAGE

##### Dockerfile指令
- FROM <image>[:<tag>] 指定基础镜像
- MAINTAINER <name> 指定构建者的信息
- RUN <command>  指定当前镜像中运行的命令
    (shell模式) /bin/sh -c command 执行命令 eg:RUN echo hello
    (exec模式) RUN ["executable","param1","param2"] eg:RUN ["/bin/bash","-c","echo hello"]
- EXPOSE <port>[<port>...] 指定运行该镜像的容器使用的端口 （可以多次使用）
- CMD 指定启动容器时运行的命令 ，当在命令行指定了命令运行模式时，CMD指定的命令将失效
    (exec模式)  CMD ["executable","param1","param2"]
    (shell模式) CMD command param1 param2
    (与ENTERYPOINT搭配使用) CMD ["param1", ["param2"] 作为ENTERYPOINT的默认参数
- ENTERYPOINT 定启动容器时运行的命令 ，当在命令行指定了命令运行模式时，ENTERYPOINT指定的命令不会失效
- ADD 将文件复制到由dockerfile构建的镜像中
    ADD <src> ... <dst>
    ADD ["<src>" ... "<dst>"](适用于文件路径中有空格的情况)
- COPY 将文件复制到由dockerfile构建的镜像中
    COPY <src> ... <dst>
    COPY ["<src>" ... "<dst>"](适用于文件路径中有空格的情况)
- ADD vs COPY 
    ADD报刊类似tar的解压功能，如果是单纯的复制文件，推荐使用COPY
- VOLUME ["/data"] 向基于镜像创建的容器添加卷
- WORKDIR /path/to/workdir 在使用镜像启动新的容器时，在容器中指定工作目录，CMD指定的命令都会在这个目录下执行，通常指定绝对路径
- ENV <key> <value> 设置环境变量，供后面的指令使用
    ENV <key>=<value>...
- USER daemon 指定基于镜像的容器使用什么用户运行 eg:USER nginx
- ONBUIlD [INSTRUCTION] 为镜像添加触发器，当这个镜像作为其他镜像的基础镜像时，这个触发器会被执行

##### 网络互连
- docker 默认情况下允许荣期间互连的，一般通过IP，但是在每次重启容器时，IP会发生变化，所以在启动容器的时候可以使用--link选项选择要连接的容器的别名
- 在容器启动时使用 --icc-false，可以拒绝所有容器间互连
- 允许特定容器间的连接，启动选项 --icc=false --iptables=true --link

##### Docker容器与外部网络的连接
- ip_forward 系统是否会转发流量，默认值为true
    设置命令：sysctl net.ipv4.conf.all.forwarding
- iptables
    iptables是与Linux内核集成的包过滤防火墙系统，几乎所有的linux发行版本都会包含Iptables的功能。
    表(table)：
    链(chain)：
    规则(rule)：

##### Docker容器的数据卷
- 什么是数据卷
    （1）数据卷是经过特殊设计的目录，可以绕过联合文件系统（UFS）,为一个或多个容器提供访问
    （2）数据卷设计的目的，在于数据的永久化，它完全独立于容器的生命周期，因此，Docker不会再容器删除时删除其挂载的数据卷，也不会存在类似垃圾收集机制，对容器引用的数据卷进行处理
- 数据卷的特点
    （1）数据卷在容器启动时初始化，如果容器使用的镜像在挂载点包含了数据，这些数据会拷贝到新初始化的数据卷中
    （2）数据卷可以在容器之间共享和重用
    （3）容器可以对数据卷的内直接进行修改
    （4）数据卷的变化不会影响镜像的更新
    （5）卷会一直存在，即使挂载数据卷的容器已经被删除
- 为容器添加数据卷
    sudo docker run -v ~/container_data:data -it ubuntu /bin/bash
- 在Dockerfile中使用VOLUME指令添加数据卷，但是没有办法指定宿主机中映射的目录
- 数据卷容器
    （1）定义：命名的容器挂载数据卷，其他容器通过这个容器实现数据共享，挂载数据卷的容器就叫做数据卷容器
    （2）使用：docker run --volumes-from [CONTAINER NAME]
- 数据卷的备份和还原
    （1）备份
    docker run --volumes-from [container_name] -v $(pwd):/backup ubuntu
    tar cvf /backup/backup.tar [container data volume]
    过程：新启动一个容器使用数据卷容器，使用-v选项将备份目录映射到宿主机指定目录，将数据卷压缩到本分目录
    （2）还原
    docker run --volumes-from [container_name] -v $(pwd):/backup ubuntu
    tar xvf /backup/backup.tar [container data volume]

##### Docker容器跨主机网络连接
- 使用网桥实现跨主机容器连接
    （1）配置网络 vim /etc/network/interfaces
        auto lo
        iface lo inet loopback
        auto br0
        iface br0 inet static
        address 10.211.55.5
        netmask 255.255.255.0
        gateway 10.211.55.1
        bridge_ports eth0
    （2）修改docker 配置
        DOCKER_OPTS="-b=br0 --fixed-cidr='10.211.55.128/26'"
    （3）优缺点
        优点：配置简单，不依赖第三方软件
        缺点：与主机在同网段，需要小心划分IP地址，需要有网段控制权，再生产环境中不易实现不易管理
- 使用Open vSwitch实现跨主机容器连接
    （1）定义
        Open vSwitch:是一个高质量、多层虚拟交换机，使用开源Apache2.0许可协议，由Nicir Networks开发，主要实现代码为可移植的C代码。它的目的是让大规模网络自动化可以通过编程扩展，同时仍然支持标准的管理接口和协议（例如 NetFlow,sFlow,SPAN,RSPAN,CLI,LACP802.lag）
    （2）环境
        MAC OSX + Virtualbox
        两台Ubuntu14.04虚拟机
        双网卡，Host-Only & NAT
        安装Open v Switch:apt-get install opencswitch-switch
        安装网桥管理工具：apt-get install bridge-utils
        IP地址： Host1:192.168.59.103
                Host2:192.168.59.104
    （3）操作
        建立ovs网桥 
        sudo ovs-vsctl add-br obr0
        为新建的网桥添加gre连接 
        sudo ovs-vsctl add-port obr0 gre0
        sudo set interface gre0 type=gre options:remote_ip192.168.59.104
        配置docker容器虚拟网桥
        sudo ifconfig br0 192.168.1.1 netmask 255.255.255.0 
        为虚拟网桥加ovs接口
        sudo brctl addif br0 obr0
        添加不同docker容器网段路由
        sudo vim /etc/default/docker
        DOCKER_OPTS="-b=obr0"
- 使用weave实现跨主机容器连接
    （1）定义
    建立一个虚拟的网络，用于将运行在不同主机的Docker容器连接起来
    官网: <http://weave.works>
    Github网址: <https://github.com/weaveworks/weave#readme>
    （2）环境
        MAC OSX + Virtualbox
        两台Ubuntu14.04虚拟机
        双网卡，Host-Only & NAT
        IP地址： Host1:192.168.59.103
                Host2:192.168.59.104
    （3）操作
        安装weave
        sudo wget -o /usr/bin/weave https://raw.githubusercontent.com/zettio/weave/master/weave
        sudo chmod a+x /usr/bin/weave
        启动weave
        weave launch
        连接不同主机
        通过weave启动容器

#### Docker network网络
- 查看网络
    docker network ls [OPTIONS] NETWORK [NETWORK...]
        --format , -f       Format the output using the given Go template
        --verbose , -v      Verbose output for diagnostics
- 创建网络
    docker network create --subnet=192.168.0.1/24 swoft_network
        --attachable        可以手动附加容器
        --aux-address       Auxiliary IPv4 or IPv6 addresses used by Network driver
        --config-from       API 1.30+The network from which copying the configuration
        --config-only       API 1.30+Create a configuration only network
        --driver , -d   bridge  Driver to manage the Network
        --gateway       IPv4 or IPv6 Gateway for the master subnet
        --ingress       API 1.29+Create swarm routing-mesh network
        --internal      Restrict external access to the network
        --ip-range      Allocate container ip from a sub-range
        --ipam-driver       IP Address Management Driver
        --ipam-opt      Set IPAM driver specific options
        --ipv6      Enable IPv6 networking
        --label     Set metadata on a network
        --opt , -o      Set driver specific options
        --scope     API 1.30+Control the network’s scope
        --subnet        Subnet in CIDR format that represents a network segment
- 将一个容器连接到一个网络
    docker network connect [OPTIONS] NETWORK CONTAINER
        --alias     Add network-scoped alias for the container
        --driver-opt        driver options for the network
        --ip        IPv4 address (e.g., 172.30.100.104)
        --ip6       IPv6 address (e.g., 2001:db8::33)
        --link      Add link to another container
        --link-local-ip     Add a link-local address for the container
- 将一个容器从一个网络断开
  docker network disconnect  [OPTIONS] NETWORK CONTAINER
    --force , -f
- 查看某个网络的详细信息
  docker network inspect [OPTIONS] NETWORK [NETWORK...]
- 清空所有没有使用的网络
  docker network prune  
- 删除一个或多个网络
  docker network rm 

#### Docker集群
    
##### 容器技术
- LXC(Linux Container)技术
    在同一台主机运行多个隔离的Linux系统的虚拟技术
    通过Linux内核功能实现轻量级虚拟机 --命名空间，cgroups，chroot
    将Linux操作系统和一个或多个应用容器化
- 容器集群管理
    容器调度
    配置管理
    服务发现
    日志/监控/报警
- 