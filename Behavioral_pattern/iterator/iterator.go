package iterator

type Iterator interface {
	HasMore() bool
	Next() interface{}
}

type ConcreteIterator struct {
	collection *ConcreteIteratorCollection
	curIndex   int
}

func (c *ConcreteIterator) HasMore() bool {
	if c.curIndex >= len(c.collection.data) {
		return false
	}
	return true
}

func (c *ConcreteIterator) Next() interface{} {
	res := c.collection.data[c.curIndex]
	c.curIndex += 1
	return res
}

type IterableCollection interface {
	// could be more than one iterator creation function
	CreateIterator() Iterator
}

type ConcreteIteratorCollection struct {
	data []string
}

func (c *ConcreteIteratorCollection) CreateIterator() Iterator {
	return &ConcreteIterator{collection: c}
}

func (c *ConcreteIteratorCollection) AddItem(item string) {
	c.data = append(c.data, item)
}

func NewConcreteIteratorCollection() *ConcreteIteratorCollection {
	return &ConcreteIteratorCollection{}
}
