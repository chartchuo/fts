package main

import (
	"log"

	"github.com/chartchuo/fts"
)

func main() {
	index := fts.New()
	index.Add("1", "aaa")
	index.Add("2", "aaaa")
	index.Add("3", "bbb")
	index.Add("4", "bbbbb")
	index.Add("5", "aaabbbbb")

	result := index.Search("aaa")
	log.Printf("found %d\n", len(result))
	for _, r := range result {
		log.Println(r)
	}

	result = index.Search("bbb")
	log.Printf("found %d\n", len(result))
	for _, r := range result {
		log.Println(r)
	}
}
