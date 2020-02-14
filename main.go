package main

import (
	"fmt"
	"log"
	"regexp"
	"st-crawler/fetcher"
)

func main() {

	bytes, err := fetcher.Fetcher("https://xuangubao.cn/live")
	if err != nil {
		log.Printf("fetcher error :%v",err)
	}
	fmt.Printf("%s",bytes)
}
func printHeaderList(content []byte)  {

	compile := regexp.MustCompile(`<header.+class="title_[0-9a-zA-Z]+"><a.+>([^<]+)</a>.+</header>`)
	header := compile.FindAllSubmatch(content,-1)
	for _,s := range header{
		for k,title := range s{

			if k != 1 {
				continue;
			}
			fmt.Printf("%s\n", title)
		}
	}
	fmt.Printf("%d/n",len(header))
}
