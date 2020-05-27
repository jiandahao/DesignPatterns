package prototype

/*
解决的问题：
如果你有一个对象， 并希望生成与其完全相同的一个复制品，
你该如何实现呢？ 首先， 你必须新建一个属于相同类的对象。
然后， 你必须遍历原始对象的所有成员变量， 并将成员变量值复制到新对象中。

但有个小问题。 并非所有对象都能通过这种方式进行复制， 因为有些对象可能
拥有私有成员变量， 它们在对象本身以外是不可见的。

直接复制还有另外一个问题。 因为你必须知道对象所属的类才能创建复制品， 所以代码必须依赖该类。
*/

/*
在原型模式中， 你可以使用一系列预生成的、 各种类型的对象作为原型。
客户端不必根据需求对子类进行实例化， 只需找到合适的原型并对其进行克隆即可。
*/

type ProtoType interface {
	/*也可以添加除Clone之外的其他方法*/
	Clone() ProtoType
}

type ConcreteProtoTypeA struct {
	Name 	string
	Age 	int
}

func (p *ConcreteProtoTypeA) Clone() ProtoType {
	return &ConcreteProtoTypeA{
		Name: p.Name,
		Age:  p.Age,
	}
}

type ConcreteProtoTypeB struct {
	Nickname 	string
	PhoneNumber string
}

func (p *ConcreteProtoTypeB) Clone() ProtoType {
	return &ConcreteProtoTypeB{
		Nickname:    p.Nickname,
		PhoneNumber: p.PhoneNumber,
	}
}

// 创建一个中心化原型注册表， 用于存储常用原型
type Registry struct {
	items map[string]ProtoType
}

func (r *Registry) Register(id string, protoType ProtoType) {
	if r.items ==  nil {
		r.items = map[string]ProtoType{}
	}

	r.items[id] = protoType
}

func (r *Registry) CreateProtoTypeById(id string) ProtoType {
	if r.items == nil {
		return nil
	}

	if item, ok := r.items[id]; ok {
		return item.Clone()
	}

	return nil
}