package main

import "fmt"

type LRUCache interface {
	set(key, value int) bool
	get(key int)  int
}

type val struct{
	accessCount int
	value int
}

type lrucache struct {
	size int
	cache map[int]val
}

func newLRUCache(size int) *lrucache{
	return &lrucache{size:size, cache:make(map[int]val)}
}

func (c *lrucache)get(key int)  int{
		val, ok := c.cache[key]

		if ok{
			val.accessCount++
			c.cache[key] = val
			return val.value
		}
		return -1
}

func deleteLeastUsed(c *lrucache){
	var min, ckey int
	for k,v := range c.cache{
		if v.accessCount <= min {
			min = v.accessCount
			ckey = k
		}
	}

	delete(c.cache, ckey)
}

func (c *lrucache)set(key , value int){
	if len(c.cache) >= c.size {
			deleteLeastUsed(c)
	}
	c.cache[key] = val{value:value}
}

func main() {
	lru := newLRUCache(4)
	lru.set(1, 1)
	lru.set(2, 2)
	lru.set(3, 3)
	lru.set(4, 4)
	lru.get(1)
	lru.get(3)
	lru.get(4)

	fmt.Println(lru.cache)

	lru.set(5,5)

	fmt.Println(lru.cache)

}
