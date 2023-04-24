# 安装链码到指定通道

现在我们网络里有3个组织orderer组织，澳交所（macaoe） 和 spv。 

macaoE有两个节点peer0 和 peer1，spv组织只有一个节点peer0。每个节点都加入了channel3这一个通道。

在hyperledger fabric 的网络里，链码是需先安装到peer节点上，然后orderer排序节点再再通道里同意每个链码的定义。

因为我们在configtx定义了链码需要被上面#两个组织通过才能使用，我们需要在两个组织分别安装链码：

## 打包及下载链码

在123上执行以下命令将链码下载到澳交所的peer0和peer1上

```bash
#使用peer0
export CORE_PEER_ADDRESS=peer0.macaoE.microconnect.com:8054
#下载对应链码包
peer lifecycle chaincode package basic.tar.gz --path /root/test2/project_contract --label basic
#在peer0节点下载链码
peer lifecycle chaincode install basic.tar.gz #这一步会得到一串数据，将这串数据写进环境变量
export PACKAGE_ID=basic_1:06613e463ef6694805dd896ca79634a2de36fdf019fa7976467e6e632104d718
#在peer0节点查询已下载的所有链码（看下图）docke
peer lifecycle chaincode queryinstalled


#使用peer1
export CORE_PEER_ADDRESS=peer1.macaoE.microconnect.com:8057
#下载对应链码包
peer lifecycle chaincode package basic.tar.gz --path asset-transfer-basic/chaincode-go --label basic
#在peer0节点下载链码
peer lifecycle chaincode install basic.tar.gz
#在peer0节点查询已下载的所有链码
peer lifecycle chaincode queryinstalled
```
![Untitled (11)](https://user-images.githubusercontent.com/101753393/233906179-9baec9bd-7a9e-4991-ab50-4daa476204ce.png)
![Untitled (12)](https://user-images.githubusercontent.com/101753393/233906198-746f8625-3a79-4090-922f-eebccc2d87b9.png)
![Untitled (13)](https://user-images.githubusercontent.com/101753393/233906214-38e63623-8a77-4043-b15e-32f6ae149163.png)

在124上执行以下命令将链码下载到spv的peer0上

```bash
#记得要先确认我们在使用peer0的环境变量
export CORE_PEER_ADDRESS=peer0.spv.microconnect.com:8055
#下载对应链码包
peer lifecycle chaincode package basic.tar.gz --path asset-transfer-basic/chaincode-go --label basic
#在peer0节点下载链码
peer lifecycle chaincode install basic.tar.gz
#在peer0节点查询已下载的所有链码
peer lifecycle chaincode queryinstalled
```

## 组织批准链码

我们需要连接一个排序节点orderer去批准代码在**此通道**上允许运行

注意，我们在用`peer lifecycle chaincode approveformyorg`的时候需要确定需不需要 `--init-required`，如果加了这个命令的话后面的所有命令都要加，如果没有，则后面的指令都不用加
需要保持一致，不然会报错

回到123虚拟机
```bash
#批准申请,这里orderer用general端口就行
peer lifecycle chaincode approveformyorg -o orderer.microconnect.com:8053 --tls --cafile $ORDERER_CA  --channelID channel3 --name basic --version 1.0 --sequence 1 --waitForEvent --init-required --package-id $PACKAGE_ID
#查看已批准链码
peer lifecycle chaincode queryapproved -C channel3 -n basic --sequence 1
```
![Untitled (14)](https://user-images.githubusercontent.com/101753393/233906994-88485fd4-298f-4ba0-bb35-ea285f7354b8.png)
在完成第一个组织的批准后，我们可以用命令
`peer lifecycle chaincode checkcommitreadiness --channelID channel3 -n basic -v 1.0 --sequence 1 --output json --init-required`来查看所有组织批准情况

![Untitled (15)](https://user-images.githubusercontent.com/101753393/233907132-6581fb4f-da07-4412-9e62-6f176925b2e8.png)

到124虚拟机去让spv组织同意链码定义

```bash
#批准申请,这里orderer用general端口就行
peer lifecycle chaincode approveformyorg -o orderer.microconnect.com:8053 --tls --cafile $ORDERER_CA  --channelID channel3 --name basic --version 1.0 --sequence 1 --waitForEvent --init-required --package-id $PACKAGE_ID
#查看已批准链码
peer lifecycle chaincode queryapproved -C channel3 -n basic --sequence 1
```

## 提交链码

当两个组织都同意后，我们用命令提交到通道后
```bash
peer lifecycle chaincode commit -o orderer.microconnect.com:8052 --tls --cafile $ORDERER_CA --channelID channel3 --name basic --version 1.0 --sequence 1 --peerAddresses peer0.macaoE.microconnect.com:8054 --tlsRootCertFiles $PEER0_macaoE_CA --peerAddresses peer0.spv.microconnect.com:8055 --tlsRootCertFiles $PEER0_spv_CA
```
这个链码（智能合约）就可以在通道 channel3 使用了

## 使用链码

以下命令可以用来调用链码里的function，如果之前加了 `--init-required` 的，一定要先call initledger的function
```bash
peer chaincode invoke -o orderer.microconnect.com:8052 --tls --cafile $ORDERER_CA --channelID channel3 --name contract_2 --peerAddresses peer0.macaoE.microconnect.com:8054 --tlsRootCertFiles $PEER0_macaoE_CA --peerAddresses peer0.spv.microconnect.com:8055 --tlsRootCertFiles $PEER0_spv_CA -c '{"Args":["InitLedger"]}'
peer chaincode invoke -o orderer.microconnect.com:8052 --tls --cafile $ORDERER_CA --channelID channel3 --name contract_2 --peerAddresses peer0.macaoE.microconnect.com:8054 --tlsRootCertFiles $PEER0_macaoE_CA --peerAddresses peer0.spv.microconnect.com:8055 --tlsRootCertFiles $PEER0_spv_CA -c '{"Args":["GetAllAssets"]}'
```

