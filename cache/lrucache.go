package main

import (
	"container/list"
	"sync"
)

type Cacher interface {
	Get(key string)
	Set(key string, val interface{})
	Delete()
}


type Cache struct {
	mu sync.Mutex

	list  *list.List
	table map[string]*list.Element

	maxLen int
}

func NewCache(maxLen int) Cacher {
	return &Cache{
		list: list.New(),
		table: make(map[string]*list.Element, maxLen),
		maxLen: maxLen,
	}
}