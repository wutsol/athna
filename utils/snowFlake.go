package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strconv"
)

var node *snowflake.Node

func InitSnowFlake() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func UniqueId() int64 {
	id := node.Generate()
	fmt.Printf("unique id:%v",id)
	rs := []rune(id.String())
	realId, _ := strconv.ParseInt(string(rs[len(rs)-15:]), 10, 64) //取15位
	return realId
}
