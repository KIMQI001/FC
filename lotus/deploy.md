## lotus Deploy

### Procedure

    replace /usr/local/go  by go 1.13
    sudo apt install bzr
    
    git clone https://github.com/filecoin-project/lotus.git
    cd lotus/
    make
    sudo make install 
    
    rm -rf ~/.lotus ~/.lotusstorage
    lotus daemon
    
    watch lotus sync status
    lotus wallet new bls
    
    
go to https://lotus-metrics.kittyhawk.wtf/ check height

go to https://lotus-faucet.kittyhawk.wtf/ create miner
    
    lotus-storage-miner run
    lotus-storage-miner info
    lotus-storage-miner store-garbage
    

### Make a deal
    
    dd if=/dev/urandom of=TradeFile bs=1M count=244
    lotus client import ./TradeFile
    lotus client query-ask <miner>
    lotus client deal <Data CID> <miner> 2210565980160 2880 
### Pond solution

    进入sudo
    npm install -g npm@latest 
    rm -rf node_modules 
    npm install
    sudo npm remove csstools/normalize.css
    sudo npm install csstools/normalize.css
    
    退出sudo
    make pond