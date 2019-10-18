# gen
./gengen -config=/home/kehu/file/hknet/0.5.6/bin/gen/setup.json -out-json=gen.json -test-proofs-mode=true -json=true -out-car=hkcloud.car
## setup.json
```
{
  "keys": 5,
  "preAlloc": [
    "1000000000000",
    "1000000000000",
    "1000000000000",
    "1000000000000",
    "1000000000000"
  ],
  "miners": [{
    "owner": 0,
    "numCommittedSectors": 1
  }],
  "network": "alpha2"
}
```
## gen.json
```
{
  "Keys": [
    {
      "privateKey": "NXMFg7cZi20u5Qvlyk6OsqP8nA8Dajqb3H13C9u4CT8=",
      "curve": "secp256k1"
    },
    {
      "privateKey": "3DWmPZSJGLL+aizkgsUpTLHQzUBqjZU9AFfC3rAFjsY=",
      "curve": "secp256k1"
    },
    {
      "privateKey": "LXYqfeWCsj3p0tLcjU5+9DDJzulDwQ6KHDR1GM3Ni+Y=",
      "curve": "secp256k1"
    },
    {
      "privateKey": "AC6AtzCNFzgdHIHbS3cik2fOsEsPpYrV3FtXGYYgCys=",
      "curve": "secp256k1"
    },
    {
      "privateKey": "OvcqUB0akfbEdPB0TEItzu1k5HaEX+mlEXbAPtkmIE8=",
      "curve": "secp256k1"
    }
  ],
  "Miners": [
    {
      "Owner": 0,
      "Address": "t23ubhfrqg46pgsypk777c3efwj7ajnrgco5iucyi",
      "Power": "1024"
    }
  ],
  "GenesisCid": {
    "/": "bafy2bzaceanb4vbxzlz4nxdqbyypjibl5zsi5fuw5il67louv32tmtuapgqd6"
  }
}
```
# car
./genesis-file-server -genesis-file-path=/home/kehu/file/hknet/0.5.6/bin/gen/hkcloud.car -port=8092 >genesis-file-server.log 2>&1 &
# init
./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --with-miner=t23ubhfrqg46pgsypk777c3efwj7ajnrgco5iucyi --repodir="/home/kehu/file/hknet/0.5.6/boots/9000"

./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --repodir="/home/kehu/file/hknet/0.5.6/boots/9001"
./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --repodir="/home/kehu/file/hknet/0.5.6/boots/9002"
./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --repodir="/home/kehu/file/hknet/0.5.6/boots/9003"
./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --repodir="/home/kehu/file/hknet/0.5.6/boots/9004"
./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --repodir="/home/kehu/file/hknet/0.5.6/boots/9005"

# config backup
cp /home/kehu/file/hknet/0.5.6/boots/9000/config.json /home/kehu/file/hknet/0.5.6/backup/9000-config.json
cp /home/kehu/file/hknet/0.5.6/boots/9001/config.json /home/kehu/file/hknet/0.5.6/backup/9001-config.json
cp /home/kehu/file/hknet/0.5.6/boots/9002/config.json /home/kehu/file/hknet/0.5.6/backup/9002-config.json
cp /home/kehu/file/hknet/0.5.6/boots/9003/config.json /home/kehu/file/hknet/0.5.6/backup/9003-config.json
cp /home/kehu/file/hknet/0.5.6/boots/9004/config.json /home/kehu/file/hknet/0.5.6/backup/9004-config.json
cp /home/kehu/file/hknet/0.5.6/boots/9005/config.json /home/kehu/file/hknet/0.5.6/backup/9005-config.json

cp /home/kehu/file/hknet/0.5.6/backup/9000-config.json /home/kehu/file/hknet/0.5.6/boots/9000/config.json
cp /home/kehu/file/hknet/0.5.6/backup/9001-config.json /home/kehu/file/hknet/0.5.6/boots/9001/config.json
cp /home/kehu/file/hknet/0.5.6/backup/9002-config.json /home/kehu/file/hknet/0.5.6/boots/9002/config.json
cp /home/kehu/file/hknet/0.5.6/backup/9003-config.json /home/kehu/file/hknet/0.5.6/boots/9003/config.json
cp /home/kehu/file/hknet/0.5.6/backup/9004-config.json /home/kehu/file/hknet/0.5.6/boots/9004/config.json
cp /home/kehu/file/hknet/0.5.6/backup/9005-config.json /home/kehu/file/hknet/0.5.6/boots/9005/config.json

