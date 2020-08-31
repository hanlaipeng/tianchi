package priorites

import "tianchi/pkg/rule"

type PrioritesOptions struct {
	PriOpts rule.Options
}

const (
	BaseScoreQuota       = int64(1000)
	BaseScoreCompactness = int64(100)
)

//base score 100
func (p *PrioritesOptions) gerGPUQuotaScore() int64 {
	subGpuQuota := p.PriOpts.NodeNow.Gpu - p.PriOpts.SchedApp.Gpu
	score := CountScoreQuota(subGpuQuota)
	return score
}

func (p *PrioritesOptions) getCPUQuotaScore() int64 {
	subCpuQuota := p.PriOpts.NodeNow.Cpu - p.PriOpts.SchedApp.Cpu
	score := CountScoreQuota(subCpuQuota)
	return score
}

func (p *PrioritesOptions) getRamQuotaScore() int64 {
	subRamQuota := p.PriOpts.NodeNow.Ram - p.PriOpts.SchedApp.Ram
	score := CountScoreQuota(subRamQuota)
	return score
}

func (p *PrioritesOptions) getGroupScore() int64 {
	insHost := int64(0)
	for _, v := range p.PriOpts.SchedNode.RunPod {
		if v.Group == p.PriOpts.SchedApp.Group {
			insHost++
		}
	}
	//堆叠需求
	if p.PriOpts.GroupRule.Compactness == true {
		return BaseScoreCompactness * insHost
	}
	//松散需求
	if p.PriOpts.GroupRule.Compactness == false {
		return BaseScoreCompactness * (p.PriOpts.GroupRule.MaxInstancePerHost - insHost)
	}
	return 0
}

func CountScoreQuota(resNum int64) int64 {
	if resNum == 0 {
		return BaseScoreQuota
	}
	return BaseScoreQuota / resNum
}
