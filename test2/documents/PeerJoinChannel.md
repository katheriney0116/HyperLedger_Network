Peer节点需要通过访问对应通道的block创世区块文件来加入。

在123虚拟机里，我们用以下环境变量去指定澳交所组织的peer0节点
```bash
export ORDERER_CA=/root/test2/orgs/ordererorgs/microconnect.com/tlsca/tlsca.microconnect.com-cert.pem
export FABRIC_CFG_PATH=/root/test2/config
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="macaoEMSP"
export PEER0_macaoE_CA=/root/test2/orgs/peerOrganizations/macaoE.microconnect.com/tlsca/tlsca.macaoE.microconnect.com-cert.pem
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_macaoE_CA
export CORE_PEER_MSPCONFIGPATH=/root/test2/orgs/peerOrganizations/macaoE.microconnect.com/users/Admin@macaoE.microconnect.com/msp

#这个指定的是我们在docker-compose里定义的peer0的值
export CORE_PEER_ADDRESS=localhost:8054
```

然后，我们用命令`peer channel join -b /root/test2/channel-artifacts/genesis.block` 去将macaoE组织的peer0加入通道channel3

