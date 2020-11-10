# dip-bridge
以太坊DIP ERC20映射Dipper Network主网中继服务

## how to use
### 1. Prepare ENV: 
- ethereum geth RPC
- redis server
- dipper network Node
- dipper network RPC

### 2. deploy dipManager contract on ethereum
```
# deploy ./contracts/eth_land_contracts/dip_manager.sol
```

### 3. deploy dipManager contract on dipper network
```
# deploy ./contracts/dip_land_contracts/dip_manager.sol
bash ./contracts/dip_land_contracts/deploy.sh [dip addr1] [dip addr2]

# dip addr1: the singer of the transaction
# dip addr2: the admin of the dipManager contract who can do mintToken
```

### 4. compile
```
make install
```

### 5. config
- config ./config/keystore
```
# save the output of the commond below as keystore file
dipcli keys export [your dipper network account address]
```
- config ./config/dip_sdk.yaml
- config ./config/bridge.yaml

### 6. run
```
dip-bridge ./config/bridge.yaml
```

## how to do ERC20 mapping
please read UserGuide.md