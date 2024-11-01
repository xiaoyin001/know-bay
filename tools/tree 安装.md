在 Ubuntu 22.04 (或其它版本的 Ubuntu) 上安装 `tree` 命令是十分简单的。你只需打开你的终端（通常可以通过按 `Ctrl + Alt + T` 快捷键来打开）并输入以下命令：

```bash
sudo apt update
sudo apt install tree
```

这里的每一步解释如下：

1. `sudo apt update`: 更新你的包管理器的包索引，确保你安装的是最新的软件包版本。`sudo` 是（superuser do 的缩写）使命令以超级用户（或叫 root 用户）的权限来执行，这通常是必要的，因为安装软件包需要对系统的更深层次修改。

2. `sudo apt install tree`: 安装 `tree` 软件包。这条命令会下载并安装 `tree` 命令到你的系统中。

完成后，你就可以通过在终端中输入 `tree` 来使用此命令了。`tree` 命令可以帮助你递归地列出目录结构并以树状图形式展示。如果你想了解更多关于 `tree` 命令的信息，可以通过 `man tree` 命令来访问它的手册页。