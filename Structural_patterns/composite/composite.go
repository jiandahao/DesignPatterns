package composite

// 组合模式，对象树
// a. 如果你需要实现树状对象结构， 可以使用组合模式。
//    组合模式为你提供了两种共享公共接口的基本元素类型：
//    简单叶节点和复杂容器。 容器中可以包含叶节点和其他容器。 这使得你可以构建树状嵌套递归对象结构。
// b. 如果你希望客户端代码以相同方式处理简单和复杂元素， 可以使用该模式。
//    组合模式中定义的所有元素共用同一个接口。 在这一接口的帮助下， 客户端不必在意其所使用的对象的具体类。
import (
	"fmt"
	"reflect"
)

// 组件接口会声明组合中简单和复杂对象的通用操作。
type BaseComponent interface {
	Draw()
	Move(x float32, y float32)
}

// 叶节点类代表组合的终端对象。叶节点对象中不能包含任何子对象。叶节点对象
// 通常会完成实际的工作，组合对象则仅会将工作委派给自己的子部件。
// 这边以"点"类为例, 该叶节点（简单对象）支持Draw和Move操作
type Dot struct {
	X float32
	Y float32
}

func (d *Dot) Draw() {
	fmt.Printf("Draw a dot at (%v,%v)\n", d.X, d.Y)
}

func (d *Dot) Move(x float32, y float32) {
	fmt.Printf("Move dot from (%v,%v) to (%v,%v)\n", d.X, d.Y, x, y)
	d.X = x
	d.Y = y
}

func NewDot(x float32, y float32) *Dot {
	return &Dot{
		X: x,
		Y: y,
	}
}

/**
 * The Composite class represents the complex components that may have children.
 * Usually, the Composite objects delegate the actual work to their children and
 * then "sum-up" the result.
 */
type Composite struct {
	children []BaseComponent
}

func NewComposite() *Composite {
	return &Composite{}
}

func (c *Composite) AddComponent(component BaseComponent) {
	if reflect.DeepEqual(component, c) {
		fmt.Println("error")
		return
	}
	if c == nil {
		return
	}
	c.children = append(c.children, component)
}

func (c *Composite) RemoveComponent(component BaseComponent) {
	if c == nil {
		return
	}
	for index, child := range c.children {
		if reflect.DeepEqual(child, component) {
			c.children = append(c.children[:index], c.children[index+1:]...)
			return
		}
	}
}

func (c *Composite) Draw() {
	if c == nil {
		return
	}
	for _, child := range c.children {
		child.Draw()
	}
}

func (c *Composite) Move(x float32, y float32) {
	if c == nil {
		return
	}
	for _, child := range c.children {
		child.Move(x, y)
	}
}
