
# Jenkins 部署

## 一、前置准备工作

### 1. 安装 JRE
``` bash
# 因为不需要Java的开发环境，所以不用装JDK而是直接装JRE即可

# 安装 JER 17
$ sudo apt update
$ sudo apt install -y openjdk-17-jre

# 安装的版本验证
$ java -version
```

### 2. 安装 Jenkins
``` bash
# 在安装之前需要安装一些必要的工具
$ sudo apt install -y ca-certificates curl gnupg net-tools

# ======================== 使用官方源 ================================ Begin
# 使用官方源 国内很慢
# 导入官方 GPG 公钥
# 把 Jenkins 官方在 2023 年新换的 GPG 公钥下载到 /usr/share/keyrings/jenkins-keyring.asc
$ sudo wget -O /usr/share/keyrings/jenkins-keyring.asc \
  https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key

# 写入软件源列表
# 在 /etc/apt/sources.list.d/jenkins.list 里新增一条 APT 源记录，告诉系统“以后还可以去 https://pkg.jenkins.io/debian-stable 找软件”
$ echo "deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
  https://pkg.jenkins.io/debian-stable binary/" | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null

# 安装
$ sudo apt update
$ sudo apt install -y jenkins=2.361.1

# ======================== 使用官方源 ================================ End

# ======================== 使用国内源 ================================ Begin
# 清华大学镜像站: https://mirrors.tuna.tsinghua.edu.cn
# 阿里镜像站: https://developer.aliyun.com/mirror

# 国内镜像站的好像都只能手动进行安装了
# 下载对应的安装包
$ wget https://mirrors.tuna.tsinghua.edu.cn/jenkins/debian-stable/jenkins_2.504.3_all.deb

# 安装deb包
sudo dpkg -i jenkins_2.504.3_all.deb

# 修复可能的依赖问题
sudo apt --fix-broken install

# ======================== 使用国内源 ================================ End

# 检查安装的版本
$ unzip -p /usr/share/java/jenkins.war META-INF/MANIFEST.MF | grep "Jenkins-Version"

# 检查 Jenkins 服务状态
$ sudo systemctl status jenkins

# 启动服务
$ sudo systemctl start jenkins

# 查看是否开机自启动
$ systemctl is-enabled jenkins

# 设置开机自启动
$ sudo systemctl enable jenkins

# 检查Jenkins是否创建了系统用户
$ id jenkins

# 检查服务端口是否监听
$ netstat -tulpn | grep 8080

# 查看防火墙状态 未启动: inactive  启动: active
$ sudo ufw status
# 如果防火墙是开启的状态需要添加一条允许8080端口的规则
$ sudo ufw allow 8080
```

## 二、初始化环境

### 1. 解锁Jenkins​​
``` bash
# 查看 Jenkins 的初始默认密码
$ sudo cat /var/lib/jenkins/secrets/initialAdminPassword
# 密码内容
# 5c3f38f999854f3e83ca6d110878da2a
```

``` txt
在 Windows 机器上使用浏览器进入 Jenkins 界面:
> http://<虚拟机IP>:8080
> 示例: http://192.168.234.130:8080
```

### 2. 安装推荐插件
密码输入完毕后，就会有一个安装推荐插件的按钮，点击即可，等待其所有插件安装完毕

### 3. 创建管理员账户
这个不建议跳过，顺着引导设置即可

## 三、配置打包

### 1. 参考文档
- https://blog.csdn.net/weixin_39655220/article/details/146591174
- https://cloud.tencent.com/developer/article/2085114?policyId=1004
- https://blog.csdn.net/weixin_53510183/article/details/144584490

### 2. 安装所需插件(Plugins)
- ✅ Gitee
- ✅ SSH Agent（远程部署时需要）
- ✅ Go（管理Go环境）
- ✅ Docker（可选，如需容器化部署）
- ✅ Publish Over SSH

### 3. 在jenkins服务的虚拟机上生成ssh-key
``` bash
# 生成ssh-key
$ ssh-keygen -t rsa

# 查看私钥内容
$ cat ~/.ssh/id_rsa
```


