package static

import (
	"tianchi/pkg/core"
	"tianchi/pkg/scheduler"
	"encoding/json"
	"tianchi/pkg/util"
	"tianchi/pkg/score"
	"fmt"
)

//静态布局
func StaticScheduler() {
	flag := true
	for flag {
		flag = scheduler.Run()
	}

	//adjustList := adjust.AdjustData{
	//	ReNodeList: core.NodeList,
	//}
	//
	//adjustList.SelectPodForReScheduler()

	data, _ := json.Marshal(&core.NodeList)
	staticResult := util.JsonOptions{
		OutPath: "./output/schedule.result",
		Data: data,
	}
	staticResult.WriteJson()
	ss := score.ScoreStruct{
		CountType: core.Static,
		DataList: core.NodeList.GetNodeList(),
	}
	scoreRes := score.Factory(ss)
	fmt.Println("final score: ", scoreRes)
}
