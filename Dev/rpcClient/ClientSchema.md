# 设计client端

### 概述
    1.接收子节点注册并返回miner
    2.改动原AddPiece逻辑，让其在函数调用时先执行selector，再执行rpcAddPiece
    3.改动Poller底层findMetaData，遍历所有map节点，执行rpcGetMetaData
    4.改动原Post逻辑，让其在函数调用时执行rpcPost
    
    
### 重点
    1.poller 逻辑梳理
    2.post  逻辑梳理
    