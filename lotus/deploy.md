## lotus Deploy

### Procedure

    replace /usr/local/go  by go 1.13
    sudo apt install bzr
    
    git clone https://github.com/filecoin-project/lotus.git
    cd lotus/
    vim lotus-storage-miner/main.go 
    (replace FlagStorageRepo/value to the storage location u want) 
    make
    sudo make install 
    
    rm -rf ~/.lotus ~/.lotusstorage
    lotus daemon
    
    watch lotus sync status
    lotus wallet new bls
    (height status：sync persistence complete)
    
    
go to https://lotus-metrics.kittyhawk.wtf/ check height
    
go to https://lotus-faucet.kittyhawk.wtf/ create miner
    
    lotus-storage-miner run
    lotus-storage-miner info
    lotus-storage-miner store-garbage
    

### Make a deal
    
    dd if=/dev/urandom of=TradeFile bs=100KB count=166
    lotus client import ./TradeFile
    lotus client query-ask <miner>
    lotus client deal <Data CID> <miner> 50000000 2880 
    
### Current schema
    
    use the storagegarbage.sh to store 16.8M*14400=236.25G
    lotus-storage-miner store-garbage
    //每两个16.8M*2 耗时2+min （i7-8700 12核）
    //一个小时大概1000M,一天约20G
    
### Pond solution

    进入sudo
    npm install -g npm@latest 
    rm -rf node_modules 
    npm install
    sudo npm remove csstools/normalize.css
    sudo npm install csstools/normalize.css
    
    退出sudo
    make pond
    ./pond run &