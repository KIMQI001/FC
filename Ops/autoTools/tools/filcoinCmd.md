# filecoin command(daily)
## running Filecoin
+ rm -rf ~/.filecoin
+ go-filecoin init --devnet-user --genesisfile=https://genesis.user.kittyhawk.wtf/genesis.car --repodir=''
    
###
    //修改config
    + vim ~/.filecoin/config.json
    + i
    + Replace("address": "/ip4/127.0.0.1/tcp/3453")("address": "/ip4/127.0.0.1/tcp/3454")
    + Replace("address": "/ip4/0.0.0.0/tcp/6000")("address": "/ip4/0.0.0.0/tcp/6001")
    + Esc
    + :wq


+ go-filecoin daemon --repodir=''
+ go-filecoin config heartbeat.nickname "Pizzanode" --repodir=''
+ go-filecoin config heartbeat.beatTarget "/dns4/backend-stats.kittyhawk.wtf/tcp/8080/ipfs/QmUWmZnpZb6xFryNDeNU7KcJ1Af5oHy7fB9npU67sseEjR" --repodir=''

## wait for chain sync
+ watch -n 10 'go-filecoin show block $(go-filecoin chain head --repodir='' | head -n 1) --repodir='''

## get FIL from the Filecoin faucet
+ export WALLET_ADDR=`go-filecoin address ls --repodir=''`
+ curl -X POST -F "target=${WALLET_ADDR}" "http://user.kittyhawk.wtf:9797/tap" | cut -d" " -f4
+ go-filecoin wallet balance ${WALLET_ADDR} --repodir=''

## start mining
+ ./go-sectorbuilder/paramcache
+ go-filecoin miner create 100 --gas-price=0.001 --gas-limit=300 --repodir=''
+ go-filecoin mining start --repodir=''
+ go-filecoin miner set-price --gas-price=0.001 --gas-limit=1000 0.0000000000000001 2880 --repodir=''

## store data
+ export CID='go-filecoin client import ./hello.txt'
+ go-filecoin client list-asks --enc=json | jq
+ go-filecoin client propose-storage-deal <miner> $CID <ask> <duration> --allow
