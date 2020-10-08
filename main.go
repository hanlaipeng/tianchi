package main

import (
	"time"
	"tianchi/pkg/core"
	"tianchi/pkg/static"
	"fmt"
)

func main() {
	t1 := time.Now()
	core.InitStatic()
	static.StaticScheduler()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).String())
}
