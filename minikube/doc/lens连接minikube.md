要在 Windows 10 上的 Lens 中查看在 Ubuntu 22.04 虚拟机上部署的 Minikube 的相关信息，需要执行以下几个步骤：

1. **确保 Minikube 在 Ubuntu 虚拟机上正常运行**：
   在 Ubuntu 虚拟机中打开一个终端，运行下面的命令来检查 Minikube 状态：
   ```bash
   minikube status
   ```

2. **获取 kubeconfig 文件**：
   Lens 需要 kubeconfig 文件来连接 Kubernetes 集群。此文件通常位于 Ubuntu 虚拟机中的 `$HOME/.kube/config` 路径下。你可以通过运行以下命令查找此文件：
   ```bash
   cat $HOME/.kube/config
   ```
   你需要将这个配置文件或其内容传输到你的 Windows 10 系统中。同时还需要看配置中有一些证书，也是需要一同弄到 Windows 10 上的。

   这是我的在lens中设置的config文件内容，可以参考一下，我们直接使用minikube集群中的配置可能还有一些可选的参数，拿到原始的配置是需要修改部分参数后才可以用到lens中的，我这里就改了证书所在路径
   ```sh
    # 这是一个Kubernetes配置文件，用来定义集群、用户和上下文信息。
    apiVersion: v1  # 配置文件的API版本。

    clusters:  # 集群配置条目的列表。
    - cluster:
        # certificate-authority: D:\MyFile\minikube\ca.crt  # 证书授权文件的路径（在这里被注释掉了，因为跳过了TLS验证）。
        insecure-skip-tls-verify: true  # 表示不应验证服务器的证书（不推荐用于生产环境）。
        server: https://192.168.116.136:8443  # Kubernetes API服务器的URL地址（我这里填的是我虚拟机的IP，做了端口转发）。
    name: minikube  # 集群的名称。

    contexts:  # 上下文配置条目的列表。
    - context:
        cluster: minikube  # 该上下文对应的集群名称。
        namespace: xiaoyin  # 该上下文使用的默认命名空间。
        user: minikube  # 该上下文使用的用户名称。
    name: minikube  # 上下文的名称。

    current-context: minikube  # 当前正在使用的上下文名称。

    kind: Config  # API对象的类型（在这里是kubectl的配置文件）。

    preferences: {}  # 用于保存命令行标志的全局默认值。

    users:  # 用户配置条目的列表。
    - name: minikube  # 用户条目的名称。
    user:
        client-certificate: D:\MyFile\minikube\client.crt  # 用户认证的客户端证书文件路径。
        client-key: D:\MyFile\minikube\client.key  # 用户认证的客户端密钥文件路径。
   ```
   上面证书如果不想每次都复制文件和配置一起的话，也可以将证书字段后面加上 `-data`例如`client-key-data: <base64的内容>`，然后在将文件进行base64编码，将编码后的内容填入即可，跟使用文件的效果相同
   ```sh
   # Linux下将执行文件进行bases64编码，这里的-w 0 表示不用将生成的编码换行输出
   base64 -w 0 /home/xiaoyin01/.minikube/ca.crt
   ```

   关于ca证书我本地是直接屏蔽掉了，minikube安装启动啥的会有一个证书，但是这个ca正式只能本地用所以我想要在lens上用的话就需要有公共的ca证书或者自己认证签发的

3. **配置网络**：
   需要确保 Windows 10 能够访问 Ubuntu 虚拟机中 Minikube 创建的 Kubernetes API 服务器。通常，VM 网络设置为 NAT，因此可能需要设置端口转发规则。

   - 添加一条新的端口转发规则，主机端口设为 `6443`，客户端端口设置为 Minikube 监听的端口，也通常是 `6443`。
   ```sh
   # 设置端口转发规则，这种是一次性的，重启后就需要重新设置
   sudo iptables -t nat -A PREROUTING -p tcp --dport 8443 -j DNAT --to-destination 192.168.49.2:8443

   # 查看端口转发规则
   sudo iptables -t nat -L PREROUTING -v -n

   # 如果需要删除规则，可以使用以下命令
   sudo iptables -t nat -D PREROUTING -p tcp --dport 8443 -j DNAT --to-destination 192.168.49.2:8443
   ```

4. **使用 Lens**：
   在 Windows 10 上打开 Lens 应用程序。
   - 如果是第一次使用 Lens，它将提示导入 kubeconfig 文件。
   - 如果已经使用过 Lens，可以通过文件菜单选择 "Add Cluster"，然后导入您之前从 Ubuntu 虚拟机中传输到 Windows 10 的 kubeconfig 文件。

5. **连接到集群**：
   导入配置文件后，Lens 应该能够显示集群信息。如果一切配置正确，你应该能够看到你的 Minikube 集群的节点、工作负载等信息。

请注意，细节可能会有所不同，具体取决于你的虚拟机配置和网络设置。