# 设计server端

### 概述
    1.子节点发送register，主节点返回miner
    2.子节点用miner生成builder
    3.开始监听
    4.proof任务来时，schedule进行Post锁判断，如无锁则直接进入任务
    5.若Post处于Locked状态，则将当前任务放入任务队列等待

### 参数确定
    1.测试一台机器可以计算Post的sectorID数量
    2.确定机器数量
    3.设计Register表
    
    
### 日志

    11.1 +确定GRPC 长短链接
         +Post Lock
         +AddPiece IO处理
         
### 问题
    1.sb不能自己生成postReq，可能需要构建node 的offline
    2.到时需替换node的sb指针调用：node.SectorStorage.sectorBuilder