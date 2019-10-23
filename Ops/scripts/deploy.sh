#!/usr/bin/env bash
mkdir -p ${GOPATH}/src/github.com/filecoin-project
git clone https://github.com/filecoin-project/go-filecoin.git ${GOPATH}/src/github.com/filecoin-project/go-filecoin
wait
cd $GOPATH/src/github.com/filecoin-project/go-filecoin
git fetch origin
wait
git checkout "${1}"
git submodule update --init --recursive
wait
FILECOIN_USE_PRECOMPILED_RUST_PROOFS=true go run ./build deps
wait
go run ./build build
wait
cp go-filecoin $GOPATH/bin

