# filecoin command(daily)
## running Filecoin
+ rm -rf $(repodir)
+ go-filecoin init --devnet-user --genesisfile=https://genesis.user.kittyhawk.wtf/genesis.car --repodir='$(repodir)'    
+ go-filecoin daemon --cmdapiaddr=/ip4/127.0.0.1/tcp/3454 --swarmlisten=/ip4/127.0.0.1/tcp/6001 --repodir='$(repodir)'
+ go-filecoin config heartbeat.nickname "Pizzanode" --repodir='$(repodir)'
+ go-filecoin config heartbeat.beatTarget "/dns4/backend-stats.kittyhawk.wtf/tcp/8080/ipfs/QmUWmZnpZb6xFryNDeNU7KcJ1Af5oHy7fB9npU67sseEjR" --repodir='$(repodir)'

## wait for chain sync
+ watch -n 10 'go-filecoin show block $(go-filecoin chain head --repodir='$(repodir)' | head -n 1) --repodir='$(repodir)''

## get FIL from the Filecoin faucet
+ export WALLET_ADDR=`go-filecoin address ls --repodir=''`
+ curl -X POST -F "target=${WALLET_ADDR}" "http://user.kittyhawk.wtf:9797/tap" | cut -d" " -f4
+ go-filecoin wallet balance ${WALLET_ADDR} --repodir=''

## start mining
+ ./go-sectorbuilder/paramcache
+ go-filecoin miner create 100 --gas-price=0.001 --gas-limit=300 --repodir='$(repodir)'
+ go-filecoin mining start --repodir='$(repodir)'
+ go-filecoin miner set-price --gas-price=0.001 --gas-limit=1000 0.0000000000000001 2880 --repodir='$(repodir)'

## store data
+ export MINER='go-filecoin mining address'
+ dd if=/dev/urandom of=TradeFile bs=1M count=244
+ export CID='go-filecoin client import ./TradeFile'
+ go-filecoin client list-asks | grep $MINER
+ go-filecoin client propose-storage-deal $MINER $CID 0 2880 --allow-duplicates
