要在Ubuntu 22.04中将指定文件夹共享给本地Windows使用：

***

## 🔧 Samba共享（推荐网络环境使用）

适用场景：物理机/虚拟机均可，需网络互通。

***

### 一、 Ubuntu端操作：
1. 安装Samba
    ``` bash
    sudo apt update
    sudo apt install samba
    ```

2. 配置共享文件夹  
    右键点击目标文件夹 → 选择 “本地网络共享” → 勾选 “共享此目录”。  
    设置共享名称（避免空格），勾选 “允许其他人创建和删除文件”（需写入权限时）。

3. 设置访问用户  
    若需密码访问，终端添加Samba用户：  
    ``` bash
    语法：sudo smbpasswd -a <你的Ubuntu用户名>
    示例：sudo smbpasswd -a xiaoyin01
    ```

4. 重启服务：  
    ``` bash
    sudo systemctl restart smbd nmbd
    ```

***

### 二、Windows端访问：
    打开文件资源管理器 → 地址栏输入： \\<Ubuntu的IP地址>\<共享名称>
    PS: 上面的共享命令就是上面配置选择指定文件作为共享文件设置的名字
    例如：`\\192.168.234.128\samba_share`
    输入 xiaoyin01和密码（若设置）

***
到此就结束啦
***

### 三、番外
我本地共享git仓库的文件，直接使用git查看记录，会有报错  
我查了一下大概得意思就是因为git为了安全  
进行git操作的时候会检查当前仓库账户是否在预期内  
因为我们是共享目录中，是别的用户拉取的仓库  
这里我们就跟着提示的信息，将这个账户的信息加入到Windows这边的配置中即可

``` bash
git config --global --add safe.directory "//192.168.234.128samba_share/game-kj"
```
