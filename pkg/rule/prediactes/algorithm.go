package prediactes

import (
	"tianchi/pkg/rule"
	"tianchi/pkg/core"
)

type PredicateOptions struct {
	PreOpts rule.Options
}

func (p *PredicateOptions) checkGPUQuota() bool {
	if p.PreOpts.NodeNow.Gpu < p.PreOpts.SchedApp.Gpu {
		return false
	}
	return true
}

func (p *PredicateOptions) checkCPUQuota() bool {
	if p.PreOpts.NodeNow.Cpu < p.PreOpts.SchedApp.Cpu {
		return false
	}
	return true
}

func (p *PredicateOptions) checkDiskQuota() bool {
	if p.PreOpts.NodeNow.Disk < p.PreOpts.SchedApp.Disk {
		return false
	}
	return true
}

func (p *PredicateOptions) checkRamQuota() bool {
	if p.PreOpts.NodeNow.Ram < p.PreOpts.SchedApp.Ram {
		return false
	}
	return true
}

func (p *PredicateOptions) checkEniQuota() bool {
	if p.PreOpts.NodeNow.Eni <= 0 {
		return false
	}
	return true
}

func (p *PredicateOptions) checkDefaultMaxInstancePerHost() bool {
	insHost := int64(0)
	for _, v := range p.PreOpts.SchedNode.RunPod {
		if v.Group == p.PreOpts.SchedApp.Group {
			insHost++
		}
	}
	if insHost < core.RuleRes.DefaultMaxInstancePerHost {
		return true
	}
	return false
}

func (p *PredicateOptions) checkTimeLimit() bool {
	return true
}

func (p *PredicateOptions) checkGroupRule() bool {
	insHost := int64(0)
	for _, v := range p.PreOpts.SchedNode.RunPod {
		if v.Group == p.PreOpts.SchedApp.Group {
			insHost++
		}
	}
	if insHost < p.PreOpts.GroupRule.MaxInstancePerHost {
		return true
	}
	return false
}

func (p *PredicateOptions) checkAppRule() bool {
	insHost := int64(0)

	for _, v := range p.PreOpts.SchedNode.RunPod {
		if v.Group == p.PreOpts.SchedApp.Group {
			insHost++
		}
	}

	for _, v := range p.PreOpts.AppRule {
		if v.MaxInstancePerHost <= insHost {
			return false
		}
	}

	return true
}
