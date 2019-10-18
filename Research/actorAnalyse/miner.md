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
        2.miner的sector满的时候，自动启动seal(off chain)
        3.seal完成后,调用commitSector(on chain)
        4.如果是第一次commitSector,将会启动Post周期
        
               





### 改动
       
       1.原先ReplicaID由ProverID和SectorID合成
       接下来由block和CommD合成，便无法去提前跑PoRep
       2.SectorID将由MinerID和SectorNum合成，保证全球唯一性
       


### 参考
    https://filecoin-project.github.io/specs/#systems__filecoin_mining