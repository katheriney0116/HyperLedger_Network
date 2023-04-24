# 配置通道

从分析的角度来看fabric，fabric是由底层的区块链配置（由orderer节点组成的区块链网路）和上层的application配置（由peer节点和智能合约）组成。

底层的配置主要包括orderer节点的共识配置

系统的初始化配置都保存在创世区块中而上层的application配置储存在交易和peer中

在Fabric最新版本中（fabric 2.3），fabric简化了通道创建的流程并增加了隐私性和通道的拓展性。现在创建应用通道可以用以下方法来实现

## orderer证书目录
 
因为创造通道的指令必须部署到orderer的节点上，我们通道上至少要存在一个orderer节点。这一步我们已经在证书部分完成
![Untitled (4)](https://user-images.githubusercontent.com/101753393/233881853-8ee7c871-1379-4127-ba15-be7b39143f3b.png)

在上一步部署节点里，我们已经启动了所有的节点。但是，目前的节点都没有加入任何通道。现在我们要首先要创造通道的创世区块。

在执行指令之前，我们用[configtx.yaml](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/config/configtx.yaml)来配置通道

- tx是交易的缩写，所有关于交易相关的所有配置，例如应用通道，锚节点，orderer服务等都是在此文件里配置
- 配置工具（bin）： configtxgen （读取configtx。形成创世区块的输出）

利用configtx里profile里的配置去创建系统通道，输出保存在genensis.block里

文件中有以下配置我们是需要进行设置的
![Untitled (5)](https://user-images.githubusercontent.com/101753393/233882182-211258f3-36bf-45db-83c1-f74466f2d88e.png)

具体配置请看
[configtx.yaml](https://github.com/katheriney0116/HyperLedger_Network/blob/main/test2/config/configtx.yaml)

当编辑完此文件后，我们可以执行以下命令来创建新的创世区块和通道名称.

```bash
configtxgen -profile TwoOrgsGenesis -outputBlock ./channel-artifacts/genesis.block -channelID channel3
```
- profile是configtx.yaml 用来创建通道的。（注意，这里的profile name一定要和 configtx里的profile label 一样）
- outputBlock生成创始区块的保存路径
- channelID通道的名称，名字必须全小写，少于250个字符并且匹配表达式[a-z][a-z0-9.-]*。

当命令执行成功时，我们会看到
![Untitled (6)](https://user-images.githubusercontent.com/101753393/233882499-4be1a530-aa27-4642-b9cf-886cd5407161.png)

## 使用osnadmin把第一个orderer加到通道
现在创始区块已经创建了，第一个orderer节点接收”osnadmin channel join“命令激活 “activates” 通道。

尽管通道并不是完全可操作的，直到法定达到法定人数（如果你的profile列出了三个共识者（orderer），则需要两个以上，使用osnadmin channel join命令加入到通道中）。
注意，在osnadmin命令里后面的 -o 一定要是docker compose里定义的admin的端口，不然会报错
```yaml
export ORDERER_CA=/root/test2/orgs/ordererorgs/microconnect.com/tlsca/tlsca.microconnect.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=/root/test2/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=/root/test2/orgs/ordererorgs/microconnect.com/orderers/orderer.microconnect.com/tls/server.key

osnadmin channel join --channelID channel3 --config-block ./channel-artifacts/genesis.block -o orderer.microconnect.com:8053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
```
当我们看到这个就表示第一个orderer加入区块链成功
![Untitled (7)](https://user-images.githubusercontent.com/101753393/233882686-c297585b-12a6-410b-8c52-810314db5513.png)

使用一下命令能看到当前orderer所加入的channel (Hyperledger 里允许节点加入多个通道，每个通道交易完全独立）
![Untitled (8)](https://user-images.githubusercontent.com/101753393/233882775-a73cbd4a-047b-443b-8b43-b5c045372a72.png)

用相同方式把剩下两个orderer加入通道，通道（此区块链账本）就可以开始使用
```yaml
#启动ordererA
export ORDERER_CA=/root/test2/orgs/ordererorgs/microconnect.com/tlsca/tlsca.microconnect.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=/root/test2/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=/root/test2/orgs/ordererorgs/microconnect.com/orderers/ordererA.microconnect.com/tls/server.key

osnadmin channel join --channelID channel3 --config-block ./channel-artifacts/genesis.block -o ordererA.microconnect.com:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
```

换到服务器124启动ordererB
```yaml
export ORDERER_CA=/root/test2/orgs/ordererorgs/microconnect.com/tlsca/tlsca.microconnect.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=/root/test2/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=/root/test2/orgs/ordererorgs/microconnect.com/orderers/ordererB.microconnect.com/tls/server.key

osnadmin channel join --channelID channel3 --config-block ./channel-artifacts/genesis.block -o ordererB.microconnect.com:9052 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
```
