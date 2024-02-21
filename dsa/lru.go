package main

type LRUCacheNode[T any] struct {
	Value T
	next  *LRUCacheNode[T]
	prev  *LRUCacheNode[T]
}

type LRUCache[K comparable, V any] struct {
	Length   int
	Capacity int
	Head     *LRUCacheNode[V]
	Tail     *LRUCacheNode[V]

	lookup        map[K]*LRUCacheNode[V]
	reverseLookup map[*LRUCacheNode[V]]K
}

func (node *LRUCacheNode[V]) createNode(value V) *LRUCacheNode[V] {
	return &LRUCacheNode[V]{Value: value}
}

func (cache *LRUCache[K, V]) NewLRUCache(capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		Length:        0,
		Capacity:      capacity,
		Head:          nil,
		Tail:          nil,
		lookup:        make(map[K]*LRUCacheNode[V]),
		reverseLookup: make(map[*LRUCacheNode[V]]K),
	}
}

func (cache *LRUCache[K, V]) Update(key K, value V) {
	node := cache.lookup[key]

	if node == nil {
		node = node.createNode(value)
		cache.Length++
		cache.prepend(node)
		cache.trimCache()

		cache.lookup[key] = node
		cache.reverseLookup[node] = key
	} else {
		cache.detach(node)
		cache.prepend(node)
		node.Value = value
	}
}

func (cache *LRUCache[K, V]) Get(key K) *V {
	node := cache.lookup[key]

	if node == nil {
		return nil
	}

	cache.detach(node)
	cache.prepend(node)

	return &node.Value
}

func (cache *LRUCache[K, V]) detach(node *LRUCacheNode[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if cache.Head == node {
		cache.Head = cache.Head.next
	}

	if cache.Tail == node {
		cache.Tail = cache.Tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (cache *LRUCache[K, V]) prepend(node *LRUCacheNode[V]) {
	if cache.Head == nil {
		cache.Head = node
		cache.Tail = node
		return
	}

	node.next = cache.Head
	cache.Head.prev = node
	cache.Head = node
}

func (cache *LRUCache[K, V]) trimCache() {
	if cache.Length <= cache.Capacity {
		return
	}
	tail := cache.Tail
	cache.detach(tail)

	key := cache.reverseLookup[tail]
	delete(cache.lookup, key)
	delete(cache.reverseLookup, tail)
	cache.Length--
}