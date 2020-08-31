package adjust

import (
	"tianchi/pkg/core"
	"sort"
	"fmt"
)

type AdjustData struct {
	ReNodeList core.NodeLists
	NullNodeList []core.SchedulerNode
}

func (ad *AdjustData) DepartNullNodeForReScheduler() {
	sort.Sort(ad.ReNodeList)
	for k, v := range ad.ReNodeList.NodeList {
		fmt.Println("node name:", v.SchedNode.SmName, "pod num:", len(v.RunPod))
		if len(v.RunPod) > 0 {
			ad.ReNodeList.NodeList = ad.ReNodeList.NodeList[k:]
			break
		}
		ad.NullNodeList = append(ad.NullNodeList, v)
	}
}

//func (ad *AdjustData) SelectPodForReScheduler() {
//	for _, v := range ad.ReNodeList.GetNodeList() {
//
//	}
//}


