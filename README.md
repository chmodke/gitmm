# 使用说明

> 集成一些批量操作git仓库的工具

## 安装卸载方法

### 安装

编译后，复制`gitmm.exe`到`C:\Windows`目录下即可全局使用。

### 卸载

删除`C:\Windows\gitmm.exe`即可。

## 环境要求

git > 2.28.0

## 仓库访问方式说明

### ssh方式（推荐）

使用ssh方式访问仓库请配置ssh私钥到代码托管平台。

```shell
# 生成密钥对，执行ssh-keygen命令，一路回车
ssh-keygen
# 查看生成的公钥信息，这个公钥就是需要配置到平台上的内容
cat ~/.ssh/id_rsa.pub

# 也可以自定义公钥
# 自定义公钥的配置方式请参考ssh客户端配置在 ~/.ssh/config 文件中配置，示例如下
Host gitee.com
    HostName gitee.com
    IdentityFile ~/.ssh/kehao@kehaopc
    User git

Host *
    IdentitiesOnly yes
    AddressFamily inet
    Protocol 2
    Compression yes
    ServerAliveInterval 60
    ServerAliveCountMax 20
    HostkeyAlgorithms +ssh-rsa
    PubkeyAcceptedKeyTypes +ssh-rsa

```

配置ssh私钥到代码托管平台的方法各有差异，请查阅托管平台配置项。

### http/https方式

使用http/https方式访问仓库请配置认证信息持久化存储。

配置方法

```shell
git config --global credential.helper store
```

## 配置文件说明

配置文件 `repo.yaml` 用来配置主从仓库上游信息，示例内容如下：

```yaml
# main_group from
main_group: "git@gitee.com:chmodke"
# origin_group to
origin_group: "ssh://git@192.168.100.100:2222/chmodke"
# repos repository list
repos:
  - arpc
  - ftrans
  - gitmm
```

- main_group : 主仓库上游地址
- origin_group : 从仓库(fork)上游地址
- repos : 仓库名称列表

<font style="color:red">注意：`main_group` ：是主仓库上游地址，`origin_group` 是从仓库上游地址，一定不能配反了。</font>

## 全局选项

- -x: --debug，控制显示详细的执行细节，取值范围debug/info/warn/error或者d/i/w/e，默认值info

## 子命令说明

> 所有子命令都支持`-h`或`--help` 参数查看帮助。

### config

> 生成示例配置文件

用于在当前目录下生成示例配置文件，校验当前目录下的配置文件

#### 执行格式

```shell
Usage:
  gitmm config [flags]
  gitmm config [command]

Available Commands:
  generate    生成示例配置文件
  verify      校验配置文件

Flags:
  -h, --help   help for config

Global Flags:
  -x, --debug string   show more detail. (default "info")
# 示例，校验配置文件
gitmm config verify
```

#### 参数

- 无参数

### clone

> 批量克隆仓库

执行命令会读取当前目录下`repo.yaml`配置文件，遍历`repos`配置项，从`origin_group`克隆代码到当前目录下`work_dir`指定的文件夹中。

#### 执行格式

```shell
Usage:
  gitmm clone [flags]

Examples:
gitmm clone -w tmp -b master

Flags:
  -h, --help                 help for clone
  -b, --work_branch string   克隆代码的分支 (default "master")
  -w, --work_dir string      克隆代码的存放路径 (default "master")

Global Flags:
  -x, --debug string   show more detail. (default "info")

# 示例：下面将克隆仓库的master分支到本地tmp目录中
gitmm clone -w tmp -b master
```

#### 参数

- work_dir: 必填项，克隆代码的存放路径
- work_branch: 可选项，克隆代码的分支，缺省值`master`

### sync

> 批量同步主从仓库

执行命令会读取当前目录下`repo.yaml`配置文件，遍历`repos`配置项，从`main_group`强制同步全部内容到`origin_group`中，需要用户对`origin_group`有强制写权限（取消分支保护）。

<font style="color:red">注意：会强制以 `main_group` 中的内容覆盖 `origin_group` 中的内容。</font>

