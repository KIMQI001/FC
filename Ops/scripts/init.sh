#!/usr/bin/env bash
# This is the fileCoin deploy script

rm -rf "${1}"
echo "start init"
go-filecoin init \
  --devnet-user \
  --repodir="${1}" \
  --sectordir="${2}" \
  --genesisfile=https://genesis.user.kittyhawk.wtf/genesis.car

go-filecoin daemon \
  --cmdapiaddr=/ip4/127.0.0.1/tcp/"${3}" \
  --swarmlisten=/ip4/127.0.0.1/tcp/"${4}" \
  --repodir="${1}" &

sleep 5;

echo "nickname"
go-filecoin config heartbeat.nickname "${5}" --repodir="${1}"

echo "heartbeat"
go-filecoin config heartbeat.beatTarget "/dns4/backend-stats.kittyhawk.wtf/tcp/8080/ipfs/QmUWmZnpZb6xFryNDeNU7KcJ1Af5oHy7fB9npU67sseEjR" \
  --repodir="${1}"