# 使用说明

> 集成一些批量操作git仓库的工具

## 安装

复制`gitmm.exe`到`C:\Windows`目录下即可全局使用。

## 卸载

删除`C:\Windows\gitmm.exe`即可。

## 环境要求

git > 2.28.0

## 配置文件说明

配置文件 `repo.yaml` 用来配置主从仓库上游信息，示例内容如下：

```yaml
remote:
  # upstream from
  upstream: 'git@gitee.com:chmodke'
  # origin to
  origin: 'ssh://git@192.168.100.100:2222/chmodke'
# repos repository list
repos:
  - arpc
  - ftrans
  - gitmm
```

- upstream : 主仓库上游地址
- origin : 从仓库(fork)上游地址
- repos : 仓库名称列表

<font style="color:red">注意：`upstream` ：是主仓库上游地址，`origin` 是从仓库上游地址，一定不能配反了。</font>

## 命令说明

> 所有子命令都支持`-h`或`--help` 参数查看帮助。

```shell
PS C:\Users\kehao> gitmm --help
git多仓库管理工具，通过简单的配置对仓库进行批量管理。

Usage:
  gitmm [command]

Available Commands:
  batch       批量执行提供的git命令
  branch      批量分支操作
  clone       批量克隆仓库
  completion  Generate the autocompletion script for the specified shell
  config      生成示例配置文件，校验配置文件
  fetch       批量拉取仓库
  help        Help about any command
  list        展示工作路径下的Git仓库信息
  pull        批量拉取仓库
  remote      批量远程地址管理
  sync        批量同步主从仓库
  version     Show tool version

Flags:
  -h, --help      help for gitmm
  -v, --version   show tool version.

Use "gitmm [command] --help" for more information about a command.
```

### 命令列表

```shell
gitmm clone 
gitmm sync 
gitmm pull 
gitmm batch 
gitmm branch 
gitmm branch create 
gitmm branch delete 
gitmm branch list 
gitmm branch rename 
gitmm branch switch 
gitmm fetch 
gitmm list 
gitmm remote 
gitmm remote add 
gitmm remote remove 
gitmm remote show 
gitmm config 
gitmm config generate 
gitmm config verify 
```

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
