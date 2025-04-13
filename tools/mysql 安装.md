# 在 Ubuntu 22.04 上安装 MySQL

## 1. 更新系统包
确保你的系统包是最新的：
```bash
sudo apt update
sudo apt upgrade -y
```

## 2. 安装 MySQL
运行以下命令安装 MySQL：
```bash
sudo apt install mysql-server -y
```

安装完成后，MySQL 服务会自动启动。

## 3. 配置 MySQL
安装完成后，运行 MySQL 的安全配置脚本，以提高安全性：
```bash
sudo mysql_secure_installation
```

按照提示进行配置：
1. **设置 MySQL root 用户的密码**：选择 `Y`，然后输入并确认新密码。
2. **移除匿名用户**：选择 `Y`。
3. **禁止 root 用户远程登录**：选择 `Y`（如果你不需要远程登录 MySQL，建议选择 `Y`）。
4. **删除测试数据库**：选择 `Y`。
5. **重新加载权限表**：选择 `Y`。

## 4. 验证 MySQL 是否安装成功
检查 MySQL 服务状态：
```bash
sudo systemctl status mysql
```
如果服务正在运行，你会看到类似 `active (running)` 的状态。

## 5. 登录 MySQL
使用以下命令登录 MySQL：
```bash
sudo mysql -u root -p
```
输入你在 `mysql_secure_installation` 中设置的密码。

## 6. 创建新用户（可选）
为了安全起见，建议不要直接使用 root 用户操作数据库，可以创建一个新用户：
```sql
CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'newuser'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

## 7. 允许远程访问（可选）
如果你需要从远程连接到 MySQL 数据库，可以修改 MySQL 配置文件：
```bash
sudo nano /etc/mysql/mysql.conf.d/mysqld.cnf
```

找到 `bind-address` 行，将其注释掉或改为 `0.0.0.0`：
```ini
# bind-address = 127.0.0.1
bind-address = 0.0.0.0
```

然后重启 MySQL 服务：
```bash
sudo systemctl restart mysql
```

最后，确保防火墙允许 MySQL 的端口（默认是 `3306`）：
```bash
sudo ufw allow 3306
```

## 8. 安装 MySQL 客户端（可选）
如果你需要在其他机器上连接到 MySQL，可以安装 MySQL 客户端：
```bash
sudo apt install mysql-client -y
```

## 总结
通过以上步骤，你已经成功在 Ubuntu 22.04 上安装并配置了 MySQL。如果需要进一步操作，可以参考 MySQL 官方文档或社区资源。
