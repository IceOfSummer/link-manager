# symbolic-link-manager

一个管理系统软连接的工具，通常用于快速切换 SDK 工具的版本。

## 安装

从 [Release](https://github.com/IceOfSummer/symbolic-link-manager/releases) 列表下载最新版本，然后放在任意目录中.

推荐添加如下环境变量(可选, 不添加不影响使用。**如果不提供，默认使用可执行文件所在的目录**):

- `SLINK_MANAGER_HOME`: 存放数据的目录

> [!TIP]
> 添加完成后同样推荐添加 `Path` 路径，方便后续使用。

## 使用用例

### 切换 Java 版本

在系统中有两个 `Java` 版本:

- `C:\Program Files\Java\jdk-17`
- `C:\Program Files\Java\jdk8u432-b06`

依次使用下列命令来管理这两个版本:

```shell
# 声明一个(软)链接，名称为 java
slm add link java

# 为链接打标签
slm add tag java 17 "C:\Program Files\Java\jdk-17"
slm add tag java 8 "C:\Program Files\Java\jdk8u432-b06"

# 切换版本
slm use java 17
```

此时将会在 `$SLINK_MANAGER_HOME/app/` 目录创建一个名称为 `java` 的软连接。之后你需要将 `JAVA_HOME` 的值设置为对应的软连接路径:

- windows: `%SLINK_MANAGER_HOME%\app\java`
- linux: `$SLINK_MANAGER_HOME/app/java`

如果你在环境变量中没有定义 `SLINK_MANAGER_HOME`，这里需要自己手动替换为响应的路径。

> [!NOTE]
> 这里还需确保你已经添加了 `%JAVA_HOME%/bin`(windows) / `$JAVA_HOME/bin`(linux) 到 `Path` 中。

**仅第一次设置时需要重新打开终端才会生效，之后切换版本时不需要再重新打开**。

设置完成后，查看 Java 版本:

![java-switch](/doc/java-switch.png)

#### 当切换 Java 版本后自带切换 Maven 版本

当切换到 Java17 后，自动将 Maven 切换到 `3.8.8`。当切换到 Java8 后，自动将 Maven 切换到 `3.6.3`。

Maven 目录:

- `D:\DevelopmentTool\apache-maven-3.6.3`
- `D:\DevelopmentTool\apache-maven-3.8.8`

管理这两个版本:

```shell
# 声明一个(软)链接，名称为 maven
slm add link maven

# 为链接打标签
slm add tag maven 3.6 "D:\DevelopmentTool\apache-maven-3.6.3"
slm add tag maven 3.8 "D:\DevelopmentTool\apache-maven-3.8.8"

# 和 Java 绑定(Java 的链接需要提前创建)
slm add bind java:17 maven:3.8
slm add bind java:8 maven:3.6

# 先使用一次，不然不会创建软连接
slm use maven 3.6
```

执行完成后，将环境变量 `MAVEN_HOME` 设置为 `%SLINK_MANAGER_HOME%\app\maven` (windows).

此时完成设置，切换 Java 版本将会跟着切换 Maven 版本：

![绑定](/doc/bind.png)


### 切换 NodeJs 版本

NodeJS 目录:

- `D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v22.11.0-win-x64`
- `D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v20.18.0-win-x64`

管理这两个版本:

```shell
slm add link node

slm add tag node 20 "D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v20.18.0-win-x64"
slm add tag node 22 "D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v22.11.0-win-x64"

slm use node 20
```

添加环境变量(Path): `D:\DevelopmentTool\symbolic-link-manager\app\node` 如果你设置了 `SLINK_MANAGER_HOME`，
也可以使用 `%SLINK_MANAGER_HOME%\app\node`.