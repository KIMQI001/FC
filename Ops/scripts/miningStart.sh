#!/usr/bin/env bash

go-filecoin miner create 100 --gas-price=0.001 --gas-limit=300 --repodir="${1}"
wait
go-filecoin miner set-price --gas-price=0.001 --gas-limit=1000 0.0000000000000001 2880 --repodir="${1}"
wait
go-filecoin mining start --repodir="${1}"