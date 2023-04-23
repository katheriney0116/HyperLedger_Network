# Fabric CA 介绍
在搭建网络前，我们得先给管理者和网络节点颁发证书。这个证书可以用来识别属于每个组织的组件。

由服务端(Fabric-ca-server)和客户端(fabric-ca-client)组件组成。

正常情况下，生产环境中，不同组织之间都有自己的CA证书服务器。

server可以看成是一个web服务，执行go代码编译生成的二进制文件后，会监听一个端口，处理收到的请求。

client是向ca服务端发送请求的程序。

## 启动server
可在本地或docker内启动服务，如在docker内启动需要用docker-compose。

新建一个compose 文件夹，我们先写一个 [ca-compose.yaml](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/compose/ca-compose.yaml) 去定义 server所需的变量。

在这里，我们定义了三个ca组织，分别是澳交所的ca组织，spv的ca组织，和orderer的ca组织
### 启动docker-compose获取server config文件
```bash
docker-compose up -f ca-compose.yaml -d
```
[pic]

第一次启动server后，我们在Fabric_CA_SERVER_HOME 下有server的配置文件：fabric-ca-server-config.yaml和其他相关证书信息

如需要修改（比如说x509证书相关的信息，我们可以在此文件里进行修改，如图所示）

- CN是常见名称
- O是组织名称
- OU 是组织单位
- L 是位置或城市
- ST 是国家
- C是国家
-[pic]

然后删除除config外其他所有file，重新启动server
```bash
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker-compose -f ca-compose.yaml up -d
```
当重新启动server时，我们获得了新的证书
【pic】
### server可以生成的证书定义
1. ca-cert.pem - 根证书文件（同个组织节点有相同的根证书）
2. tls-cert.pem - 由ca-cert.pem签发的证书，是fabric-ca-server的tls通讯证书。当fabric-ca-client向fabric-ca-server发出通讯请求时，fabric-ca-server就展示tls-cert的证书。fabric-ca-client有根证书就可以用根证书文件去进行验证
3. msp - 包含keystore，CA服务器的私钥
4. fabric-ca-server.db ：CA默认使用的嵌入型数据库 SQLite,存储发放证书信息，可以通过sqllite3或其他可视化工具查看详细信息。

我们可以通过openssl来查看证书内容
```bash
openssl x509 -in /root/test2/orgs/macaoE.microconnect.com/ca/crypto/ca-cert.pem -inform pem -noout -text
```

## 当server生成完证书后，我们可以使用fabric-ca-client来申请并颁发用户、节点证书
有两个步骤：

1. fabric-ca-client register - 登记注册用户的过程
2. fabric-ca-celint enroll -  实际颁发用户证书的过程 - 将证书下载到指定目录

必须得是admin才可以颁发账户

因为在我们启动docker server的时候，我们已经运行了以下命令，相当于register了admin，我们现在只需要enroll admin去获得admin的实际证书
```bash
command: sh -c 'fabric-ca-server start -d -b ca-admin:ca-adminpw --port 7050'
```
### 准备证书
之前开启了ca-client 和 ca-server之间的tls通信。当ca-client 发起 tls通讯的时候，需要通过根证书ca-cert.per来验证server所提供tls-cert.perm证书的有效性

首先，我们需要设置client端的工作路径 - FABRIC_CA_CLIENT_HOME用于指定client的工作路径

为了提供根证书，我们可以直接把ca-client端的（TLS_CERTIFILES)环境变量直接指定到存储ca-cert的地方
```bash
#创建澳交所证书目录
mkdir -p orgs/peerOrganizations/macaoE.microconnect.com/
#enroll组织默认管理员账户，其配置对应在compose/compose-ca.yaml的command中，enroll过程会获取该账户的全部证书并保存至FABRIC_CA_CLIENT_HOME目录下

#caname 在compose-ca.yaml里
export FABRIC_CA_CLIENT_HOME=/root/test2/orgs/peerOrganizations/macaoE.microconnect.com/
export FABRIC_CA_CLIENT_MSPDIR=msp
```
### 获取admin证书（发起enroll）
“enroll”它会以我们刚刚设置的环境变量里的FABRIC_CA_CLIENT_TLS_CERTFILES 指向的CA服务器根证书加密通信并将生成的身份证书保存在环境变量FABRIC_CA_CLIENT_HOME 里
```bash
fabric-ca-client enroll -u https://admin:adminpw@localhost:7050 --caname ca-macaoE --tls.certfiles "/root/test2/orgs/fabric-ca/macaoE/ca-cert.pem"
```
上面的操作会将证书和私钥都保存在client端
![enroll图](https://user-images.githubusercontent.com/101753393/233832877-14be9650-fa40-48b5-8f5e-6d87522ae984.png)

