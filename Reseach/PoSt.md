## 如何启动分发PoSt任务

### PoSt逻辑交互流程

### 方案猜想一
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
    + 主节点submitPost启动时
        1.主节点收集commits并产生seed
        2.主节点不执行post算法
        3.主节点直接手动拆分commits成不同的inputs，然后包装成request分发给对应的子节点
        4.子节点根据request启动Post算法并生成结果
        5.子节点返回结果并在主节点处合并
        
    + 难点：
        1.拆分commits成不同的inputs
        2.合并Response
        