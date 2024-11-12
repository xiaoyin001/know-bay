在Ubuntu 22.04 (LTS) 中安装 `kubectl` 可以通过几种方法完成。以下是一些安装 `kubectl` 的推荐方法。

### 方法 1: 用 `apt` 包管理器安装

1. 更新包列表：
   ```bash
   sudo apt-get update
   ```

2. 安装包管理工具的依赖：
   ```bash
   sudo apt-get install -y apt-transport-https ca-certificates curl
   ```

3. Google Cloud 提供的公共软件包签名密钥添加到你的系统中：
   ```bash
   sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
   ```

4. Kubernetes 的 `apt` 仓库添加到你的系统的仓库列表中：
   ```bash
   echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
   ```

5. 再次更新包列表，以确保可以访问新添加的 Kubernetes 仓库：
   ```bash
   sudo apt-get update
   ```

6. 安装 `kubectl`：
   ```bash
   sudo apt-get install -y kubectl
   ```

### 方法 2: 使用 `snap` 包管理器安装

如果你更喜欢使用 `snap` 来管理你的应用，你也可以用它来安装 `kubectl`：

1. 通过 Snap 安装：
   ```bash
   sudo snap install kubectl --classic
   ```

`snap` 会自动更新 `kubectl` 到最新版本。

### 方法 3: 手动下载并安装 `kubectl` 的二进制文件（亲测可以）

1. 下载最新版本的 `kubectl`:
   ```bash
   curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
   ```

2. 使二进制文件可执行：
   ```bash
   chmod +x ./kubectl
   ```

3. 将二进制文件移动到 PATH 中的某个目录：
   ```bash
   sudo mv ./kubectl /usr/local/bin/kubectl
   ```

4. 测试以确保版本为最新：
   ```bash
   kubectl version --client
   ```

确保选择与你管理员身份或个人使用习惯最符合的方法。如果你在一个企业环境中，可能需要遵循特定的安装指南或限制。