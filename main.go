package main

import (
	"st-crawler/XuanGuBao/parser"
	"st-crawler/engine"
)

func main() {

	engine.Run(engine.Request{Url:"https://xuangubao.cn/live",ParseFunction:XuanGuBao.LiveParse})
	// bytes, err := fetcher.Fetcher("https://xuangubao.cn/live")
	// if err != nil {
	// 	log.Printf("fetcher error :%v",err)
	// }
	// fmt.Printf("%s",bytes)
}

