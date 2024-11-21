需要你已经安装了 Docker，接下来将只关注安装 Docker Compose。从 Docker Compose v2 开始，它作为 Docker CLI 的一个插件存在。以下是在 Ubuntu 22.04 上安装 Docker Compose 插件的步骤：

1. **下载 Docker Compose**

    首先，你需要从 GitHub 上获取 Docker Compose 的最新版本。你可以访问 [Docker Compose 的 GitHub Release 页面](https://github.com/docker/compose/releases) 来查找最新的版本号，或者运行以下命令来自动检索最新的稳定版本编号：

    ```sh
    COMPOSE_VERSION=$(curl --silent "https://api.github.com/repos/docker/compose/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")')
    ```

    解释：
    - `curl --silent`：`curl` 是一个工具，用于传输数据。在这里，`--silent` 或 `-s` 选项使 `curl` 不会显示进度条和错误消息。
    - `"https://api.github.com/repos/docker/compose/releases/latest"`：这是 GitHub API 的 URL，用于获取 Docker Compose 最新版本的信息。
    - `grep -Po`：`grep` 是用来搜索文本的工具。`-P` 选项允许使用 Perl 兼容的正则表达式，而 `-o` 选项仅输出匹配的部分。
    - `'"tag_name": "\K.*?(?=")'`：这是一个正则表达式，用来匹配最新版本号。

2. **安装 Docker Compose**

    使用上一步获取的版本号来下载并安装 Docker Compose：

    ```sh
    sudo curl -L "https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    ```

    解释：
    - `sudo`：以超级用户权限运行命令。
    - `curl -L`：`curl` 用于获取文件，并且 `-L` 选项会让 curl 跟随重定向。
    - `"https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)"`：这是 Docker Compose 下载链接，`${COMPOSE_VERSION}` 是上一步获取的最新版本号，`$(uname -s)` 和 `$(uname -m)` 分别用 shell 命令 `uname` 获取操作系统名称和机器硬件名称。
    - `-o /usr/local/bin/docker-compose`：指定下载的文件保存的位置。

3. **给 Docker Compose 设置执行权限**

    下载完成后，需要给 `docker-compose` 文件设置可执行权限：

    ```sh
    sudo chmod +x /usr/local/bin/docker-compose
    ```

    解释：
    - `sudo chmod +x /usr/local/bin/docker-compose`：`chmod +x` 给文件添加执行（`x`）权限。这使得执行 `docker-compose` 命令时操作系统可执行该文件。

4. **测试 Docker Compose 是否安装成功**

    安装完成后，你可以通过检查其版本来验证安装是否成功：

    ```sh
    docker-compose --version
    ```

    这条命令应当返回 Docker Compose 的版本信息，表明它已经正确安装在 `/usr/local/bin/docker-compose` 路径下。

按照上面的步骤，你应该能够成功地在你的 Ubuntu 22.04 系统上安装 Docker Compose。如果遇到任何问题，你可能需要检查 Docker Compose 的文档或 GitHub Release 页面，以确认是否有任何安装步骤的变更。