# Docker Compose

## 简介
1. 用于定义和运行多容器Docker应用程序的工具，通过Compose，可以使用yml文件来配置应用程序需要的所有服务，然后，使用一个命令，就可以从yml文件配置中创建并启动所有服务。
2. 使用的三个步骤
    * 使用Dockerfile定义应用的环境
    * 使用Docker-compose.yml定义应用程序的服务，这样他们可以在隔离环境中一起运行
    * 最后执行docker-compose up命令来启动并运行整个应用程序
3. docker-compose.yml 的配置案例
    
    ```
    version: '3'
    services:
      web:
        build: .
        ports:
       - "5000:5000"
        volumes:
       - .:/code
        - logvolume01:/var/log
        links:
       - redis
      redis:
        image: redis
    volumes:
      logvolume01: {}    
    ```

## Compose 安装
1. 从GitHub下载二进制包 [(最新发行的版本地址)](https://github.com/docker/compose/releases)

    `sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose` 
    * 要安装其他版本的 Compose，请替换 1.24.1。
2. 将可执行权限应用于二进制文件：
    `sudo chmod +x /usr/local/bin/docker-compose`
3. 创建软链：
    `sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose`
4. 测试是否安装成功：
    `docker-compose --version`

**注意:** 对于 alpine，需要以下依赖包： py-pip，python-dev，libffi-dev，openssl-dev，gcc，libc-dev，和 make

## 使用
1. 准备测试目录
    ```
    mkdir compose_test
    cd compose_test
    ```
    
2. 在测试目录中创建一个名为 app.py 的文件
    composetest/app.py 文件代码
    ```
    import time
    import redis
    from flask import Flask
    
    app = Flask(\__name__)
    cache = redis.Redis(host='redis', port=6379)
    
    def get_hit_count():
        retries = 5
        while True:
            try:
                return cache.incr('hits')
            except redis.exceptions.ConnectionError as exc:
                if retries == 0:
                    raise exc
                retries -= 1
                time.sleep(0.5)
    
    @app.route('/')
    def hello():
        count = get_hit_count()
        return 'Hello World! I have been seen {} times.\n'.format(count)    
    ```
    
3. 在 composetest 目录中创建另一个名为 requirements.txt 的文件，内容如下：
    ```
    flask
    redis
    ``` 
    
4. 创建Dockerfile文件
    在 composetest 目录中，创建一个名为的文件 Dockerfile，内容如下：
    ```
    FROM python:3.7-alpine
    WORKDIR /code
    ENV FLASK_APP app.py
    ENV FLASK_RUN_HOST 0.0.0.0
    RUN apk add --no-cache gcc musl-dev linux-headers
    COPY requirements.txt requirements.txt
    RUN pip install -r requirements.txt
    COPY . .
    CMD ["flask", "run"]
    ```
    
5. 创建docker-compose.yml的文件
    在测试目录中创建一个名为 docker-compose.yml 的文件，然后粘贴以下内容：
    ```
    # yaml 配置
    version: '3'
    services:
      web:
        build: .
        ports:
         - "5000:5000"
      redis:
        image: "redis:alpine"   
    ``` 
    
## 使用 Compose 命令构建和运行您的应用
1. 在测试目录中，执行以下命令来启动应用程序
    `docker-compose up`
2. 如果你想在后台执行该服务可以加上 -d 参数：
    `docker-compose up -d` 
   
    