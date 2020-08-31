package core

var NodeList NodeLists

type NodeLists struct {
	NodeList []SchedulerNode
}

type SchedulerNode struct {
	SchedNode Node  `json:"node"`
	RunPod    []Pod `json:"pods"`
}

type NodeOperate interface {
	AddRunPod(runPod Pod)
	PopRunPod(idx int)
	GetNodeList() []SchedulerNode
}

//添加一个调度成功的pod
func (sn *SchedulerNode) AddRunPod(runPod Pod) {
	sn.RunPod = append(sn.RunPod, runPod)
}

//删除一个需要移动的pod
func (sn *SchedulerNode) PopRunPod(idx int) {
	if idx < 0 {
		return
	}
	sn.RunPod = append(sn.RunPod[:idx], sn.RunPod[idx+1:]...)
}

//获取node列表
func (n *NodeLists) GetNodeList() []SchedulerNode {
	return n.NodeList
}


//sort by len(pods)
func (ns NodeLists) Len() int {
	return len(ns.NodeList)
}

func (ns NodeLists) Less(i, j int) bool {
	return len(ns.NodeList[i].RunPod) < len(ns.NodeList[j].RunPod)
}

func (ns NodeLists) Swap(i, j int) {
	ns.NodeList[i], ns.NodeList[j] = ns.NodeList[j], ns.NodeList[i]
}
