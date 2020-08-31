package core

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func InitStatic() {
	fmt.Println("start init .....")
	//读取app配置文件
	err := initApp("/Users/hanlaipeng/Desktop/tianchi/schedule.app.source")
	//err := initApp("./conf/schedule.app.source")
	Check(err)
	//读取node配置文件
	err = initNode("/Users/hanlaipeng/Desktop/tianchi/schedule.node.source")
	//err = initNode("./conf/schedule.node.source")
	Check(err)
	//读取rule配置文件
	err = initRule("/Users/hanlaipeng/Desktop/tianchi/rule.source")
	//err = initRule("./conf/rule.source")
	Check(err)
	fmt.Println("initRule finally!")
	fmt.Println("end init .....")
}

func InitDynamic() {
	fmt.Println("start init dynamic.....")
	//读取app配置文件
	err := initApp("/Users/hanlaipeng/Desktop/tianchi/reschedule.source")
	//err := initApp("./conf/reschedule.source")
	Check(err)
	fmt.Println("end init dynamic.....")
}

func initApp(path string) error {
	//读取app配置文件
	var app []App
	err := LoadJsonFile(path, &app)
	if err != nil {
		return err
	}
	//初始化队列
	for _, v := range app {
		for i := int64(0); i < v.Replicas; i++ {
			pod := Pod{
				AppName: v.AppName,
				Group: v.Group,
				Gpu: v.Gpu,
				Cpu: v.Cpu,
				Ram:v.Ram,
				Disk:v.Disk,
			}
			podLink := PodGroup{
				SchedPod: pod,
				SchedApp:v,
			}
			ActiveQ.PushPod(&podLink)
		}
	}
	fmt.Println("initApp finally!")
	return nil
}

func initNode(path string) error {
	var nl []Node
	err := LoadJsonFile(path, &nl)
	if err != nil {
		return err
	}
	//初始化节点列表
	var nodeList []SchedulerNode
	for _,v := range nl {
		node := SchedulerNode{
			SchedNode: v,
		}
		nodeList = append(nodeList, node)
	}
	NodeList.NodeList = nodeList
	fmt.Println("initNode finally!")
	return nil
}

func initRule(path string) error{
	return  LoadJsonFile(path, &RuleRes)
}

func initReschedulerNode(path string) error {
	return LoadJsonFile(path, &NodeList)
}

func LoadJsonFile(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("in LoadJsonFile ReadFile: %s", err.Error())
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("in LoadJsonFile Unmarshal: %s", err.Error())
	}
	return nil
}

func Check(err error) {
	if err != nil{
		panic(err)
	}
}

func CovertWeightsToMap() map[string]NodeResourceWeight {
	res := make(map[string]NodeResourceWeight)
	for _, v := range RuleRes.NodeResourceWeights {
		key := fmt.Sprintf("%s+%s", v.Resource, v.SmName)
		res[key] = v
	}
	return res
}

func CovertAppRuleToMap() map[string]GroupMaxInstancePerHost {
	res := make(map[string]GroupMaxInstancePerHost)
	for _, v := range RuleRes.GroupMaxInstancePerHosts {
		res[v.Group] = v
	}
	return res
}

func CovertAppGroupToMap(podList []Pod) map[string]int64 {
	res := make(map[string]int64)
	for _, v := range podList {
		if _, ok := res[v.Group]; ok {
			res[v.Group]++
		}else{
			res[v.Group] = 1
		}
	}
	return res
}

