package rpcServer

import "log"

// 筛选出合适子节点去执行任务

// register应包含：ip	; sectorID range ; TaskNum

type RegisterMap struct {
	Map 	[]Register
}

type Register struct {
	Ip		 	string
	IDRange	 	uint64
	TaskNum 	uint64
}

func Initial() *RegisterMap{
	rgsMap:=NewRgsMap().Load().Update()
	return rgsMap
}

func NewRgsMap() RegisterMap {
	return RegisterMap{}
}

func (r *RegisterMap)Load() *RegisterMap{
	// load from the disk
	return r
}

func (r *RegisterMap)Update() *RegisterMap{
	// Traversal the RegisterMap and Ping the Register IP

	// if failed,delete the Register
	return r
}

func (r *RegisterMap)NewRegiser(ip string,idr,tasNum uint64)  Register{
	return Register{ip,idr,tasNum}
}

// rpc add
func (r *RegisterMap)Add() {
	NewRg:=r.NewRegiser()
	r.Map=append(r.Map,NewRg)
	log.Printf("new register is coming ip:%s : IDR:%d -- TaskNum:%d",NewRg.Ip,NewRg.IDRange,NewRg.TaskNum)
}

func (r *RegisterMap)Delete(index int)  {
	r.Map=append(r.Map[:index],r.Map[index+1:]...)
}

func (r *RegisterMap)Select() string {
	// Traversal the map and select a num smallest Register
	min:=r.Map[0].TaskNum
	ip:=r.Map[0].Ip
	for _,x :=range r.Map{
		if x.TaskNum<min{
			min = x.TaskNum
			ip=x.Ip
		}
	}
	return ip

}