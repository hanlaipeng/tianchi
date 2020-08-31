package score

import (
	"tianchi/pkg/core"
	"fmt"
)

type ScoreStruct struct {
	CountType  string
	DataList   []core.SchedulerNode
}

type ScoreAlgorithm interface {
	countResource() int64
	countInstanceGroup() int64
}

func (ss *ScoreStruct)countResource() int64 {
	score := int64(0)
	nodeWeights := make(map[string]core.NodeResourceWeight)
	nodeWeights = core.CovertWeightsToMap()
	for _, v := range ss.DataList {
		if len(v.RunPod) > 0 {
			//GPU
			key := fmt.Sprintf("GPU+%s", v.SchedNode.SmName)
			if _, ok := nodeWeights[key]; ok {
				score = score + nodeWeights[key].Weight * v.SchedNode.Gpu
			}else{
				score = score + nodeWeights["GPU+"].Weight * v.SchedNode.Gpu
			}
			//CPU
			key = fmt.Sprintf("CPU+%s", v.SchedNode.SmName)
			if _, ok := nodeWeights[key]; ok {
				score = score + nodeWeights[key].Weight * v.SchedNode.Cpu
			}else{
				score = score + nodeWeights["CPU+"].Weight * v.SchedNode.Cpu
			}
			//RAM
			key = fmt.Sprintf("RAM+%s", v.SchedNode.SmName)
			if _, ok := nodeWeights[key]; ok {
				score = score + nodeWeights[key].Weight * v.SchedNode.Ram
			}else{
				score = score + nodeWeights["RAM+"].Weight * v.SchedNode.Ram
			}
		}
	}
	return score
}

func (ss *ScoreStruct)countInstanceGroup() int64 {
	score := int64(0)
	appRule := core.CovertAppRuleToMap()
	for _, v := range ss.DataList {
		podMap := core.CovertAppGroupToMap(v.RunPod)
		for k, num := range podMap {
			if _, ok := appRule[k]; ok {
				if appRule[k].MaxInstancePerHost > 1 && appRule[k].Compactness == false {
					score = score + num - 1
				}
				if appRule[k].Compactness == true {
					score = score - num + 1
				}
			}
		}
	}

	return score
}

