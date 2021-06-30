package cache

type CacheNode struct {
	Value    string
	Key      string
	Next     *CacheNode
	Previous *CacheNode
}

// Note:
// Enqueueing is done at the head
// Dequeueing is done from the tail
type LRUCache struct {
	Head      *CacheNode
	Tail      *CacheNode
	Length    int
	MaxLength int
	CacheMap  map[string]*CacheNode
}

func CreateCache(maxLength int) *LRUCache {
	cache := LRUCache{}
	cache.CacheMap = make(map[string]*CacheNode)
	cache.MaxLength = maxLength
	cache.Head = &CacheNode{}
	cache.Tail = &CacheNode{}
	cache.Head.Next = cache.Tail
	cache.Tail.Previous = cache.Head

	return &cache
}

func (l *LRUCache) Enqueue(node *CacheNode) {
	l.Length += 1
	l.Head.Next.Previous = node
	node.Next = l.Head.Next
	node.Previous = l.Head
	l.Head.Next = node
}

func (l *LRUCache) Dequeue() *CacheNode {
	l.Length -= 1
	dequeuedNode := l.Tail.Previous
	l.Tail.Previous = dequeuedNode.Previous
	l.Tail.Previous.Next = l.Tail

	return dequeuedNode
}

func (l *LRUCache) Walk(callback func(*CacheNode)) {
	currentNode := l.Head.Next

	for currentNode.Next != nil {
		callback(currentNode)
		currentNode = currentNode.Next
	}
}

func (l *LRUCache) AddNewNode(operationKey string, answerValue string) {
	node := &CacheNode{Value: answerValue, Key: operationKey}

	l.CacheMap[operationKey] = node
	l.Enqueue(node)

	if l.Length > l.MaxLength {
		removedNode := l.Dequeue()
		delete(l.CacheMap, removedNode.Key)
	}
}

type Operation struct {
	Num1     int
	Operator string
	Num2     int
}
