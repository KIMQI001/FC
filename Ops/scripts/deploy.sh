#!/usr/bin/env bash
export http_proxy="http://127.0.0.1:12333"
export https_proxy="http://127.0.0.1:12333"

mkdir -p ${GOPATH}/src/github.com/filecoin-project
git clone https://github.com/filecoin-project/go-filecoin.git ${GOPATH}/src/github.com/filecoin-project/go-filecoin
wait
cd $GOPATH/src/github.com/filecoin-project/go-filecoin
git fetch origin
wait
git checkout -b "${1}"
git submodule update --init --recursive
wait
FILECOIN_USE_PRECOMPILED_RUST_PROOFS=true go run ./build deps
wait
go run ./build build
wait
cp go-filecoin $GOPATH/bin

