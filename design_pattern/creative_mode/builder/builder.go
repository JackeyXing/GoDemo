// Builder 设计模式_创建型模式_生成器
package main 

import "fmt"

//Builder 生成器接口定义
type Builder interface {
	Produce()   //生产
	Check()     //检测
	DeepCheck() //深度检测
}

//FactoryA 工厂A
type FactoryA struct {}

func (f *FactoryA) Produce() { fmt.Println("FactoryA::Produce") }
func (f *FactoryA) Check() { fmt.Println("FactoryA::Check") }
func (f *FactoryA) DeepCheck() { fmt.Println("nothing") }

//FactoryB 工厂B
type FactoryB struct {}

func (f *FactoryB) Produce() { fmt.Println("FactoryB::Produce") }
func (f *FactoryB) Check() { fmt.Println("FactoryB::Check") }
func (f *FactoryB) DeepCheck() { fmt.Println("FactoryB::DeepCheck") }


//导向器 （主管/导演） ---控制builder
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

//Build 普通生产
func (c *Director) Build() {
	c.builder.Produce()
	c.builder.Check()
}

//BuildForVip 高端生产
func (c *Director) BuildForVip() {
	c.builder.Produce()
	c.builder.Check()
	c.builder.DeepCheck()
}

//FindBuilder 创建生成器
func FindBuilder(money int) Builder {
	if money > 10000 {
		return &FactoryB{}
	} else {
		return &FactoryA{}
	}
}

func main() {
	normal := FindBuilder(2000)
	normalDirectory := NewDirector(normal)
	normalDirectory.Build()
	normalDirectory.BuildForVip()

	vip := FindBuilder(20000)
	vipDirectory := NewDirector(vip)
	vipDirectory.Build()
	vipDirectory.BuildForVip()

}