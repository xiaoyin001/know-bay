## 解析 Proto 文件

### 1. 背景
在开发的过程中会用到pb，有时候很多代码都是重复的填入“有规律”的内容，
而这些内容都是来源于 `.proto` 文件中，所以需要一个工具解析 `.proto` 文件，
并且将解析后的内容填入预设的位置，在使用的时候执只需调用生成代码中的某个方法即可。

### 2. 关于 .proto 内容解析
内容的解析具体可以看 `proto\doc\proto结构.go`，目前只写了我用到的部分，
以后如果还有机会会继续补充。

### 3. 使用
可以查看[proto](https://github.com/emicklei/proto)作者的使用方式，
或者也可以直接看看我的自己的示例 `proto\demo` 这个路径下面也会有一份比较简单的示例。

