package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key   string
	value string
}

type LRUCache struct {
	Capacity int
	list     *list.List
	hashMap  map[string]*list.Element
}

// Get get key and put elem to the front
func (l *LRUCache) Get(key string) (string, bool) {
	if elem, ok := l.hashMap[key]; ok {
		l.list.MoveToFront(elem)
		fmt.Printf("Get %s, Front (%s,%s), Back (%s,%s) \n", key, l.list.Front().Value.(*entry).key,
			l.list.Front().Value.(*entry).value, l.list.Back().Value.(*entry).key, l.list.Back().Value.(*entry).value)

		return elem.Value.(*entry).value, true
	}
	return "", false
}

func (l *LRUCache) Put(key string, value string) {
	if _, ok := l.Get(key); ok { //
		l.hashMap[key].Value.(*entry).value = value
	} else {
		if len(l.hashMap) == l.Capacity {
			fmt.Printf(" -- delete tail (%s, %s)\n", l.list.Back().Value.(*entry).key, l.list.Back().Value.(*entry).value)
			back := l.list.Back()
			l.list.Remove(back)
			delete(l.hashMap, back.Value.(*entry).key)
		}

		e := l.list.PushFront(&entry{key, value})
		l.hashMap[key] = e
		fmt.Printf("Put (%s,%s), Front (%s,%s), Back (%s,%s), cap: %d len map: %d\n", key, value,
			l.list.Front().Value.(*entry).key, l.list.Front().Value.(*entry).value,
			l.list.Back().Value.(*entry).key, l.list.Back().Value.(*entry).value, l.Capacity, len(l.hashMap))
	}
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.list = list.New()
	lruCache.Capacity = capacity
	lruCache.hashMap = make(map[string]*list.Element, capacity)
	return lruCache
}

func testContainer() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	fmt.Println(e1, e4)
}

func lrutest() {
	l := Constructor(3)
	l.Put("k1", "v1")
	l.Put("k2", "v2")
	l.Put("k3", "v3")
	fmt.Println(l.Get("k1"))
	l.Put("k2", "v22")
	l.Put("k4", "v4")
	fmt.Println(l.Get("k2"))

}

func main() {
	// testContainer()
	lrutest()
}
