package fts

import (
	"regexp"
	"strings"
)

const minLen = 3

type Index struct {
	maxDbId int
	db      map[int]string
	index   map[string]map[int]bool
}

func word2token(word string) map[string]bool {
	word = strings.ToLower(word)
	tokens := make(map[string]bool)
	for i := minLen; i <= len(word); i++ {
		for j := 0; j <= len(word)-i; j++ {
			tokens[word[j:j+i]] = true
		}
	}

	return tokens
}

func New() Index {
	var idx Index
	idx.maxDbId = 0
	idx.db = make(map[int]string)
	idx.index = make(map[string]map[int]bool)
	return idx
}

func (idx Index) Add(id string, searchText string) {

	idx.maxDbId++
	idx.db[idx.maxDbId] = id

	searchText = strings.TrimSpace(searchText)
	space := regexp.MustCompile(`\s+`)
	searchText = space.ReplaceAllString(searchText, " ")
	words := strings.Split(searchText, " ")
	for _, w := range words {
		tokens := word2token(w)
		for t := range tokens {
			if idx.index[t] == nil {
				idx.index[t] = make(map[int]bool)
			}
			idx.index[t][idx.maxDbId] = true
		}
	}
}

func (idx Index) Search(searchText string) []string {
	searchText = strings.TrimSpace(searchText)
	space := regexp.MustCompile(`\s+`)
	searchText = space.ReplaceAllString(searchText, " ")

	result := make(map[int]int)
	words := strings.Split(searchText, " ")

	wordCount := 0
	for _, w := range words {
		if len(w) < minLen {
			continue
		}
		wordCount++
		w = strings.ToLower(w)

		for id := range idx.index[w] {
			if _, exist := result[id]; exist {
				result[id] = result[id] + 1
			} else {
				result[id] = 1
			}
		}

	}

	var ids []string

	for id, count := range result {
		if count == wordCount {
			ids = append(ids, idx.db[id])
		}
	}
	return ids
}
