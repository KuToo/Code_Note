# Elasticsearch 视频教程笔记

## 基础篇

### 第一节

#### 1.1 Elasticsearch概念
1. 一个采用Restful API标准的、**高扩展性**和**高可用性**的、**实时**数据分析的全文搜索工具
2. 几个重要概念
    * **Node(节点):**单个的装有 Elasticsearch服务并且提供故障转移和扩展的服务器
    * **Cluster(集群):**一个集群就是由一个或多个node组织在一起，共同工作，共同分享整个数据，具有负载均衡功能的集群
    * **Document(文档):**一个文档是一个可被索引的基础信息单元
    * **Index(索引):**索引就是一个拥有几分相似特征的文档的集合
    * **Type(类型):**一个索引中，你可以定义一种或者多种类型
    * **Field(列):**Field是Elasticsearch的最小单位，相当于数据的某一列
    * **Shards(分片):**Elasticsearch将索引分成若干份，每一部分就是一个shard
    * **Replicas(复制):**Replicas是索引一份或多份拷贝

3. Mysql与Elasticsearch对比

| 关系型数据库Mysql | 非关系型数据库Elasticsearch |
| --- | --- |
| 数据库 Database | 索引 Index |
| 表 Table | 类型 Type |
| 数据行 Raw | 文档 Document |
| 数据列 Column | 字段 Field |

#### 1.2 Elasticseach架构

### 第二节 Restful 以及CURL的介绍

#### 2.1 RestfulAPI介绍
1. Restful:Representnation State Transfer(表现层状态转化)，操作资源（Get,Post,Put,Delete...）
2. API：Application Programming Interface的缩写，中文意思就是应用程序接口
3. XML：可扩展语言，是一种程序与程序之间传输数据的标记语言
4. JSON：英文JavaScript Object Notation的缩写，它是一种新型的轻量级数据交换格式

#### 2.2 Curl讲解
1. 就是以命令的方式来执行HTTP协议的请求工具
2. 可以通过Curl操作HTTP的Get/Post/Put/Delete方法
3. 使用示例

    ```
    # 访问一个网页
    curl www.baidu.com
    # 访问并保存一个网页
    curl -o baidu.html www.baidu.com
    # 显示http response的头信息
    curl -i www.baidu.com
    # 显示一次http请求的通信过程
    curl -v www.baidu.com  //粗略
    curl --trace output.txt www.baidu.com //更及详细
    # Curl执行Get/Post/Delete操作
    curl -X GET/POST/PUT/DELETE www.example.com
    ```
    
#### 2.3 Elasticsearch API文档
[文档链接:https://www.elastic.co/guide/en/elasticsearch/reference/current/docs.html](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs.html)

### 第三节 Elasticsearch

#### 3.1 Elasticsearch 安装
1. Docker安装

    ```
    # 下载镜像
    docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
    # 启动容器（单节点）
    docker run --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2
    ```
    
1. Linux
    * 配置jdk环境：下载jdk包，解压，配置环境变量
    * 下载ES的tar包，解压，添加环境变量
2. Windows
3. MacOS

4. 目录介绍

| 文件夹 | 作用 |
| --- | --- |
| /bin | 运行Elasticsearch实例和管理插件的一些脚本 |
| /Config | 配置文件路径，包含Elasticsearch.yml |
| /data | 在节点上每个索引/碎片的数据文件的位置，可以有多个目录 |
| /lib | Elasticsearch使用的库 |
| /logs | 日志的文件夹 |
| /Plugins | 已经安装的插件存放位置 |

#### 3.1 Elasticsearch 相关插件
安装插件：/bin/plugin -install 插件
1. Head插件：是一个 Elasticsearch 集群管理插件
    * github地址：[https://github.com/mobz/elsticsearch-head](https://github.com/mobz/elsticsearch-head)
    * 安装：`/bin/plugin -install mobz/elsticsearch-head`
    * 查看：[localhost:9200/_plugin/head](localhost:9200/_plugin/head)
2. Bigdesk插件:是一个 Elasticsearch 集群监控工具，可以通过它来查看集群的各种状态，如：cpu,内存使用情况，索引数据，搜索情况，http连接数等信息
    * github地址：[https://github.com/lukas-vlcek/bigdesk](https://github.com/lukas-vlcek/bigdesk)
    * 安装：`/bin/plugin -install lukas-vlcek/bigdesk/2.5`
    * 查看：[localhost:9200/_plugin/head](localhost:9200/_plugin/bigdesk)
    



