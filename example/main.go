package main

import (
	"log"

	"github.com/chartchuo/fts"
)

func test(index fts.Index, txt string) {
	result := index.Search(txt)
	log.Printf("found %d %s\n", len(result), txt)
	for _, r := range result {
		log.Println(r)
	}

}

func main() {
	index := fts.New()
	index.Add("string#1", "aaa")
	index.Add("string#2", "aaaa")
	index.Add("string#3", "bbb")
	index.Add("string#4", "bbbbb")
	index.Add("string#5", "aaabbbbb")
	index.Add("string#6", "0123456789012345678901234567890123456789")
	index.Add("thai#1", "สวัสดีปีใหม่")
	index.Add("thai#2", "สวัสดี")

	test(index, "aaa")
	test(index, "bbb")
	test(index, "01234567890123456789012345678901234567")
	test(index, "ปีใหม่")
	test(index, "สวัสดี")
	test(index, "ดีปี")

}
