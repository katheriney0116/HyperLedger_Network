# Fabric-Explorer 学习及应用

本实验学习官方[Fabric-Explorer](https://github.com/hyperledger-labs/blockchain-explorer)文件并应用在了test2的实验网络中


### 下载三个官方配置文件 docker-compose.yaml,config.json,test-network.json

```bash
wget https://raw.githubusercontent.com/hyperledger/blockchain-explorer/main/examples/net1/config.json
wget https://raw.githubusercontent.com/hyperledger/blockchain-explorer/main/examples/net1/connection-profile/test-network.json -P connection-profile
wget https://raw.githubusercontent.com/hyperledger/blockchain-explorer/main/docker-compose.yaml
```

### 将test-network.json放到connection-profile的目录下，剩余放在explorer同级文件夹中
```bash
explorer
├── config.json
├── connection-profile
│   ├── macaoe.json
│   └── spv.json
├── docker-compose.yaml
└── orgs
```

### 修改三个配置文件

**docker-compose.yaml**

1. 修改networks配置，保持与之前peer和orderer节点的network相同

 ![image](https://user-images.githubusercontent.com/101753393/233938996-0696a9c3-3eb1-42c8-aaad-b48b28b9a4f9.png)
 
 ![image](https://user-images.githubusercontent.com/101753393/233939281-10abdc46-392d-43bc-ac4c-26e0a8e689fb.png)


2. 修改explorer.mynetwork.com的volumes，保证与你本地路径一致
![image](https://user-images.githubusercontent.com/101753393/233939231-d11b180b-a1c8-4ff7-9fa6-24c87472d5b3.png)


**test-network.yaml**

1.可改命名 -> macaoe.yaml

2. 将此处改为admin的密钥和证书

![image](https://user-images.githubusercontent.com/101753393/233940428-0c855830-3075-485a-85a0-281789109fd1.png)


3. 根据peer和channel名字修改代码，注意，这里是你登录explorer浏览器时的账号和密码

![image](https://user-images.githubusercontent.com/101753393/233940316-8fa960fe-9a41-42e8-93ec-0a7eff277d00.png)


**config.json**

保证profile的路径改成之前的`test-network.yaml`，现在的`macaoe.yaml`

![image](https://user-images.githubusercontent.com/101753393/233940625-c798a1e1-9915-4151-8b00-923f820c4c19.png)


### 启动

注意，一定要设置好以下三个变量，不然会报错
```bash
export EXPLORER_CONFIG_FILE_PATH=./config.json
export EXPLORER_PROFILE_DIR_PATH=./connection-profile
export FABRIC_CRYPTO_PATH=./orgs
```

启动区块链浏览器
```bash
docker-compose up -d
```
打开本地浏览器输入
```bash
10.10.10.124:8080
```
![image](https://user-images.githubusercontent.com/101753393/233941684-5a29b91e-9c49-42a1-b071-1d75106b45ed.png)

输入账号`exploreradmin`, 密码`exploreradminpw`

![image](https://user-images.githubusercontent.com/101753393/233941882-6ba35b6d-8928-4ff6-a53a-76e335b7e25c.png)




