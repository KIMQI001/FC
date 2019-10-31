# FC的部署
    
    git clone https://github.com/filecoin-project/go-filecoin.git
    cd $GOPATH/src/github.com/filecoin-project/go-filecoin
    git fetch origin
    git checkout -b 0.5.7
    git submodule update --init --recursive
    FILECOIN_USE_PRECOMPILED_RUST_PROOFS=true go run ./build deps
    go run ./build build
    cp go-filecoin $GOPATH/bin


## 利用autoTools部署

## 利用Docker部署

## 利用K8s的部署集群研究