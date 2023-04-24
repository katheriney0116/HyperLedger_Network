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

注意，这里的host代表生成的证书能在哪些机器使用。如果需要有修改，我们需要重新enroll获得证书（如上述表述server config时一样）

![Untitled (1)](https://user-images.githubusercontent.com/101753393/233878261-c1c390b2-d1ab-4f08-a260-670ef4f4f6c0.png)
我们必须创建一个OU配置(config.yaml)并写到澳交所的证书下面（需要为每个组织做这一步）当连接fabric网络的时候，他会查看每个组织下面的msp folder有没有这个文件，没有的话会报错
```yaml
echo 'NodeOUs:
Enable: true
ClientOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: client
PeerOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: peer
AdminOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: admin
OrdererOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: orderer' > "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/config.yaml"
```
### 创造每个组织所需目录
1. 构造tlscacert目录（用于不同组织通信）
```bash
# 由于该CA同时充当组织CA和tlsca，因此直接将CA启动时生成的组织根证书复制到组织级CA和TLS CA目录中
mkdir -p "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/msp/tlscacerts"
cp "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem" "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/msp/tlscacerts/ca.crt"
```
2. 构造tlsca和ca目录（用于组织内客户端通信）
```bash
mkdir -p "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/tlsca"
cp "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem" "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/tlsca/tlsca.macaoE.microconnect.com-cert.pem"
mkdir -p "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/ca"
cp "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem" "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/ca/ca.macaoE.microconnect.com-cert.pem"
```
![Untitled (2)](https://user-images.githubusercontent.com/101753393/233878504-6a2fa403-899f-48f1-a607-e2d45b3627a4.png)

## 为澳交所注册新账户（peer节点、admin、client等）
当我们拥有ca-admin之后，我们可以使用刚刚得到的证书去给别的用户颁发证书

### register
注意register的时候client_home 要指向ca-admin的地址
```yaml
export FABRIC_CA_CLIENT_HOME=${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client register --caname ca-macaoE --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
fabric-ca-client register --caname ca-macaoE --id.name user1 --id.secret user1pw --id.type client --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
fabric-ca-client register --caname ca-macaoE --id.name org1admin --id.secret org1adminpw --id.type admin --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
fabric-ca-client register --caname ca-macaoE --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
```

### enroll
构造澳交所peer0的身份证书目录
```yaml
# 构造peer0的msp证书目录，证书文件会存在-M指定的文件夹下
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7050 --caname ca-macaoE -M "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/msp" --csr.hosts peer0.macaoE.microconnect.com --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
cp "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/msp/config.yaml" "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/msp/config.yaml"
# 构造peer0的msp-tls证书目录
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7050 --caname ca-macaoE -M "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls" --enrollment.profile tls --csr.hosts peer0.macaoE.microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
# 构造peer0的tls证书目录并格式化文件名——用于启动peer docker容器
cp "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls/ca.crt"
cp "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls/signcerts/"* "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls/server.crt"
cp "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls/keystore/"* "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/peers/peer0.macaoE.microconnect.com/tls/server.key"
```
构造user1的msp证书目录，因为user是通过peer进行通信，他不用与组织间通信，也不同配置tls
```bash
fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-macaoE -M "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/users/User1@macaoE.microconnect.com/msp" --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
cp "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/msp/config.yaml" "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/users/User1@macaoE.microconnect.com/msp/config.yaml"

# 构造org1admin的msp证书目录
fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-macaoE -M "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/users/Admin@macaoE.microconnect.com/msp" --tls.certfiles "${PWD}/orgs/fabric-ca/macaoE/ca-cert.pem"
cp "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/msp/config.yaml" "${PWD}/orgs/peerOrganizations/macaoE.microconnect.com/users/Admin@macaoE.microconnect.com/msp/config.yaml"
```

### 用相同方式建造spv和orderer组织的证书目录
```yaml
mkdir -p orgs/peerOrganizations/spv.microconnect.com/
export FABRIC_CA_CLIENT_HOME=${PWD}/orgs/peerOrganizations/spv.microconnect.com/
#dont forget to edit client.config
fabric-ca-client enroll -u https://admin:adminpw@localhost:7051 --caname ca-spv --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"

echo 'NodeOUs:
Enable: true
ClientOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: client
PeerOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: peer
AdminOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: admin
OrdererOUIdentifier:
Certificate: cacerts/localhost-8054-ca-spv.pem
OrganizationalUnitIdentifier: orderer' > "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/config.yaml"

mkdir -p "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/tlscacerts"
cp "${PWD}/orgs/fabric-ca/spv/ca-cert.pem" "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/tlscacerts/ca.crt"
mkdir -p "${PWD}/orgs/peerOrganizations/spv.microconnect.com/tlsca"
cp "${PWD}/orgs/fabric-ca/spv/ca-cert.pem" "${PWD}/orgs/peerOrganizations/spv.microconnect.com/tlsca/tlsca.spv.microconnect.com-cert.pem"
mkdir -p "${PWD}/orgs/peerOrganizations/spv.microconnect.com/ca"
cp "${PWD}/orgs/fabric-ca/spv/ca-cert.pem" "${PWD}/orgs/peerOrganizations/spv.microconnect.com/ca/ca.spv.microconnect.com-cert.pem"

fabric-ca-client register --caname ca-spv --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"
fabric-ca-client register --caname ca-spv --id.name user1 --id.secret user1pw --id.type client --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"
fabric-ca-client register --caname ca-spv --id.name org2admin --id.secret org2adminpw --id.type admin --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"

fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7051 --caname ca-spv -M "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/msp" --csr.hosts peer0.spv.microconnect.com --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"
cp "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/config.yaml" "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/msp/config.yaml"
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7051 --caname ca-spv -M "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls" --enrollment.profile tls --csr.hosts peer0.spv.microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"

cp "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls/ca.crt"
cp "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls/signcerts/"* "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls/server.crt"
cp "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls/keystore/"* "${PWD}/orgs/peerOrganizations/spv.microconnect.com/peers/peer0.spv.microconnect.com/tls/server.key"

fabric-ca-client enroll -u https://user1:user1pw@localhost:7051 --caname ca-spv -M "${PWD}/orgs/peerOrganizations/spv.microconnect.com/users/User1@spv.microconnect.com/msp" --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"
cp "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/config.yaml" "${PWD}/orgs/peerOrganizations/spv.microconnect.com/users/User1@spv.microconnect.com/msp/config.yaml"

fabric-ca-client enroll -u https://org2admin:org2adminpw@localhost:7051 --caname ca-spv -M "${PWD}/orgs/peerOrganizations/spv.microconnect.com/users/Admin@spv.microconnect.com/msp" --tls.certfiles "${PWD}/orgs/fabric-ca/spv/ca-cert.pem"
cp "${PWD}/orgs/peerOrganizations/spv.microconnect.com/msp/config.yaml" "${PWD}/orgs/peerOrganizations/spv.microconnect.com/users/Admin@spv.microconnect.com/msp/config.yaml"
```
```yaml
mkdir -p orgs/ordererorgs/microconnect.com
export FABRIC_CA_CLIENT_HOME=${PWD}/orgs/ordererorgs/microconnect.com

fabric-ca-client enroll -u https://admin:adminpw@localhost:7049 --caname ca-orderer --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"

echo 'NodeOUs:
Enable: true
ClientOUIdentifier:
Certificate: cacerts/localhost-9054-ca-orderer.pem
OrganizationalUnitIdentifier: client
PeerOUIdentifier:
Certificate: cacerts/localhost-9054-ca-orderer.pem
OrganizationalUnitIdentifier: peer
AdminOUIdentifier:
Certificate: cacerts/localhost-9054-ca-orderer.pem
OrganizationalUnitIdentifier: admin
OrdererOUIdentifier:
Certificate: cacerts/localhost-9054-ca-orderer.pem
OrganizationalUnitIdentifier: orderer' > "${PWD}/orgs/ordererorgs/microconnect.com/msp/config.yaml"

mkdir -p "${PWD}/orgs/ordererorgs/microconnect.com/msp/tlscacerts"
cp "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem" "${PWD}/orgs/ordererorgs/microconnect.com/msp/tlscacerts/tlsca.microconnect.com-cert.pem"
mkdir -p "${PWD}/orgs/ordererorgs/microconnect.com/tlsca"
cp "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem" "${PWD}/orgs/ordererorgs/microconnect.com/tlsca/tlsca.microconnect.com-cert.pem"

fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
fabric-ca-client enroll -u https://orderer:ordererpw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/msp" --csr.hosts microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"

cp "${PWD}/orgs/ordererorgs/microconnect.com/msp/config.yaml" "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/msp/config.yaml"

fabric-ca-client enroll -u https://orderer:ordererpw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls" --enrollment.profile tls --csr.hosts orderer.microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"

cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/ca.crt"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/signcerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/server.crt"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/keystore/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/server.key"
mkdir -p "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/msp/tlscacerts"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/msp/tlscacerts/tlsca.microconnect.com-cert.pem"

fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/users/Admin@microconnect.com/msp" --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
cp "${PWD}/orgs/ordererorgs/microconnect.com/msp/config.yaml" "${PWD}/orgs/ordererorgs/microconnect.com/users/Admin@microconnect.com/msp/config.yaml"

#OrdererA 和OrdererB的注册
fabric-ca-client register --caname ca-orderer --id.name ordererA --id.secret ordererApw --id.type orderer --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
fabric-ca-client register --caname ca-orderer --id.name ordererB --id.secret ordererBpw --id.type orderer --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
fabric-ca-client enroll -u https://ordererA:ordererApw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/msp" --csr.hosts microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
fabric-ca-client enroll -u https://ordererB:ordererBpw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/msp" --csr.hosts microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"

cp "${PWD}/orgs/ordererorgs/microconnect.com/msp/config.yaml" "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/msp/config.yaml"
cp "${PWD}/orgs/ordererorgs/microconnect.com/msp/config.yaml" "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/msp/config.yaml"

fabric-ca-client enroll -u https://ordererA:ordererApw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls" --enrollment.profile tls --csr.hosts ordererA.microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"
fabric-ca-client enroll -u https://ordererB:ordererBpw@localhost:7049 --caname ca-orderer -M "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls" --enrollment.profile tls --csr.hosts ordererB.microconnect.com --csr.hosts localhost --tls.certfiles "${PWD}/orgs/fabric-ca/ordererOrg/ca-cert.pem"

cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/ca.crt"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/signcerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/server.crt"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/keystore/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/server.key"
mkdir -p "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/msp/tlscacerts"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/msp/tlscacerts/tlsca.microconnect.com-cert.pem"

cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/ca.crt"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/signcerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/server.crt"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/keystore/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/server.key"
mkdir -p "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/msp/tlscacerts"
cp "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/tlscacerts/"* "${PWD}/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/msp/tlscacerts/tlsca.microconnect.com-cert.pem"
```
到此，我们完成所有证书的注册
