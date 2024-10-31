在Ubuntu 22.04中安装`stern`，你可以按照以下步骤进行：

1. **下载Stern的最新版本**：
   首先，你需要从GitHub上的Stern仓库获取最新的二进制文件。你可以访问[Stern Releases](https://github.com/stern/stern/releases)页面来找到适合你系统架构的最新版本。假设你使用的是64位的系统，你可以用以下命令来下载：

   ```bash
   wget https://github.com/stern/stern/releases/download/v1.31.0/stern_1.31.0_linux_amd64.tar.gz
   ```

   请确保将链接中的版本号替换为最新的版本号。

2. **解压缩文件**:
   首先，使用 `tar` 命令解压下载的 `.tar.gz` 文件。将 `stern_1.31.0_linux_amd64.tar.gz` 替换为你实际下载的文件名：

   ```bash
   tar -zxvf stern_1.31.0_linux_amd64.tar.gz
   ```

   这将解压缩 `stern` 二进制文件。

3. **添加执行权限**:
   给解压后的 `stern` 文件添加执行权限：

   ```bash
   chmod +x stern
   ```

4. **移动到系统路径中**:
   然后，将`stern`可执行文件移动到系统的路径中，例如 `/usr/local/bin`：

   ```bash
   sudo mv stern /usr/local/bin/
   ```

5. **验证安装**:
   最后，验证 `stern` 是否安装成功，可以执行以下命令：

   ```bash
   stern --version
   ```
