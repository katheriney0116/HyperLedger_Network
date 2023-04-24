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
[Fabric-CA介绍](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/Fabric-CA.md)

在此文章里，我们用123虚拟机生成所有证书文件

然后使用`docker stop $(docker ps -aq)`关掉CA服务容器
