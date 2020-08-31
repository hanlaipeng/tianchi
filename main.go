package main

import (
	"regexp"
	"fmt"
)

func main() {
	//t1 := time.Now()
	//core.InitStatic()
	//static.StaticScheduler()
	//t2 := time.Now()
	//fmt.Println(t2.Sub(t1).String())

	name := "aaaaa中文"
	//r, _ := regexp.Compile("[^\u4e00-\u9fa5]+")
	//matchString := r.FindString(name)
	//
	//if name != matchString {
	//	fmt.Println("name must not consist of chinese characters")
	//}

	match, err := regexp.MatchString("[^\u4e00-\u9fa5]+", name)
	if err != nil {
		fmt.Println(err)
	}
	if !match {
		fmt.Println("name must not consist of chinese characters")
	}
}
