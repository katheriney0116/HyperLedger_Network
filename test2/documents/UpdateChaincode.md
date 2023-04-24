# 更新链码步骤

## 仅仅更新背书策略

1. 组织同意新的背书策略：每次升级sequence都需+1
```bash
#例子：通过同意策略（需要org1和org2都为交易签名的mycc链码定义
peer lifecycle chaincode approveformyorg --channelID mychannel --signature-policy "AND('Org1.member', 'Org2.member')" --name mycc --version 1.0 --sequence 2.0 --package-id mycc_1:3a8c52d70c36313cfebbaf09d8616e7a6318ababa01c7cbe40603c373bcfe173 --sequence 1 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --waitForEvent
```

2. 提交：当足够数量的通道成员批准新的链码定义时，一个组织可以提交新定义以将链码定义升级到通道
```bash
#例子
peer lifecycle chaincode commit -o orderer2.council.ifantasy.net:7054 --tls --cafile $ORDERER_CA --channelID testchannel --name basic --version 1.0 --sequence 1 --peerAddresses peer1.soft.ifantasy.net:7251 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer1.web.ifantasy.net:7351 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```
3. 如果在不更改链码版本的情况下更新了链码定义(如仅修改背书策略的情况)，则链码容器将保持不变，并且无需调用Init函数。

## 更新链码文件里的内容

1. 重新打包链码

```bash
peer lifecycle chaincode package .....
```

2. 重新安装链码

```bash
peer lifecycle chaincode install
```

3. 组织同意新的链码定义

注意要更新链码定义中的链码版本（version）和链码sequence，packageID也会返回新的
```bash
peer lifecycle chaincode approveformyorg -o localhost:8053 --tls --cafile $ORDERER_CA  --channelID channel3 --name basic --version 2.0 --sequence  2 --waitForEvent --package-id $PACKAGE_ID
```

4. 提交链码定义:当足够数量的通道成员批准了新的链码定义时，一个组织可以提交新定义以将链码定义升级到通道。

```bash
#例子
peer lifecycle chaincode commit -o orderer2.council.ifantasy.net:7054 --tls --cafile $ORDERER_CA --channelID testchannel --name basic --version 2.0 --sequence 2 --peerAddresses peer1.soft.ifantasy.net:7251 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer1.web.ifantasy.net:7351 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
```
