package score

import (
	"tianchi/pkg/core"
	"fmt"
)

func Factory(ss ScoreStruct) int64 {
	score := int64(0)
	switch ss.CountType {
	case core.Static:
		score1 := ss.countResource()
		fmt.Println("score1: ", score1)
		score2 := ss.countInstanceGroup()
		fmt.Println("score2: ", score2)
		score = score1 + score2
	}
	return score
}