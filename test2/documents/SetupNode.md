# 节点部署 

## Peer节点
当Peer节点启动时，会按照以下优先级从高到低的顺序依次尝试从中读取配置信息
1. 命令行参数、(cli)
2. 环境变量 (docker-compose)
3. 配置文件(core) - 必须存在，不然会报错

[具体peer配置解析翻译](https://blog.csdn.net/weixin_45839894/article/details/123450112)

不论哪种方式，我们都需要注意改写以下的几个值

![Untitled (9)](https://user-images.githubusercontent.com/101753393/233886147-143e8c48-441d-41cb-a91b-e27c26a96d84.png)



## orderer节点
同样，当Orderer节点启动时，会按照以下优先级从高到低的顺序依次尝试从中读取配置信息
1. 命令行参数、(cli)
2. 环境变量 (docker-compose)
3. 配置文件(orderer) - 不一定存在，不会报错如果不存在

不论哪种方式，我们都需要注意改写以下的几个值

![Untitled (10)](https://user-images.githubusercontent.com/101753393/233888164-4ce65245-0a4e-4608-bb17-54d717247904.png)

## 启动节点

当orderer和peer的配置写好后，我们就可以启动他们的容器

在此实验里，我们先在`compose/docker-base`定义[peer和orderer节点的基础配置](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/compose/docker-base.yaml)

在 [net-compose](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/compose/net-compose.yaml)里定义每个节点具体的值

再在每个服务器（123和124）启动相应的hostn.yaml的配置来启动所有节点(可以将net-compose里的配置写入hostn.yaml来去繁）

[123所在节点配置](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/compose/host1.yaml)
```bash
#在123服务器里的compose文件夹里启动
docker-compose -f host1.yaml -d
```
[124所在节点配置](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/compose/host2.yaml)
```bash
#在124服务器里的compose文件夹里启动
docker-compose -f host2.yaml -d
```
