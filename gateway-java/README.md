# Gateway-java 学习

## Fabric-gateway是什么？

Fabric-gateway是一个提供给我们与fabric网络交互的SDK，以java来实现。

在2.4版本后，Fabric提倡使用[Fabric Gateway Client API](https://hyperledger.github.io/fabric-gateway/)，实现用最少的代码与链码交互

### 流程

1. 配置钱包
2. 配置gateway
3. 获取网络
4. 获取合约
5. 调用合约

  更详细资料，见[此处](https://github.com/katheriney0116/HyperLedger_Network/blob/main/gateway-java/HyperLedger%20Fabric%20API.pdf)

## 配置

1. Maven依赖

将以下代码加入project的`pom.xml`文件
```bash
<dependency>
  <groupId>org.hyperledger.fabric</groupId>
  <artifactId>fabric-gateway</artifactId>
  <version>1.1.0</version>
</dependency>
```

2. Gradle依赖
将以下代码加入project的`build.gradle`文件
```bash
implementation 'org.hyperledger.fabric:fabric-gateway:1.1.0'
```

3. 将一个peer节点的所有证书保存在本地project上 （如图）
![image](https://user-images.githubusercontent.com/101753393/233925727-e718a6d0-b57f-4397-be07-e880c6e882e2.png)

4. 将节点信息、证书写入 App.java 用于访问
注意，这里所有路径都是本地路径
![image](https://user-images.githubusercontent.com/101753393/233925944-8c6ffe50-120e-480c-b703-c0358f47abb2.png)

## Gateway-client api java

所有package和method可以在[此处](https://hyperledger.github.io/fabric-gateway/main/api/java/)找到

## App.java

这里gateway的示例是从[fabric-samples里面](https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-basic/application-gateway-java)摘录