# daemon
./go-filecoin daemon --repodir="/home/kehu/file/hknet/0.5.6/boots/9000" >9000-run.log 2>&1 &
./go-filecoin daemon --repodir="/home/kehu/file/hknet/0.5.6/boots/9001" >9001-run.log 2>&1 &
./go-filecoin daemon --repodir="/home/kehu/file/hknet/0.5.6/boots/9002" >9002-run.log 2>&1 &
./go-filecoin daemon --repodir="/home/kehu/file/hknet/0.5.6/boots/9003" >9003-run.log 2>&1 &
./go-filecoin daemon --repodir="/home/kehu/file/hknet/0.5.6/boots/9004" >9004-run.log 2>&1 &
./go-filecoin daemon --repodir="/home/kehu/file/hknet/0.5.6/boots/9005" >9005-run.log 2>&1 &

# bootstrap
/ip4/116.31.96.179/tcp/9000/ipfs/QmNfa31tWFJfoRTd6rGpnbvMZm6jqJubPbLfB3CdLuSMaQ
/ip4/116.31.96.179/tcp/9001/ipfs/QmSrLXzedXNn4WqP3k5EHgmviPD5yzi8KzZTavNPSwR2wf
/ip4/116.31.96.179/tcp/9002/ipfs/QmbLT8xtEqA9VHMRgPjWnqVy3Nsi8ekN5yZu5iy4vaoQKk
/ip4/116.31.96.179/tcp/9003/ipfs/QmVH75zZbv6eDCc8zy65YKKhk4uBsNiroJJd3ycC9Bh7PK
/ip4/116.31.96.179/tcp/9004/ipfs/Qmb9gZvd1reUQbZtr1QAzKJCLiESzDpM62Hc2csC7nhyoY
/ip4/116.31.96.179/tcp/9005/ipfs/QmQmFnSV18ei3va4uVj4kFYeWJP1ek1hTJ2NyMRYEQ9ij2

# boots import
./go-filecoin wallet import /home/kehu/file/hknet/0.5.6/bin/gen/0.key --repodir="/home/kehu/file/hknet/0.5.6/boots/9000"
## result t1dvwn4stcb4vwvw45zttnj5pzuljmd3nu2junpty

# update peerid
./go-filecoin miner update-peerid t23ubhfrqg46pgsypk777c3efwj7ajnrgco5iucyi QmNfa31tWFJfoRTd6rGpnbvMZm6jqJubPbLfB3CdLuSMaQ --from=t1dvwn4stcb4vwvw45zttnj5pzuljmd3nu2junpty --gas-price=300 --gas-limit=300 --repodir="/home/kehu/file/hknet/0.5.6/boots/9000"
## result bafy2bzacea3loultygracnwc4s5y3p5crxejoissal47erhxm2iwbb6n5oqmw

# mining start
./go-filecoin mining start --repodir="/home/kehu/file/hknet/0.5.6/boots/9000"

# stats
npm run dev >backend.log 2>&1 &
```
修改filecoin-network-stats/frontend/src/components/GlobalNav.tsx 
Block Explorer、Faucet 超链接地址修改
```
npm run dev >frontend.log 2>&1 &

# explorer
```bash
yarn install
HOST=116.31.96.179 PORT=8082 REACT_APP_API_URL=http://116.31.96.179:3453 yarn start >explorer.log 2>&1 &
```
# faucet
./faucet -faucet-val=10000 -fil-api=116.31.96.179:3453 -fil-wallet=t1dvwn4stcb4vwvw45zttnj5pzuljmd3nu2junpty >faucet.log 2>&1 &

# localhost
./go-filecoin init --genesisfile=http://116.31.96.179:8092/genesis.car --repodir="/home/sywyn/file/0.5.6/miners/9000"

