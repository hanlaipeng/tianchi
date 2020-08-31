package core

var ActiveQ = &SchedulerPodQueue{}

type PodGroup struct {
	SchedPod Pod
	SchedApp App
}

type PodOperate interface {
	NextPod() *PodGroup
	PushPod(podGroup *PodGroup)
}

type SchedulerPodQueue struct {
	PodData *PodGroup
	Next    *SchedulerPodQueue
}

//链表头出法，选择最新的pod出队列
func (spq *SchedulerPodQueue) NextPod() *PodGroup {
	if spq.Next == nil {
		return nil
	}
	nextPod := spq.Next.PodData
	spq.Next = spq.Next.Next
	return nextPod
}

//链表尾插法， 入队列
func (spq *SchedulerPodQueue) PushPod(podGroup *PodGroup) {
	p := spq
	for p.Next != nil {
		p = p.Next
	}
	podQ := &SchedulerPodQueue{
		PodData: podGroup,
	}
	p.Next = podQ
}