#### 执行格式

```shell
Usage:
  gitmm sync [flags]

Examples:
gitmm sync

Flags:
  -h, --help   help for sync

Global Flags:
  -x, --debug string   show more detail. (default "info")
```

#### 参数

- 无参数

### pull

> 批量拉取仓库

执行命令会遍历`work_dir`目录下中的git仓库，并执行分支拉取操作。

#### 执行格式

```shell
Usage:
  gitmm pull [flags]

Examples:
gitmm pull -w tmp

Flags:
  -f, --force             强制拉取
  -h, --help              help for pull
  -w, --work_dir string   本地代码的存放路径 (default ".")

Global Flags:
  -x, --debug string   show more detail. (default "info")
# 示例：下面拉取当前目录下tmp文件夹中所有仓库的最新代码
gitmm pull -w tmp
```

#### 参数

- work_dir: 可选项，仓库的存放路径，默认值当前目录
- f: 可选项，强制拉取，强制拉取时会回退本地所有修改到远程的最新记录

### create

> 批量创建分支

执行命令会遍历`work_dir`中的git仓库，并执行分支创建操作。

#### 执行格式

```shell
Usage:
  gitmm create [flags]

Examples:
git create -w tmp -b develop

Flags:
  -h, --help                help for create
  -b, --new_branch string   新分支名称 (default "master")
  -r, --refs string         新分支起点 (default "HEAD")
  -w, --work_dir string     本地代码的存放路径 (default ".")

Global Flags:
  -x, --debug string   show more detail. (default "info")
# 示例：下面将对当前目录下tmp文件夹中所有仓库基于当前节点创建develop分支
gitmm create -w tmp -b develop
```

#### 参数

- work_dir: 可选项，仓库的存放路径，默认值当前目录
- new_branch: 必填项，新分支名称
- start_point: 可选项，分支起始点，默认`HEAD`

### switch

> 批量切换分支

执行命令会遍历`work_dir`中的git仓库，并执行分支切换操作。

#### 执行格式

```shell
Usage:
  gitmm switch [flags]

Examples:
gitmm switch -w tmp -b develop

Flags:
  -b, --branch string     目标分支/tag/commit (default "master")
  -f, --force             强制切换
  -h, --help              help for switch
  -w, --work_dir string   本地代码的存放路径 (default ".")

Global Flags:
  -x, --debug string   show more detail. (default "info")
# 示例：下面将切换当前目录下tmp文件夹中所有仓库到develop分支
gitmm switch -w tmp -b develop
```

#### 参数

- work_dir: 可选项，仓库的存放路径，默认值当前目录
- branch: 必填项，目标分支名称
- f: 可选项，强制切换，强制切换时会回退本地所有修改

### remote

> 批量查看仓库远程信息

执行命令会遍历`work_dir`目录下中的git仓库，并查看仓库远程信息。

#### 执行格式

```shell
Usage:
  gitmm remote [flags]

Examples:
gitmm remote -w tmp

Flags:
  -h, --help              help for remote
  -w, --work_dir string   本地代码的存放路径 (default ".")

Global Flags:
  -x, --debug string   show more detail. (default "info")
# 示例：下面将查看当前目录下master文件夹中所有仓库的远程信息
gitmm remote -w tmp
```

#### 参数

- work_dir: 可选项，仓库的存放路径，默认值当前目录

### batch

> 批量执行提供的git命令

执行命令会遍历`work_dir`中的git仓库，并执行提供的git命令。

#### 执行格式

```shell
Usage:
  gitmm batch [flags]

Examples:
gitmm batch -w tmp 'log --oneline -n1'

Flags:
  -h, --help              help for batch
  -w, --work_dir string   本地代码的存放路径 (default ".")

Global Flags:
  -x, --debug string   show more detail. (default "info")
# 示例：下面将查看当前目录下tmp文件夹中所有仓库最新一条提交记录
gitmm batch -w tmp 'log --oneline -n1'
```

#### 参数

- work_dir: 可选项，仓库的存放路径，默认值当前目录
- command: 必填项，git命令
