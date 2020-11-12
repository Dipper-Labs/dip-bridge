# 如何映射

## 1. 浏览器安装MetaMask并导入以太坊账号
参考MetaMask官方: https://metamask.io/

### 1.1 浏览器安装MetaMask
### 1.2 导入要映射的以太坊账号到MetaMask


## 2. 创建Dipper Network主网地址
可以通过官方钱包完成也可以通过节点客户端dipcli完成

### 2.1 官方钱包下载地址(目前只支持安卓): http://fir.highstreet.top/dipperwallet
安装后根据提示创建账号

### 2.2 节点客户端dipcli创建账号
- 如果没有安卓手机可以采用这种方式
- 需要linux或者mac osx操作系统，参考文档安装Dipper Network节点程序
```
浏览: https://docs.dippernetwork.com/software/how-to-install.html#%E5%AE%89%E8%A3%85%E4%B8%BB%E7%BD%91并安装主网程序
```

- 安装完毕后执行创建账号
```
dipcli keys add [account name]
# e.g.
dipcli keys add alice

Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:
{
  "name": "alice",
  "type": "local",
  "address": "dip1kku6g066hzl6pxm7ldf8ydl2wn7damy3q3awl8",
  "pubkey": "dippub1addwnpepqdfwplcx7gq8grqaeq9tqc45799jyd4qa2qgspx8r03vwk5ueppu7tknz8w",
  "mnemonic": "spirit uncle slender inform fit ignore pretty lawn link depart spike panic hen tip raw jewel armor ensure oak success again swallow wing obscure"
}

mnemonic为助记词
address为账号地址

# 根据提示输入两次密码后程序会在本地创建账号，并输出助记词，账号地址等信息，请妥善保管你的助记词等信息，防止资产丢失
```


## 3. 授权官方锁仓合约ERC20
### 3.1 打开以太坊浏览器: https://etherscan.io/
### 3.2 在以太坊浏览器搜索DIP的ERC20地址: 0x97af10D3fc7C70F67711Bf715d8397C6Da79C1Ab
### 3.3 点击Contract标签下的Write Contract标签
### 3.4 确保MetaMask安装完毕并已成功导入待映射以太坊账号
### 3.5 点击Connect to Web3,弹出的Connect a Wallet 对话框中选择MetaMask,根据提示完成连接,确保状态更新为Connected的状态
### 3.6 点击approve
### 3.7 输入官方的锁仓合约地址: 0xD14f2f1e32e2Dd4C17DFA27f1393815674e2adA2,输入要映射的数量,推荐将所有DIP ERC20做映射,点击Write,弹出MetaMask,确认交易后点击确认发送交易
### 3.8 确保交易完成


## 4. 映射
成功完成了步骤3的交易，以下步骤中的4.4, 4.5可以跳过

### 4.1 打开以太坊浏览器: https://etherscan.io/
### 4.2 在以太坊浏览器搜索DIP官方的锁仓合约地址: 0xD14f2f1e32e2Dd4C17DFA27f1393815674e2adA2
### 4.3 点击Contract标签下的Write Contract标签
### 4.4 确保MetaMask安装完毕并已成功导入待映射以太坊地址
### 4.5 点击Connect to Web3, 弹出的Connect a Wallet 对话框中选择MetaMask, 根据提示完成连接
### 4.6 点击LockToken
### 4.7 dipAddr参数输入步骤2中创建的主网地址(例如: dip1kku6g066hzl6pxm7ldf8ydl2wn7damy3q3awl8,注意: 这里是举例,请输入自己的地址),amount字段输入要映射的数量,最多不能超过步骤3中授权给合约的数量,如果超过交易会失败,点击Write,弹出MetaMask,确认交易后点击确认发送交易
### 4.8 确保交易完成
### 4.9 交易完成后,需要等待50个区块确认后映射才会被执行

## 5. 查看主网账号到账
### 5.1 步骤4的交易完成后,需要等待50个以太坊确认数映射才会最终执行,请等待至少50个区块再执行5.2以后的步骤
### 5.2 打开主网浏览器: https://explorer.dippernetwork.com/
### 5.3 在页面输入你映射的主网地址(例如: dip1kku6g066hzl6pxm7ldf8ydl2wn7damy3q3awl8),点击回车搜索你的账号详情
