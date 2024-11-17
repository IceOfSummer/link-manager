# symbolic-link-manager

[中文文档](/README_zh.md)

A tool to manage system symbolic link, usually used to switch SDK version.

## Install

Download from [Release](https://github.com/IceOfSummer/symbolic-link-manager/releases), and put the executable file to any empty folder.

You can add the environment variable below for better use():

- `SLINK_MANAGER_HOME`: The folder to save data and executable file. **The default value is the folder where the executable exist.**
- Add the executable file to `Path`

These environment variables are **optional**, you can also continue to use if you don't provide.

## Example

### Switch Java version

We have two `Java` here:

- `C:\Program Files\Java\jdk-17`
- `C:\Program Files\Java\jdk8u432-b06`

Manage these version by the command below:

```shell
# Declare a (symbolic) link，which name is `java`
slm add link java

# Add a tag for the link
slm add tag java 17 "C:\Program Files\Java\jdk-17"
slm add tag java 8 "C:\Program Files\Java\jdk8u432-b06"

# Switch the tag
slm use java 17
```

After all, a symbolic link named `java` will be created, it will locate at `$SLINK_MANAGER_HOME/app/`. Now you should set your
`JAVA_HOME` to this directory: 

- windows: `%SLINK_MANAGER_HOME%\app\java`
- linux: `$SLINK_MANAGER_HOME/app/java`

Replace the `SLINK_MANAGER_HOME` if you don't set it in environment variable.

> [!NOTE]
> Make sure you also add `%JAVA_HOME%/bin`(windows) / `$JAVA_HOME/bin`(linux) to your `Path`。 

**Then reopen the terminal. You only need to reopen the terminal the first time you set up the configuration for it to take effect.**。

After all, check you java version:

![java-switch](/doc/java-switch.png)

#### Switch the Maven version when switch the version of Java

When we switched to Java 17, automatically switch the Maven version to `3.8.8`. And when using Java 8，switch Maven to `3.6.3`。

Maven directories:

- `D:\DevelopmentTool\apache-maven-3.6.3`
- `D:\DevelopmentTool\apache-maven-3.8.8`

Manage these two versions:

```shell
# Declare a (symbolic) link，which name is `maven`
slm add link maven

# Add a tag for the link
slm add tag maven 3.6 "D:\DevelopmentTool\apache-maven-3.6.3"
slm add tag maven 3.8 "D:\DevelopmentTool\apache-maven-3.8.8"

# Bind to Java (You should create it before you use.)
slm add bind java:17 maven:3.8
slm add bind java:8 maven:3.6

# Switch to maven 3.6
slm use maven 3.6
```

After this, set the environment variable `MAVEN_HOME` to `%SLINK_MANAGER_HOME%\app\maven` (windows).

Now, when you switch the Java version will switch bound Maven version too:

![绑定](/doc/bind.png)


### Switch Node.js version

NodeJS directories:

- `D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v22.11.0-win-x64`
- `D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v20.18.0-win-x64`

Manage these two versions:

```shell
slm add link node

slm add tag node 20 "D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v20.18.0-win-x64"
slm add tag node 22 "D:\DevelopmentTool\symbolic-link-manager\sdk\nodejs\node-v22.11.0-win-x64"

slm use node 20
```

Add the `Path` variable: `D:\DevelopmentTool\symbolic-link-manager\app\node`. If you have set the `SLINK_MANAGER_HOME`,
you can use `%SLINK_MANAGER_HOME%\app\node` instead.