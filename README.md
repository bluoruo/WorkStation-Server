# WorkStation Server
WorkStation的服务器端，支持远程管理windows、linux、mac、openwrt等设备，自带ddns功能，\
并可以远程开启和关闭客户端服务。支持监控客户端的CPU、内存、硬盘使用率。配置客户端的ddns接口（目前支持aliyun的dns）信息，并能自动更新。\
服务器端支持 Windows Linux Mac Freebsd的x86/arm等架构

## 服务器端接口:
### /v1/receive
###### 接收GET/POST的信息并完整的记录到日志文件
### /v2/config
###### 客户端初始化时候发送基本配置给客户端
Method GET:\
type=ip 客户端的IPv4地址\
Method POST:\
type=ip 客户端的IPv4地址\
type=config 服务器端配置（加密）\
type=update 客户端更新版本信息 \
无参数: Welcome to WorkStation Server.
### /v2/base
###### 基本信息接口
Method POST:
type=checkname 客户端识别名 \
type=edit 修改客户端识别名 \
type=getinfo  获取基本信息

### /v2/repo
###### 信息提交接口
Method POST:\
把客户信息记录到 wsc_info

### /v2/ddns
###### ddns接口
Method POST:\
DDNS域名配置信息

### /v2/status
###### 状态接口
Method POST:\
把客户端的CPU、内存、硬盘等信息记录到 wsc_status\
~~并可以再次插入服务器端发送给客户端的指令~~


## 版本功能
### 0.01.1
增加多配置文件读取功能
1.增加redis操作功能\
2.增加缓存机制\
3.增加执行操作其他微服务功能\
4.增加爬功能
### 0.01.2
1.增加v3加密接口\
2.增加鉴权和token功能\
3.增加和im的互动功能\