package main

import(
	"fmt"
	"web_spider/fetcher"
	"web_spider/parser"
)

var used_url = map[string]int{}

func main(){
	//种子url
	url := "https://www.hao123.com/"
	body, _ := fetcher.Fetch(url)
	n := 0
	for _, url_new := range parser.Parser(body){
		fmt.Println(url_new)
		if usedurl(url_new, used_url){
			n++
		}else{
			continue
		}		
	}
	fmt.Println(n)
}

// 判断url是否已经请求过
func usedurl(a string, m map[string]int)bool{
	_, ok := m[a]
	if !ok{
		m[a] = 0
		return !ok
	}else{
		fmt.Println(a)
		return !ok
	}
}