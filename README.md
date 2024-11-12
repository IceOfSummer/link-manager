# symbolic-link-manager

> [!IMPORTANT]
> 目前处于最终测试阶段...

一个管理系统软连接的工具，通常用于快速切换 SDK 工具的版本。


## 安装

从 [Release](https://github.com/IceOfSummer/symbolic-link-manager/releases) 列表下载最新版本，然后放在任意目录中，
并将该目录添加到下面的环境变量：

- `SLINK_MANAGER_HOME`: 存放数据的目录

> [!TIP]
> 添加完成后还需要添加 `Path` 路径，值直接使用 `SLINK_MANAGER_HOME` 即可，以便可以随处使用。

## 使用用例

### 切换 Java 版本

在系统中有两个 `Java` 版本:

- `C:\Program Files\Java\jdk-17`
- `C:\Program Files\Java\jdk-8`

依次使用下列命令来管理这两个版本:

```shell
# 声明一个(软)链接，名称为 java
slm add link java

# 为链接打标签
slm add tag java 17 "C:\Program Files\Java\jdk-17"
slm add tag java 8 "C:\Program Files\Java\jdk-8"

# 切换版本
slm use java 17
```

此时将会在 `$SLINK_MANAGER_HOME/app/` 目录创建一个名称为 `java` 的软连接。之后你需要将 `JAVA_HOME` 的值设置为对应的软连接路径:

- windows: `%SLINK_MANAGER_HOME%\app\java`
- linux: `export JAVA_HOME=$SLINK_MANAGER_HOME/app/java`

> [!NOTE]
> 这里还需确保你已经添加了 `%JAVA_HOME%/bin` / `$JAVA_HOME/bin` 到 `Path` 中。 

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
slm add bind java:17 manven:3.8
slm add bind java:8 maven:3.6

# 先使用一次，不然不会创建软连接
slm use maven 3.6
```

执行完成后，将环境变量 `MAVEN_HOME` 设置为 `%SLINK_MANAGER_HOME%\app\maven` (windows).

此时完成设置，切换 Java 版本将会跟着切换 Maven 版本：





%SLINK_MANAGER_HOME%\app\maven