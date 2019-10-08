package clusterSchema

// 梳理sectorBuilder的数据结构及功能 并构建sb
/*
	sectorBuilder相当与一个“生产证明数据的工厂”，给入合适参数调用对应证明方法其生成对应的证明数据。
	我们的目标就是：独立出工厂，使我们可控操作，复制出n份，成为n个子节点。
*/
