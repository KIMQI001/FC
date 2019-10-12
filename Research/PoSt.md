## 如何启动分发PoSt任务

### PoSt逻辑交互流程
+ 业务层逻辑（Go层）
    + 启动（OnNewHeaviestTipSet）
        + 获得commits
        + 收集inputs
        + 获得ProvingWindow
        + 确定没有在这个period提交过Post
        + ProvingWindow+challengeDelayRounds<确保当前高度<provingWindowEnd
        + miner.submitPost->sm.prover.CalculatePost(inputs)
        + 
    + 上链（minerActor.SubmitPost）
        + 获取到当前链高度
        + 计算 nextProvingPeriodEnd
        + 判断是否未超时(当前高度<nextProvingPeriodEnd)
        + 计算 provingWindowStart（intime版或late版）
        + 判断当前高度是否在 provingWindowStart 之后
        + 生成 ChallengeSeed
        + 收集当前状态下 sectorInfos并sort化
        + 构造verifyPostRequest
        + 执行VerifyPoSt
        + 改变state的Power，commits，provingSet，NextDoneSet
+ 算法层逻辑（Rust层）
### 方案猜想一
     // 业务层逻辑（Go层）
     + 主节点submitPost启动时
        1.主节点删除收集commits逻辑，只生成seed
        2.向子节点发送seed
        3.子节点接受seed开始启动Post
        4.子节点计算完毕返回
        5.主节点接受结果并合成
     
     + 难点：
        1.删除收集commits逻辑
        2.如何手动收集到子节点commits
        3.如何合并各个子节点返回的PostResponse
        
        
### 方案猜想二
    // 业务层逻辑（Go层）
    + 主节点submitPost启动时
        1.主节点收集commits并产生seed
        2.主节点不执行post算法
        3.主节点直接手动拆分commits成不同的inputs，然后包装成request分发给对应的子节点
        4.子节点根据request启动Post算法并生成结果
        5.子节点返回结果并在主节点处合并
        
    + 难点：
        1.拆分commits成不同的inputs
        2.合并Response
        
### 方案猜想三
    // 算法层逻辑（Rust层）
    + 主节点submitPost启动时
        1.完成cgo调用至Rust-PoSt
        2.分拆成micro-server，调用并返回结果
        
    + 难点：
        1.Rust-PoSt逻辑梳理
        2.micro-server的拆分