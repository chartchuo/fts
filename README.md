# FTS
Very simple and easy full text search library. 
- In-memory key value mapping.
- Suitable for small number of data (<10000 entries)
- Support non-english example Thai
  
## Example
```golang
	index := fts.New()
	index.Add("string#1", "aaa")
	index.Add("string#2", "aaaa")
	index.Add("string#3", "bbb")
	index.Add("string#4", "bbbbb")
	index.Add("string#5", "aaabbbbb")
	index.Add("string#6", "0123456789012345678901234567890123456789")
	index.Add("thai#1", "สวัสดีปีใหม่")
	index.Add("thai#2", "สวัสดี")
    
	result := index.Search("aaa")
	log.Printf("found %d %s\n", len(result), txt)
	for _, r := range result {
		log.Println(r)
	}
```
  