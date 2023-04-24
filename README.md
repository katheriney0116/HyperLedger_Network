# HyperLedger_Network
 A demo of hyperledger network including a functional test network(couchDB), two contract with fabric-gateway and fabric-explorer implemented.
 
 ## Fabric 2.4 Prerequisite
 
 所有的实验都是在Linux下完成的。
 
 - Install git `apt install git`
 
 - Install cURL  `apt install cURL`
 
 - Install golang  `apt install golang`
 
 - Install jq  `apt install jq`
 
 - Install Fabric-samples
 https://github.com/hyperledger/fabric-samples
   - fabric samples是Fabric官方的demo集合，里面包含多个实例
   - Fabric是联盟链核心开发工具，包含开发部署所有命令
   ```bash
   #下载fabric 2.4并解压
   wget https://github.com/hyperledger/fabric/releases/tag/v2.4.0/hyperledger-fabric-linux-amd64-2.4.0.tar.gz
   mkdir /usr/local/fabric
   tar -xzvf hyperledger-fabric-linux-amd64-2.3.2.tar.gz -C /usr/local/fabric
   
   #下载fabric-ca1.5.2并解压
   wget https://github.com/hyperledger/fabric-ca/releases/tags/v1.5.2/hyperledger-fabric-ca-linux-amd64-1.5.2.tar.gz
   tar -xzvf hyperledger-fabric-ca-linux-amd64-1.5.2.tar.gz
   mv bin/* /usr/local/fabric/bin
   
   #
 
 
 
 ## Contents
 - 
 - 完整部署所有节点+通道+链码安装
 - Fabric-Gateway学习
 - Fabric-Explorer学习
