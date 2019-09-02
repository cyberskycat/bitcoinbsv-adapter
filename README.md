# bitcoincash-adapter

bitcoincash-adapter继承了bitcoin-adapter，主要修改了如下内容：

- 重写了Symbol = "BCH"。
- 重写了addressDecoder，实现了BCH地址编码。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建BCH.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url
serverAPI = "http://127.0.0.1:20009"
# RPC Authentication Username
rpcUser = "user"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = true
# support segWit
supportSegWit = false
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```
