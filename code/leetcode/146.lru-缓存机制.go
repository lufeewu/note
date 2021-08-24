/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	Capacity int
	list     *list.List
	hashMap  map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.list = list.New()
	lruCache.Capacity = capacity
	lruCache.hashMap = make(map[int]*list.Element, capacity)
	return lruCache
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.hashMap[key]; ok {
		l.list.MoveToFront(elem)

		return elem.Value.(*entry).value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if l.Get(key) != -1 { //
		l.hashMap[key].Value.(*entry).value = value
	} else {
		if len(l.hashMap) == l.Capacity {
			back := l.list.Back()
			l.list.Remove(back)
			delete(l.hashMap, back.Value.(*entry).key)
		}

		e := l.list.PushFront(&entry{key, value})
		l.hashMap[key] = e
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

