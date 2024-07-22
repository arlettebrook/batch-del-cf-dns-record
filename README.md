# batch-del-cf-dns-record

> Batch delete cloudflare DNS records。批量删除cloudflare dns记录。



---



## 介绍

当我们将域名解析到*Cloudflare*时，不知道什么原因，系统可能会自动导入几百条不那么正确的解析*记录*，让人抓狂。

这些*记录*在界面中无法快速*删除*，也不支持跳过导入 ...

手动删除又过于麻烦。

网上看了一下。我们可以通过*Cloudflare* 的API功能来实现*批量删除*解析。因此写了一个Go版本的脚本。仅供参考。

地址是: https://github.com/arlettebrook/batch-del-cf-dns-record



---



## 如何使用

**帮助信息**：

```sh
$ ./batch-del-cf-dns-record.exe --help
Usage of Batch-del-cf-dns-record:
  -l, --log_level string   Log level (default "info")
  -a, --api_token string   Cloudflare API Token
  -z, --zone_id string     Cloudflare Zone ID
pflag: help requested
```



### 前提条件

- 登录cloudflare账号

- 进入要删除DNS记录域名的首页
- 滑动到底部右下角，有一个我们需要的第一个参数：区域ID(`zone_id`)
- 最后一个参数：`api_token`
  - 点击下面的[获取您的API令牌](https://dash.cloudflare.com/profile/api-tokens)，创建令牌
  - 创建一个编辑区域的DNS令牌。
  - 这个令牌就是`api_token`
  - 使用完成之后可以删除



### 源码运行

**前提条件**：需要Go环境

1. 克隆项目到本地

   ```sh
   git clone https://github.com/arlettebrook/batch-del-cf-dns-record.git
   ```

2. 进入项目，安装依赖，运行`main.go`并指定参数即可。

   ```sh
   cd batch-del-cf-dns-record
   
   go mod tidy
   
   go run main.go -a api_token -z zone_id
   ```



### 二进制文件运行[推荐]

**注意**：只提供了Windows的二进制文件，其他系统自行编译。

1. 前往[发布页面](https://github.com/arlettebrook/batch-del-cf-dns-record/releases/latest)下载，最新Windows版本

2. 解压

3. 在**Windows终端**中运行`batch-del-cf-dns-record.exe`并指定参数即可:

   ```sh
   batch-del-cf-dns-record.exe --a api_token -z zone_id
   ```

4. 如果是bash之类的终端，运行：

   ```
   ./batch-del-cf-dns-record.exe --a api_token -z zone_id
   ```

   



---



## 注意事项

- 如果是TLS握手超时，重新运行即可。
