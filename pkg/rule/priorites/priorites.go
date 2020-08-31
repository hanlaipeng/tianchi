package priorites

import (
	"tianchi/pkg/core"
	"tianchi/pkg/rule"
)

func Priorites(node core.SchedulerNode, podGroup core.PodGroup) int64 {

	opts := rule.Options{
		SchedApp:  podGroup.SchedApp,
		SchedNode: node,
	}
	preOpts := PrioritesOptions{
		PriOpts: opts,
	}

	return doPriorites(preOpts)
}

func doPriorites(opts PrioritesOptions) int64 {
	opts.PriOpts.Init()
	score := int64(0)
	score = score + opts.gerGPUQuotaScore()
	//fmt.Println("gpu score: ", opts.gerGPUQuotaScore())
	score = score + opts.getCPUQuotaScore()
	//fmt.Println("cpu score: ", opts.getCPUQuotaScore())
	score = score + opts.getRamQuotaScore()
	//fmt.Println("ram score: ", opts.getRamQuotaScore())
	if opts.PriOpts.GroupRule != nil {
		score = score + opts.getGroupScore()
		//fmt.Println("group score: ", opts.getGroupScore())
	}
	return score
}
