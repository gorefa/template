# 网关控制系统(YK-CGI)

目前我的所有服务都部署在局域网内，包括 WEB, 智能家居等，其中只是一小部分有选择的暴露在外网，并且这部分暴露的，完全依靠的手工配置，或者由一些小的脚本控制。

现在急需，统一的入口，控制所有的服务。

目前初步考虑使用 Golang 作为主力开发语言实现(主要是性能问题)，由于后续需要把所有服务全部整合到该项目上来，所以 `前期协议设计`，构思可能会和开发用的时长一样多，甚至更多。(扩展, 稳定, 做第一考虑; 性能放在其次)

同样采用 TDD, 并且使用&同时完善 [Golang工具箱](https://gitlab.commonheart-yk.com/ch-yk/w26/tree/master/%E5%B7%A5%E5%85%B7/Golang%E5%B7%A5%E5%85%B7%E7%AE%B1/golang-tools)。

> 本项目大概会在 10 月中旬左右完成，历时一个多月

* 2018.8.29 有好多好的想法，发现这些想法的功能实现最好都要借助这个系统  
  * --- 这个系统的需求处于 `Urgent`.
  * 先赶工一个原型  --- golang all the way.
  * 使用 `api.commonheart-yk.com`


**TODO**: before 2018.10.15 重点：统一调用格式，必须统一支持 `https`，最好考虑好扩展性

* 发现程序有点大了，所以 web 响应处理这一块，选用一个 web 框架吧 (Gin轻量级)
* 先考虑 http 自建 server，后续考虑在 nginx 上部署 https

## 1.概述

> 功能：请参考下面 `[使用者视角]` 部分

### 1.1 设计选型

> 谈一些技术选型，以及基于现实原因的退让
>
> * 设计思路一般都是`二次设计` + `从两头往中间掐`


API网关系统，早有人设计过了，也很成熟。

但是后续我个人肯定会根据个人需求，进行扩展，这个时候如果用别人的系统，弊端明显。



> 下面只是设计时考虑的，实际实现时，考虑的更多
>
> 请参考下面的 **相关技术**



先说部署:

* 开发环境是 `macOS`，但是最终部署是 Ubuntu Server
* 最终 部署: nginx , docker + ubuntu



- API 选型:
  - rpc: gRPC, Thrift, Protobuffer
    - rpc 一般是tcp层，性能会更好
    - 各种框架都有一套自己的 dsl 语言用于描述序列化(二进制数据)，可读性不好
    - 一般用于微服务内部，系统模块之间 — 而我的多半是对外部系统，外部client提供服务
  - restful: REST + JSON   
    - 后端必定是一个 crud 的数据库，且我的需求主要就是 crud
    - 主要方便开发过程调试 (json 可读性好)，API清晰，简单不复杂，轻量
    - 无状态，扩展也好 —— 关键是后续很多 web 请求需要扩展
    - 缺点：很多操作不一定能约定或者对应URI资源，需要提前约定好



提供几套API：

* 如果有时间应该提供两套格式的API，分别给外部和内部(自己的子系统)
* 没有时间，在同一套API里，通过 `切面` 加入内部还是外部调用判断逻辑，权限控制逻辑





API的操作:

* 用代码内部 mapping
* 用配置文件 mapping  --- 最佳
  * 也就是 API命令应该先加载配置文件，之后根据配置清单做后续操作
  * 可以考虑将这种方式写成一个简单 `框架`



配置文件

* 初始化 API 操作的 mapping
* 初始化日志系统
* 初始化数据库系统(实例)
* 初始化 HTTP(s)
  * 设置 Header
  * 注册路由
  * 注册中间件



服务器引擎:

* 处理响应，任何一个响应，都要在本地记录事务处理过程

* 客户端能够检查服务器是否存活 (专门有一个协程处理)

* 处理响应

  * 如果程序表，那么标准库的 `net` 包足够，如果路由规则多，或者后期需要扩展的多，则选用轻量级 web 框架


web框架:

* beego
* [gin](https://gin-gonic.github.io/gin/)   —— 最佳 (**要的就是轻量级**，因为此系统 web 部分不是重点，只是模块)
* Martini






### 1.2 相关技术

> 主要列举项目用到的知识点



Golang语言(包括 Gin 框架)，小型系统的设计能力，后端代码实现能力(比如查看 api status就用到了心跳监测)。

* golang 语言方面的知识
* golang 的依赖管理，最不济需要 `go get` 基本，本项目采用 `dep`
  * 注意墙
* 基本的环境配置，代码结构，配置，日志，数据库，错误处理等 —— 
* 前后端 token 问题: 登录场景， 任何 API 调用(同时辅助 API身份验证)权限认证
  - 权鉴用 `JWT` 方式
    - 永久 token — 不能 `明文传输`
    - 需要登录才能产生的有 validation duration 的 token，比如 `-XPOST -H "Authorization: token值"`
  - token 还有一个，**防止重复提交**
    - 每次 request 和reponse 都加入 32 uuid: `X-Request-Id`
* API权限: 采用 token 机制(见上)
* 查询优化: 代码端并发处理结果集, 数据库端设计表的时候就要加入分页项
* 自定义业务处理的加载时机 (包括构建 model 和解析model)
  - 包括加入 `切面操作`: 请求的日志记录, 插入某个 header在response中，API权限鉴定
  - 框架的中间件，或者自己写个中间层来处理，对于注册的函数进行调用(观察者模式)
  - 注意注册路由的时机和注册中间件的时机，先后顺序
* API版本问题 --- uri 路径解决
* API 请求记录问题
  - HTTP Request body 一旦被读取就会自动清空，所以记录API时，还需要重新赋值
  - 截获 HTTP 响应的话，则需要重定向到指定的 IO 流 — 一般只记录请求即可 (响应有后端逻辑)

* Nginx , 反向代理部署应用 
  * 最终还是算了, 直接部署更加方便 (关键是没有多台机器，不需要负载均衡)
* Docker 自定义镜像
* 容器网络知识 (容器内部访问其他容器的网络)
* 分页








### 1.3 思维脑图

> 暂定的，改项目的 road map

![16-46-02-golang网关系统.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/08/16-46-02-golang网关系统.jpg)





## 2. 使用者视角


所有的子系统，内部服务，包括是否要暴露在外部，都由该网关系统全部总揽。

**也就是说，该系统是提供给其他子系统，作为模块组合到其他系统** 一般不给普通用户使用。



后续所有的系统，代码，都从 `此系统作为入口`，提供服务。

(比如要给服务器上传文件，要从个服务器的某个数据库拿数据等等))




## 3. 开发者视角

> 下面记录了我的开发过程步骤 (step) ，可以从 git log 提交查看一些信息，这里详细记载

* 先解决 Golang 依赖管理问题， vgo 或者 dep
  * vgo 还是太新，暂时用 dep 开发吧
* API 的设计，规范

等等，这里不列举了，可以参考下面的 `version` 号的划分。


### 3.1 Version 0

**原型系统**: 用一个 demo 功能，让系统出现初步原型，方便后面调整设计方向，评估系统。



#### 3.1.1 Version 0.1 只管设计

 > 不管实现

确定了使用 RESTFUL, 而不是 RPC。(虽然 RPC 性能好一些) 

初步约定:

| 方法   | URI                 | 说明                    |
| ------ | ------------------- | ----------------------- |
| GET    | /services           | 获取 service 列表       |
| GET    | /services/service_a | 获取 service_a 的信息   |
| POST   | /services           | 创建一个新的 service    |
| PUT    | /services/service_a | 更新 service_a 的信息   |
| DELETE | /services/service_a | 删除 servcie_a 相关信息 |
| 待定   | N/A                 | N/A                     |

(`get /` 对应 `index.html`)

内外部提供同一套 API， 但是内部控制调用权限，以及内部或者外部的判断。



为了以后扩展方便，还是决定用一下外部框架 (原计划是零依赖的)。

* 虽然网络方面，标准库已经做了的很好了
* 虽然框架更多的是组合工具，语法糖
* 但是，只要能提高效率，糖未尝不好呀？



初步设计，不能太详细。



#### 3.1.2 Version 0.2 基本骨架



整体的思路有，详细的细节则比较乱，先借助 `.gitkeep` 填充目录，整理相关思路:

(采用 dep，目前不把相关依赖放置在项目内 —— `vendor` 是不算做我的项目代码的)

![09-33-30-项目骨架.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/08/09-33-30-项目骨架.jpg)

dep 下载依赖的问题， `gin` 这个包比较特殊:

```bash
## 问题:
## failed to list versions for https://github.com/gin-gonic/gin: fatal: unable to access 'https://github.com/gin-gonic/gin/': The requested URL returned error: 503
: exit status 128


##Latest git does not follow redirects by default, breaking gopkg.in
git config --global http.https://gopkg.in.followRedirects true
```



至于 :

```bash
unable to fetch raw metadata: failed HTTP request to 
URL "http://golang.org/x/net/html?go-get=1": 
```

需要用 `http` 而不是 `socks`。(懂得自然懂)



#### 3.1.3 Version0.3 跑起来

思路渐渐清晰了，先跑起来; 基于 `Gin` 的 HTTP server。

* 该部分代码实现一个完整的服务器状态查询，包括cpu, meminfo, 



| 方法 | URI                   | 说明                 | 对应的handler   |
| ---- | --------------------- | -------------------- | --------------- |
| GET  | /xserver_status/alive | ping服务器看是否能通 | 包名.AliveCheck |
| GET  | /xserver_status/disk  | 查看磁盘信息         | 包名.DiskCheck  |
| GET  | /xserver_status/cpu   | 查看cpu信息          | 包名.CPUCheck   |
| GET  | /xserver_status/mem   | 查看内存信息         | 包名.MEMCheck   |

其中包名就是请求 API 的分组名，此处是 `xserver_status`。



跑起来，测试，提交git。(先将就一下)

![13-54-13-api_gateway_v0.3_1.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/08/13-54-13-api_gateway_v0.3_1.jpg)

![13-54-27-api_gateway_v0.3_2.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/08/13-54-27-api_gateway_v0.3_2.jpg)

![13-54-52-api_gateway_v0.3_3.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/08/13-54-52-api_gateway_v0.3_3.jpg)



#### 3.1.4 Version0.4 加载配置

上一个小版本，等于说是借助 Gin 框架，把 Restful API 跑起来了。

但是具体一个日常用来不来基本不用操心的应用，还有很事情没有做。

* 最基本的 api 请求记录吧
* 日志记录
* 可扩展配置文件方案处理
  * 最起码端口可以重新配置吧
  * IP, Host地址可以配置吧
  * Ping测试几次放弃可以配置吧
* 权限检查



本小版本主要解决, 可扩展文件配置方案，即这个配置文件的问题。

主要处理:

* 命令行传入的 flag 参数
* 环境变量
* 配置文件



配置格式采用 `yaml`，然后采用 `flg` 和 `viper` 包.

(后扩后面数据库的链接配置，切换数据库等)





#### 3.1.5 Version 0.5 加载日志配置

利用 github 上已经有的工具箱，最简单的处理了日志存储，日志记录的问题。

配置文件如下:

```yaml
log:                         # 日志配置, 转储用的是 Linux 的 rotate 系统(定时检查)
  writers: file,stdout
  logger_level: DEBUG        # DEBUG, INFO, WARN, ERROR, FATAL
  logger_file: run.log
  log_format_text: False     # True 会输出 Json 格式，False则是一般 text 格式
  rollingPolicy: size       # daily 根据天进行转储，size则是根据大小
  log_rotate_date: 1         # rotate 利用的
  log_rotate_size: 10        # 10 M 大小
  log_backup_count: 100      # 日志备份个数
```





#### 3.1.6 Version0.6 添加数据库支持

如果每增加一个应用，都要去写数据库读写逻辑，这就非常烦人了。

以后的所有应用都基于此网关，路由到内部路由，其次所有数据库读写都要尝试经过这里的总路由处理。

(先确保某个数据库有可访问权限，比如规定只有内网才能访问的数据库 --- 最好新增一个用户)

* 本版本增加 api 系统，对于多个数据库的读写支持
  * 至少支持 mariadb/mysql --- 利用 gorm 框架能拿到数据库读写对象即可
  * 下一步扩展支持 mongodb



创建一个新的数据库以及表格: (后续的子系统，都将读写该数据库)

```sql
/* MariaDB 10.3.8, 仅内部网络特定账户可以访问 */
/* 创建一个数据库 */
DROP DATABASE IF  EXISTS `track_all`;
CREATE DATABASE `track_all` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `track_all`;


/*表的前缀说明: 
 * ta_xxx 数据库本身相关的
 * api_xxx API系统相关的
 * track_xxx 成就系统相关的
 * daily_xxx 习惯养成系统相关的
 */
 
 /*建立一张表格: api_services，查看该表为哪些系统服务*/
DROP TABLE IF EXISTS `api_services`;
CREATE TABLE `api_services` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `serviceName` varchar(255) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `serviceName` (`serviceName`),
  KEY `idx_api_services_deletedAt` (`deletedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


/*插入查看系统信息的服务(模块)*/
INSERT INTO `api_services` VALUES (0,'XServer_Status','2018-08-31 10:34:12','2018-09-02 13:25:33',NULL);
```

(表信息，后续再完善)



**新建用户** 利用 `phpMyAdmin` 吧(也可以用代码授权)，仅对上述数据库有读写权限。

```bash
GRANT ALL PRIVILEGES ON `track\_all`.* TO 'xxxx'@'192.168.10.%' WITH GRANT OPTION;
```



配置启动信息:

```bash
mariadb1:
  name: track_all
  addr: 114.114.114.114:3306
  username: xxxx
  password: yyyy
docker_db:
  name: track_all
  addr: 114.114.114.114:3306
  username: xxxx
  password: yyyy
mongodb1:
  name: track_all
```

条件有限，我这里最多弄两个数据库，读，写数据库。



#### 3.17 Version 0.7 定义错误信息

内部系统的话，非必要; 有的话，更加规范。主要就是自定义:

* `err.Err` 类型 (面型对象的方式)
* 填充 code 和 message
  * code 给开发者 --- 开发相关，比如 HTTP 相关
  * message 给用户 --- 业务相关



简单理解: 自定义异常类型；但返回的一般是 json 格式

```json
{
  "code": 10002,
  "message": "Error occurred while binding the request body to the struct."
}
```

后面响应处理要用到，先测试一下:

不传入参数:

![15-41-27-api_gateway_v0.7_1.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/09/15-41-27-api_gateway_v0.7_1.jpg)

只传入单个参数:

![15-43-52-api_gateway_v0.7_2.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/09/15-43-52-api_gateway_v0.7_2.jpg)



只传递后一个参数:

![15-48-55-api_gateway_v0.7_3.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/09/15-48-55-api_gateway_v0.7_3.jpg)

![15-50-45-api_gateway_v0.7_4.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/09/15-50-45-api_gateway_v0.7_4.jpg)



后面的所有功能和 Web 页面开发类似，读取 HTTP消息(头 + Body)，handler 处理(读参数: param, query, parse, 校验参数，token检查，业务逻辑处理，回写 HTTP Response)， ORM plain对象的填充与生成，数据业务执行(通常还要注意一下查询分页的情况，代码以及数据库两方面的查询优化)，Service业务处理(通常是被 Handler 调用)



返回的消息体，格式大概如下:

```json
{
  "code": 0,
  "message": "OK",
  "data": null
}
```



一般棘手的问题:

* 前后端 token 问题: 登录场景， 任何 API 调用(同时辅助 API身份验证)
  * 权鉴用 `JWT` 方式
    * 永久 token — 不能 `明文传输`
    * 需要登录才能产生的有 validation duration 的 token，比如 `-XPOST -H "Authorization: token值"`
  * token 还有一个，**防止重复提交**
    * 每次 request 和 reponse 都加入 32 uuid: `X-Request-Id`
* 查询优化: 代码端并发处理结果集, 数据库端设计表的时候就要加入分页项
* 自定义业务处理的加载时机 (包括构建 model 和解析model)
  * 包括加入 `切面操作`: 请求的日志记录, 插入某个 header在response中，API权限鉴定
  * 框架的中间件，或者自己写个中间层来处理，对于注册的函数进行调用(观察者模式)
  * 注意注册路由的时机和注册中间件的时机，先后顺序
* API版本问题 --- uri 路径解决
* API 请求记录问题
  * HTTP Request body 一旦被读取就会自动清空，所以记录API时，还需要重新赋值
  * 截获 HTTP 响应的话，则需要重定向到指定的 IO 流 — 一般只记录请求即可 (响应有后端逻辑)



> **技术演练**，到此为止；相关技术都已经非常清晰了，直接跳入 1.0 正式实现
>
> * 不同于 0.x 版本:
>   * 1.0版本会对用到的每个模块进行 test，包括 `gin` 框架本身
>   * 1.x 版本直接部署在服务器上，所以 `配置文件` 就不开源了, 即 conf 文件
>   * 部署方式采用 nginx + docker
>   * Makefile 管理源码，而不用 `go build`
> * 加入大量测试用例
>   * 如果不写测试用例的话，开发会快很多；但是后面增加模块或者原有模块出了问题就要哭了
>   * 包括性能测试

最后强调一下: `千万别明文传输`！(最好只启动 `ListenAndServeTLS`)



### 3.1 Version 1.0

在 `version 0.x` 基础上需要补充的功能:

* 删除不能是真的删除
* 可以运行时动态取消某个 API， 以及某组 API
* 运行时可以动态更换  `永久 token` (从 token list 中); 登录生成的 token 则不受影响
* 采用通行证机制: api认证的账户，后续所有通过 API 的应用都可以免认证
* 同时部署 http 和 https:
  * 内网还是 `http` 部署，直接走 IP (外网走内网IP，根本进不来)
  * 外网访问，比如查看 services， 必须走 https 部署



最终还是由于物理环境限制，不得不做很多妥协: 公网IP数量，内网服务器数量.

采用了 `supervisor` + `本地启动脚本(start|stop|status)`部署。





如果后期出现 **并发瓶颈**，那么考虑采用  `docker + nginx(upstream或者proxy_pass负载均衡)+多个程序实例`。

(应该不至于，仅内部服务于个人 + 家人 + 朋友圈)



BTW: Version 1.0 引入 `api.commonheart-yk.com/debug/pprof` 采样分析，但是不进行优化。

(如果确实后续接入系统过多或者实在存在性能问题，又无法扩充机器，再考虑代码优化)



#### 3.1.1 正式版开发过程



由于 `version 0.x` 已经把项目开发相关的各个技术都已经演练(预演)了，这里直接列举条目而不展开。

* `https://api.commonheart-yk.com` 展开服务(部署)
  * 主要解决证书问题 (https)
* 数据库表结构设计 (考虑到后续其他内部系统都会用到它)
  * 只涉及基础表
* 加载配置，日志，数据库，错误信息配置

到此，`xserver_status` 模块正常， `services` 模块不确定是否真的需要。

**暂停一下，根据后面子系统对于其的使用来确定后续开发的重点**，理想的情况：

> 除此之外，后面再无后端系统 （以后其他 web 系统都由此应用网关供能，其他系统只写前端，纯前端实现）

![08-32-44-api_gateway_halt.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/09/08-32-44-api_gateway_halt.jpg)


* 重新整理 API 版本问题 (**版本还是放在模块前面**)
* 新增 user 模块 (为成就系统)
  * 数据库只设计登录表结构 (后续补充 userinfo 表专门记录各项信息)
* 配置路由信息 (基础模块: xserver_status 和 services)
  * 包括 ORM ， 业务逻辑处理
* 加入 JWT 权限检查(永久token + 登录token)
* 加入启动脚本



* TODO: 同时加入 Swagger文档  --- 非必须
* TODO: 加入性能测试选项
* TODO: 加入 makefile


之后就是根据需要增添模块，即 handler 和 service 的逻辑。

TODO: Version2.0 解决热更新问题 (增加新模块免编译)



## 4 感谢



* 总体设计， 部分参考 `kong网关`系统。

* API 设计规范，这方面 `EasyDarwin` 网站提供的多媒体 API 编写规范可以成为不错的参考，主要就是字段，状态码，以及API格式，调用方式等等。有前辈的文档，就不要瞎折腾了。

宝典如下:

![16-50-11-api设计参考.jpg](https://hub.commonheart-yk.com/wiki/pics/2018/08/16-50-11-api设计参考.jpg)



**非常感谢开源社区** [EasyDarmin](https://www.easydarwin.org) 提供的方案，以及参考思路，感谢！



* 实现思路，感谢 腾讯后端高工 雷克斯。
* 微博开放文档- [微博API设计](http://open.weibo.com/wiki/Error_code)
* [Token认证](https://blog.csdn.net/qq_28098067/article/details/52036493)
* [《Web API 设计与开发》水野贵明(日), 工信出版社, 人民邮电出版社]()
