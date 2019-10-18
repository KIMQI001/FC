## miner功能整理

### storageMiner
      
 
+ minerActor分为worker和owner，这两个角色拥有不同的address
        
        worker职责：sign blocks create by this miner
                   submit proofs
                   commit new sectors
                   other day to day activities
                   
        owner职责： create miner
                   pay the collateral
                   accept the block rewards

+ Storage Mining Cycle

        在每一个height检查：
        1.自身是否处于Post周期（Start/InProcess/End）
        2.自身是否获得一个winningTicket去申请挖掘下一个块
              
+ deal流程：
        
        1.make deal(on chain)
        2.miner接收数据放入sector，sector满的时候，自动启动seal,即执行复制证明(off chain)
        3.seal完成后,调用commitSector到链上(on chain)
        4.如果是第一次commitSector,将会启动Post周期,开始进行第一次Post计算
        5.计算完Post后，调用submitPost到链上(on chain)
        6.从第一次启动Post证明的时刻开始，此后每隔一个Post周期(生成300个块的时间)提交一次Post.
        
        
+ 注意事项：
        
        1.storageMiner维护一个PorvingSet（需要Post的sectors集合）和一个FaultsSet（Post失败的集合）
        2.commitSector后，会将此sector放入下一个Post周期的ProvingSet，即如果在当前Post周期内commit了新的sector，不会计入当前Post计算
        3.更新Power（全网存储权重）是在submitPost执行成功后
        4.计算Post后，会返回一个proof和一个Faults，Faults包含计算Post失败的sector，会被放入FaultsSet进行重新计算。               

### retrieval miner
    
    
    参考 https://filecoin-project.github.io/specs/#systems__filecoin_markets__retrieval_market



### 改动
       
       1.原先ReplicaID由ProverID和SectorID合成
       接下来由block和CommD合成，便无法去提前跑PoRep
       2.SectorID将由MinerID和SectorNum合成，保证全球唯一性
       


### 参考
    https://filecoin-project.github.io/specs/#systems__filecoin_mining