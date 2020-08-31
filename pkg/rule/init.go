package rule

import (
	"tianchi/pkg/core"
)

type Options struct {
	SchedApp  core.App
	SchedNode core.SchedulerNode
	NodeNow   core.Node
	GroupRule *core.GroupMaxInstancePerHost
	AppRule   []core.ReplicasMaxInstancePerHost
}

func (p *Options) Init() {
	node := p.SchedNode.SchedNode
	for _, v := range p.SchedNode.RunPod {
		node.Gpu = node.Gpu - v.Gpu
		node.Cpu = node.Cpu - v.Cpu
		node.Disk = node.Disk - v.Disk
		node.Ram = node.Ram - v.Ram
		node.Eni = node.Eni - 1
	}
	p.NodeNow = node

	for _, v := range core.RuleRes.GroupMaxInstancePerHosts {
		if v.Group == p.SchedApp.Group {
			ruleGroup := v
			p.GroupRule = &ruleGroup
		}
	}

	for _, v := range core.RuleRes.ReplicasMaxInstancePerHosts {
		switch v.Restrain {
		case "le":
			if p.SchedApp.Replicas <= v.Replicas {
				p.AppRule = append(p.AppRule, v)
			}
		case "ge":
			if p.SchedApp.Replicas >= v.Replicas {
				p.AppRule = append(p.AppRule, v)
			}
		}
	}
}
