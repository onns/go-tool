package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/onns/go-tool/common"
	config_load "github.com/onns/go-tool/config-load"
)

/*
@Time : 2022/5/8 11:13
@Author : onns
@File : cal-holiday/main.go
*/

func main() {
	var (
		b   []byte
		err error
	)
	cfg := []*Holiday{}
	b, err = config_load.LoadConfig("")
	if err != nil {
		log.Printf("%v", err)
	}
	if err = json.Unmarshal(b, &cfg); err != nil {
		log.Printf("%v", err)
	}
	resList := make([]*common.MemDays, 0)
	for _,holiday :=range cfg {
		holiday.Format()
		var tempResList []*common.MemDays
		tempResList,err = holiday.Cal.Output(holiday.Start,holiday.End)
		if err != nil {
			panic(err)
		}
		resList = append(resList, tempResList...)
	}
	bs, _ := json.MarshalIndent(resList,"","  ")
	err = ioutil.WriteFile(fmt.Sprintf("other-holiday.json"), bs, 0644)
	if err !=nil {
		panic(err)
	}
}