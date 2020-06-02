package main

import (
	"fmt"
	"github.com/jiandahao/DesignPatterns/Behavioral_pattern/iterator"
)

func main() {

	collect := iterator.NewConcreteIteratorCollection()
	collect.AddItem("name")
	collect.AddItem("gender")
	collect.AddItem("age")
	TestIterator(collect)
}

// 该函数接受一个实现了迭代器的容器/集合，通过CreateIterator接口获取迭代器，对容器进行
// 遍历。当更改容器或者容器的遍历策略更改，该函数并不需要做任何更改
func TestIterator(collect iterator.IterableCollection) {
	it := collect.CreateIterator()

	for it.HasMore() {
		fmt.Println(it.Next().(string))
	}
}
