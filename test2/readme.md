# 创造一个HyperLedger网络

我们将HyperLedger 部署在两台主机上 - 10.10.10.123 和 10.10.10.124

为了简化步骤，我们会将所有证书在10.10.10.124上先启动，然后把124节点所需证书转移到124上

所有节点所需资料如下（加粗部分为实际操作中需要注意的端口）


| domain/container  | host | port  | describtion |
|---- |-- | ------|-------------|
| ca_macaoE |10.10.10.123 |7050 |澳交所组织的证书机构，会在发布完所有证书后关闭|
| ca_spv|10.10.10.123 |7051 |spv组织的证书机构，会在发布完所有证书后关闭|
| ca_orderer |10.10.10.123 |7049 |orderer组织的证书机构，会在发布完所有证书后关闭 |
| orderer.microconnect.com | 10.10.10.123| **8052(general) 8053(admin)** 9445(operation) |排序组织里的其中一个排序节点 |
| ordererA.microconnect.com | 10.10.10.123| **7052(general) 7053(admin)** 9446(operation) |排序组织里的其中一个排序节点 |
| ordererB.microconnect.com | 10.10.10.124| **9050(general) 9052(admin)** 9447(operation) |排序组织里的其中一个排序节点 |
| peer0.macaoE.microconnect.com | 10.10.10.123| **8054(peer_address)** 8055(chaincode) 8056(operation) |澳交所组织的一个peer节点 |
| peer1.macaoE.microconnect.com | 10.10.10.123| **8057(peer_address)** 8058(chaincode) 8059(operation) |澳交所组织的一个peer节点 |
| peer0.spv.microconnect.com | 10.10.10.124| **8055(peer_address)** 8052(chaincode) 9450(operation) |spv组织的一个peer节点 |

## 启动Fabric-CA服务

因为要通过124虚拟机生成所有证书文件，我们需要将本地DNS指向124
```bash
echo "127.0.0.1       macaoE.microconnect.com" >> /etc/hosts
echo "127.0.0.1       spv.microconnect.com" >> /etc/hosts
echo "127.0.0.1       microconnect.com" >> /etc/hosts

echo "127.0.0.1       orderer.microconnect.com" >> /etc/hosts
echo "127.0.0.1       ordererA.microconnect.com" >> /etc/hosts
echo "127.0.0.1       ordererB.microconnect.com" >> /etc/hosts

echo "127.0.0.1       peer1.macaoE.microconnect.com" >> /etc/hosts
echo "127.0.0.1       peer0.macaoE.microconnect.com" >> /etc/hosts
echo "127.0.0.1       peer0.spv.microconnect.com" >> /etc/hosts
```
[Fabric-CA介绍](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/documents/Fabric-CA.md)

在此文章里，我们用124虚拟机生成所有证书文件

然后使用`docker stop $(docker ps -aq)`关掉CA服务容器

将需要部署在123的节点证书copy到123虚拟机里，并将两边的dns文件都改成

```bash
echo "10.10.10.124       macaoE.microconnect.com" >> /etc/hosts
echo "10.10.10.123       spv.microconnect.com" >> /etc/hosts
echo "10.10.10.124       microconnect.com" >> /etc/hosts

echo "10.10.10.124       orderer.microconnect.com" >> /etc/hosts
echo "10.10.10.124       ordererA.microconnect.com" >> /etc/hosts
echo "10.10.10.123       ordererB.microconnect.com" >> /etc/hosts

echo "10.10.10.124       peer1.macaoE.microconnect.com" >> /etc/hosts
echo "10.10.10.124       peer0.macaoE.microconnect.com" >> /etc/hosts
echo "10.10.10.123       peer0.spv.microconnect.com" >> /etc/hosts
```

## 部署节点

Fabric节点在启动时会通过几种方式加载变量获取配置信息

默认情况下，Fabric节点主配置路径为 `FABRIC_CFG_PATH` 环境变量所指向路径。在不显式指定配置路径时，会尝试从主配置路径下查找相关的配置文件。

![image](https://user-images.githubusercontent.com/101753393/233884903-f05fca62-7cd6-4ab7-9813-4868ac2b703c.png)

所以，在启动docker之前，我们一定要先去定义`FABRIC_CFG_PATH`的环境变量，不然启动节点时会报错

在本实验中，我们设定的path是`export FABRIC_CFG_PATH=/root/test2/config`，里面包含了以下所有的配置文件

- core.yaml (peer）
- orderer.yaml(orderer)
- configtx.yaml (channel)

根据此文档，我们可依次启动Peer节点和Orderer节点

[部署peer和orderer节点](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/documents/SetupNode.md)

## 创造通道

- [创造通道的创世区块并将orderer节点加入通道](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/documents/ChannelConfig.md)

- [peer节点加入通道](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/documents/PeerJoinChannel.md)

到此为止，所有节点都已加入通道，并可使用

## 链码（智能合约）

- [安装链码到指定通道](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/documents/InstallChaincode.md)
- [更新链码步骤](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/documents/UpdateChaincode.md)

## CouchDB(数据库浏览器）

因为我们在定义peer节点的时候，添加了couchDB的环境变量，我们可以用couchDB的浏览器去查看我们之前启动智能合约时写入的数据（具体couchdb的container已经在上面部署节点的时候启动了）

具体流程可参考[此文章](https://ifantasy.net/2022/08/24/hyperledger_fabric_10_complex_contract_and_couchdb//)

![image](https://user-images.githubusercontent.com/101753393/233954996-0f4f2729-28f9-4efc-b33d-697740454610.png)

在本地浏览器输入`http://10.10.10.124:7255/_utils/#login`

![image](https://user-images.githubusercontent.com/101753393/233955624-90243acd-b563-418e-9669-23ab6810c2af.png)

![image](https://user-images.githubusercontent.com/101753393/233955768-b8e3a117-33ab-4b3f-ae89-fbd8f43df5cb.png)

![image](https://user-images.githubusercontent.com/101753393/233955822-69b12ab3-a60b-4cc3-bc9b-5ad292c9eecc.png)

![image](https://user-images.githubusercontent.com/101753393/233955854-7b6290f2-3894-4c61-85d5-21465c1de2a9.png)




## works cited

https://github.com/wefantasy/FabricLearn
