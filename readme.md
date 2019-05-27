# go-for-apache-dubbo-benchmark


## 使用方式

### 调用方式

脚本参数含义：

参数|含义|取值范围
-------------|-------------|-------------
c|并发client|大于0的int值
n|请求总数|大于0的int值
p|测试协议|dubbo or jsonrpc

#### 1.客户端和服务端在同一台机器上

sh start.sh -c 10 -n 100 -p dubbo

#### 2.客户端和服务端在不同一台机器上，包括多个服务端在不同机器上

server: 

1)需要配置server.yml参数

2)执行脚本 sh start_server.sh  -p dubbo

client: 

1)需要配置client.yml参数

2)执行脚本 sh start_client.sh  -c 10 -n 100



### 输出报表含义

`TPS`:吞吐率

`mean`: 单个请求平均耗时

`max`: 单个请求的最大耗时

`min`: 单个请求的最小耗时

`p99`: 99%的请求单个耗时


例子：

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------
100|0|0|20|0|55561
500|7|6|59|0|62593
1000|14|12|103|0|65329
2000|28|24|163|0|67033
5000|71|64|380|0|63803
