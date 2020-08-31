package scheduler

import (
	"fmt"
	"sync"
	"tianchi/pkg/core"
	"tianchi/pkg/rule/prediactes"
	"tianchi/pkg/rule/priorites"
	"tianchi/pkg/util"
	"encoding/json"
)

type schedulerOptions struct {
	Score       int64
	PodGroup    core.PodGroup
	SuggestHost core.SchedulerNode
}

func Run() bool {
	sc := schedulerOptions{}
	return sc.schedulerOne()
}

func (sc *schedulerOptions) schedulerOne() bool {
	podGroup := core.ActiveQ.NextPod()
	if podGroup == nil {
		return false
	}
	//fmt.Println("start scheduler, pod: ", podGroup)
	sc.PodGroup = *podGroup
	//队列中没有可调度pod, 结束流程
	err := sc.scheduler()
	if err != nil {
		//TODO-- 调度失败，重新加入待调度队列
		//core.ActiveQ.PushPod(&sc.PodGroup)
		var data []byte
		json.Unmarshal(data, &sc.PodGroup)
		errPodOutput := util.JsonOptions{
			OutPath: "./output/file_pod.json",
			Data: data,
		}
		errPodOutput.WriteJson()
		return true
	}
	sc.DistributeCpuID()
	nodeList := sc.bind()
	core.NodeList = core.NodeLists{
		NodeList: nodeList,
	}
	//fmt.Println("scheduler result: ", core.NodeList)

	return true

}

func (sc *schedulerOptions) scheduler() error {
	var suggestHost core.SchedulerNode
	nodeList := core.NodeList.GetNodeList()
	filterNodeList := prediactes.Predicates(nodeList, sc.PodGroup)
	//fmt.Println("filterNodeList: ", filterNodeList, "scheduler pod: ", sc.PodGroup)

	if len(filterNodeList) == 1 {
		sc.SuggestHost = filterNodeList[0]
		return nil
	}

	if len(filterNodeList) == 0 {
		return fmt.Errorf("have no node for pod")
	}

	wg := sync.WaitGroup{}
	wg.Add(len(filterNodeList))
	scoreCh := make(chan schedulerOptions, len(filterNodeList))
	for _, v := range filterNodeList {
		go func(node core.SchedulerNode) {
			defer wg.Done()
			score := priorites.Priorites(node, sc.PodGroup)
			sched := schedulerOptions{
				Score:       score,
				PodGroup:    sc.PodGroup,
				SuggestHost: node,
			}
			scoreCh <- sched
		}(v)
	}
	wg.Wait()
	close(scoreCh)

	maxScore := int64(0)
	maxPod := int(0)
	for sched := range scoreCh {
		//fmt.Println("score:", sched)
		if sched.Score > maxScore {
			maxScore = sched.Score
			suggestHost = sched.SuggestHost
			maxPod = len(sched.SuggestHost.RunPod)
		}
		if sched.Score == maxScore {
			if len(sched.SuggestHost.RunPod) > maxPod {
				maxScore = sched.Score
				suggestHost = sched.SuggestHost
				maxPod = len(sched.SuggestHost.RunPod)
			}
		}
	}
	sc.Score = maxScore
	sc.SuggestHost = suggestHost

	return nil

}

func (sc *schedulerOptions)DistributeCpuID() {
	//cpuID
	pod := sc.PodGroup.SchedPod
	cpuArr := make([]int64, sc.SuggestHost.SchedNode.Cpu)
	for _, v := range sc.SuggestHost.RunPod {
		for _, item := range v.CpuIDs {
			cpuArr[item] = 1
		}
	}
	count := pod.Cpu
	for k, v := range cpuArr {
		if v == 0 && count > 0 {
			pod.CpuIDs = append(pod.CpuIDs, k)
			count--
		}
	}

	sc.PodGroup.SchedPod = pod
}

func (sc *schedulerOptions) bind() []core.SchedulerNode{
	//fmt.Println("start to bind ...")
	nodeList := core.NodeList.GetNodeList()
	for k, v := range nodeList {
		if v.SchedNode.Sn == sc.SuggestHost.SchedNode.Sn {
			runPods := v.RunPod
			runPods = append(runPods, sc.PodGroup.SchedPod)
			v.RunPod = runPods
			nodeList[k] = v
			return nodeList
		}
	}
	//fmt.Println("end to bind ...")
	return nodeList
}
