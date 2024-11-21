要在Ubuntu 22.04中实现通过SSH让外部用户连接进来，您需要按照以下步骤进行操作：

1. **安装SSH Server**：
   - 打开终端。
   - 输入以下命令来安装OpenSSH Server。
     ```bash
     sudo apt update
     sudo apt install openssh-server
     ```
   - 安装完成后，SSH服务通常会自动启动。您可以使用以下命令来检查它的状态：
     ```bash
     sudo systemctl status ssh
     ```
     如果服务未运行，您可以使用以下命令来启动它：
     ```bash
     sudo systemctl start ssh
     ```

2. **配置SSH Server**（可选）：
   - SSH配置文件位于 `/etc/ssh/sshd_config`。
   - 您可以使用您喜欢的文本编辑器（如nano或vim）编辑这个文件来改变配置，例如修改默认端口号或禁用root登录等：
     ```bash
     sudo nano /etc/ssh/sshd_config
     ```
   - 例如，您可能想改变端口号，找到 `#Port 22`，去掉前面的 `#` 并将 `22` 更改为您所希望的端口号。
   - 完成更改后，您需要重启SSH服务使更改生效：
     ```bash
     sudo systemctl restart ssh
     ```

3. **防火墙设置**：
   - 如果您的Ubuntu系统上安装了UFW防火墙，您需要允许外部连接的SSH端口：
     ```bash
     sudo ufw allow ssh
     ```
     如果您更改了SSH端口，请确保允许那个特定的端口：
     ```bash
     sudo ufw allow <port_number>/tcp
     ```
   - 启用防火墙（如果尚未启用）：
     ```bash
     sudo ufw enable
     ```
   - 检查防火墙规则是否设置正确：
     ```bash
     sudo ufw status
     ```

如果没有什么特殊的，一般情况只需要做第1步安装即可
