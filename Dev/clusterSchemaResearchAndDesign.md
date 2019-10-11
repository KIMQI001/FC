# filecoin集群系统设计与研究
## 一.主要问题点
### 1.什么是sectorbuilder
+ 梳理sectorbuilder的数据结构及功能
+ 梳理与sb功能相关数据结构
+ 构建sb
+ 构建虚拟密封过程
+ 构建虚拟Post过程

### 2.确定远程数据交换内容（服务参数数据结构）
+ AddPiece
    1. PieceRef cid /pieceSize uint64 /pieceReader io.Reader
    2. sectorID uint64
+ GetSectorSealingStatusByID
    1. SectorSealingStatus

### 3.如何从主节点分发任务并回收结果
+ 修改主节点rustSectorBuilder
    1. 修改 AddPiece
    2. 修改 findSealedSectorMetadata
    3. 修改 ReadPieceFromSealedSector及其它（次要）

### 4.AddPiece交互流程图

![Image text](https://github.com/KIMQI001/FC/blob/master/img/RpcRep.png)

## 二.主要技术点
### 1.IO流的本质是什么，如何网络传输一个IO流

### 2.选择远程数据交换格式

