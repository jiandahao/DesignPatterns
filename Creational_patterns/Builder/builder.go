package builder

import "fmt"

// 生成器接口声明了创建产品对象不同部件的方法。
type Builder interface {
	// reset（重置）方法可清除正在生成的对象。
	BuildPartA()
	BuildPartB()
}
type ConcreteBuilder struct {
	product *Product
}
// 实现Builder的接口以构造和装配该产品的各个部件。
// 每个ConcreteBuilder包含了创建和装配一个特定产品的所有代码。
// 这些代码只需要写一次，然后不同的Director可以复用它以在相同部件集合的基础上
// 构造出不同的product
func NewConcreteBuilder() *ConcreteBuilder{
	return &ConcreteBuilder{product:&Product{}}
}
func (b *ConcreteBuilder) BuildPartA() {
	fmt.Println("Building part A")
	b.product.addItem("PartA")
}

func (b *ConcreteBuilder) BuildPartB() {
	fmt.Println("Building part B")
	b.product.addItem("PartB")
}

func (b *ConcreteBuilder) Reset() {
	b.product = nil
}

func (b *ConcreteBuilder) GetResult() []string{
	result := b.product.Content
	b.Reset()
	return result
}

// 构造一个使用builder接口的对象
type Director struct {
	builder Builder
}
func NewDirector(builder Builder) *Director{
	return &Director{builder:builder}
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
}

// 表示被构造的复杂对象，ConcreteBuilder创建该产品的内部表示并定义它的装配过程
type Product struct {
	Content 	[]string
}
func (p *Product) addItem(item string) {
	p.Content = append(p.Content, item)
}


