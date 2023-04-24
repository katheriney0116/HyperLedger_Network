# HyperLedger_Network
 A demo of hyperledger network including a functional test network(couchDB), two contract with fabric-gateway and fabric-explorer implemented.
 
 ## Fabric 2.4 Prerequisite
 
 所有的实验都是在Linux下完成的。
 
 - Install git `apt install git`
 
 - Install cURL  `apt install cURL`
 
 - Install golang  `apt install golang`
 
 - Install jq  `apt install jq`
 
 - Install Fabric-samples (fabric samples是Fabric官方的demo集合，里面包含多个实例)
 ```bash
 git clone https://github.com/hyperledger/fabric-samples
 ```

 - Install Fabric (Fabric是联盟链核心开发工具，包含开发部署所有命令)
   ```bash
   #下载fabric 2.4并解压
   wget https://github.com/hyperledger/fabric/releases/tag/v2.4.0/hyperledger-fabric-linux-amd64-2.4.0.tar.gz
   mkdir /usr/local/fabric
   tar -xzvf hyperledger-fabric-linux-amd64-2.3.2.tar.gz -C /usr/local/fabric
   ```
 - Install Fabric-ca (Fabric 自带的证书系统）
   ```bash
   #下载fabric-ca1.5.2并解压
   wget https://github.com/hyperledger/fabric-ca/releases/tags/v1.5.2/hyperledger-fabric-ca-linux-amd64-1.5.2.tar.gz
   tar -xzvf hyperledger-fabric-ca-linux-amd64-1.5.2.tar.gz
   mv bin/* /usr/local/fabric/bin
   ```
 - Set environment variable
   ```bash
    export FABRIC=/usr/local/fabric
    export PATH=$PATH:$FABRIC/bin
   ```
 - Install Docker
   ```bash
    apt install docker
    apt install docker-compose
   ```
 - Install Dockers Imager
   ```bash
   docker pull hyperledger/fabric-tools:2.4
   docker pull hyperledger/fabric-peer:2.4
   docker pull hyperledger/fabric-orderer:2.4
   docker pull hyperledger/fabric-ccenv:2.4
   docker pull hyperledger/fabric-baseos:2.4
   docker pull hyperledger/fabric-ca:1.5
   ```
 
 
 ## Contents
 - [多机完整部署所有节点+通道+链码安装+couchDB](https://github.com/katheriney0116/HyperLedger_Network/tree/main/test2)

 - [Fabric-Gateway学习](https://github.com/katheriney0116/HyperLedger_Network/tree/main/gateway-java)

 - [Fabric-Explorer学习](https://github.com/katheriney0116/HyperLedger_Network/tree/main/explorer)

 ## Folders Description
 ```bash
 ├── basic.tar.gz
├── channel-artifacts     #创世区块
│   └── genesis.block
├── compose               #docker-compose文件夹，用来启动所有节点
│   ├── ca-compose.yaml
│   ├── couchdb.yaml
│   ├── docker-base.yaml
│   ├── host1.yaml
│   ├── host2.yaml
│   └── net-compose.yaml
├── config                 #配置文件（通道，peer节点，和orderer）
│   ├── configtx.yaml
│   ├── core.yaml
│   └── orderer.yaml
├── orgs                   #所有组织证书
│   ├── fabric-ca          #各个组织公共材料目录
│   │   ├── macaoE
│   │   ├── ordererOrg
│   │   └── spv           
│   ├── ordererorgs         #排序组织目录
│   │   └── microconnect.com
│   │       ├── fabric-ca-client-config.yaml
│   │       ├── msp         #组织msp目录
│   │       ├── orderers    #各个排序节点目录
│   │       │   ├── ordererA.microconnect.com
│   │       │   ├── ordererB.microconnect.com
│   │       │   └── orderer.microconnect.com
│   │       ├── tlsca       #TLS-CA证书，用于不同组织通信
│   │       │   └── tlsca.microconnect.com-cert.pem
│   │       └── users       #各个用户目录
│   └── peerOrganizations   #各个peer目录
│       ├── macaoE.microconnect.com
|       |   ├── ca
│       │   ├── peers
│       │   │   ├── peer0.macaoE.microconnect.com
│       │   │   └── peer1.macaoE.microconnect.com
│       │   ├── tlsca
│       │   │   └── tlsca.macaoE.microconnect.com-cert.pem
│       │   └── users
│       │       ├── Admin@macaoE.microconnect.com
│       │       └── User1@macaoE.microconnect.com
│       └── spv.microconnect.com
│           ├── ca
│           │   └── ca.spv.microconnect.com-cert.pem
│           ├── fabric-ca-client-config.yaml
│           ├── msp
│           │   ├── cacerts
│           │   │   └── localhost-7051-ca-spv.pem
│           │   ├── config.yaml
│           │   ├── IssuerPublicKey
│           │   ├── IssuerRevocationPublicKey
│           │   ├── keystore
│           │   │   └── abdf1d008e8da773f97ad000d17a3f79aae3218bef772dbbbffa908a7e32d2c1_sk
│           │   ├── signcerts
│           │   │   └── cert.pem
│           │   ├── tlscacerts
│           │   │   └── ca.crt
│           │   └── user
│           ├── peers
│           │   └── peer0.spv.microconnect.com
│           │       ├── msp
│           │       │   ├── cacerts
│           │       │   │   └── localhost-7051-ca-spv.pem
│           │       │   ├── config.yaml
│           │       │   ├── IssuerPublicKey
│           │       │   ├── IssuerRevocationPublicKey
│           │       │   ├── keystore
│           │       │   │   └── 49b24d44b98ba36664d2fcf332acf1ab2e1bb8dd0554c39e6c71a3441a9e6834_sk
│           │       │   ├── signcerts
│           │       │   │   └── cert.pem
│           │       │   └── user
│           │       └── tls
│           │           ├── cacerts
│           │           ├── ca.crt
│           │           ├── IssuerPublicKey
│           │           ├── IssuerRevocationPublicKey
│           │           ├── keystore
│           │           │   └── cc24dc7dbd81586e555e219e07c18146c75282fcce1d5c7470ca228df64b0a16_sk
│           │           ├── server.crt
│           │           ├── server.key
│           │           ├── signcerts
│           │           │   └── cert.pem
│           │           ├── tlscacerts
│           │           │   └── tls-localhost-7051-ca-spv.pem
│           │           └── user
│           ├── tlsca
│           │   └── tlsca.spv.microconnect.com-cert.pem
│           └── users
│               ├── Admin@spv.microconnect.com
│               │   └── msp
│               │       ├── cacerts
│               │       │   └── localhost-7051-ca-spv.pem
│               │       ├── config.yaml
│               │       ├── IssuerPublicKey
│               │       ├── IssuerRevocationPublicKey
│               │       ├── keystore
│               │       │   └── 9300b6f8a778a2adaa407bdc3c6244dffc81c0adbfe875c43909c4b964beea97_sk
│               │       ├── signcerts
│               │       │   └── cert.pem
│               │       └── user
│               └── User1@spv.microconnect.com
│                   └── msp
│                       ├── cacerts
│                       │   └── localhost-7051-ca-spv.pem
│                       ├── config.yaml
│                       ├── IssuerPublicKey
│                       ├── IssuerRevocationPublicKey
│                       ├── keystore
│                       │   └── c029a61520b8f24ca5b30923c90b2ff0bdb83a9ec487b2e1d2b6d48d3e5061f2_sk
│                       ├── signcerts
│                       │   └── cert.pem
│                       └── user

```
