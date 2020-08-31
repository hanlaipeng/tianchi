package prediactes

import (
	"tianchi/pkg/core"
	"tianchi/pkg/rule"
)

func Predicates(nodeList []core.SchedulerNode, podGroup core.PodGroup) []core.SchedulerNode {
	var filterNode []core.SchedulerNode
	for _, v := range nodeList {
		opts := rule.Options{
			SchedApp:  podGroup.SchedApp,
			SchedNode: v,
		}
		preOpts := PredicateOptions{
			PreOpts: opts,
		}
		if doPredicates(preOpts) {
			filterNode = append(filterNode, v)
		}
	}

	return filterNode
}

func doPredicates(opts PredicateOptions) bool {
	opts.PreOpts.Init()
	if !opts.checkGPUQuota() {
		return false
	}
	if !opts.checkCPUQuota() {
		return false
	}
	if !opts.checkDiskQuota() {
		return false
	}
	if !opts.checkRamQuota() {
		return false
	}
	if !opts.checkEniQuota() {
		return false
	}
	if !opts.checkDefaultMaxInstancePerHost() {
		return false
	}
	if !opts.checkTimeLimit() {
		return false
	}
	if opts.PreOpts.GroupRule != nil {
		if !opts.checkGroupRule() {
			return false
		}
	}
	if len(opts.PreOpts.AppRule) > 0 {
		if !opts.checkAppRule() {
			return false
		}
	}

	return true
}
