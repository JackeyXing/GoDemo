//设计模式_创建型模式_抽象工厂模式
package main

import "fmt"

//Factory 一个标准工厂
type Factory interface {
	Workshop() Workshop //工作车间
	Eatery() Eatery     //食堂
}

//Workshop 车间标准
type Workshop interface {
	Work()
}

//Eatery 食堂标准
type Eatery interface {
	Eating()
}

//FactoryA 工厂A
type FactoryAInfo struct{}

func (o *FactoryAInfo) Workshop() Workshop { return &FactoryA{} } //返回实例化的工厂A
func (o *FactoryAInfo) Eatery() Eatery     { return &FactoryA{} } //返回实例化的工厂A

//FactoryA 工厂A的车间和食堂 (可以把车间和食堂分开)
type FactoryA struct{}

func (o *FactoryA) Work()   { fmt.Println("FactoryA::Work") }
func (o *FactoryA) Eating() { fmt.Println("FactoryA::Eating") }

//FactoryB 工厂B
type FactoryBInfo struct{}

func (o *FactoryBInfo) Workshop() Workshop { return &FactoryB{} } //返回实例化的工厂A
func (o *FactoryBInfo) Eatery() Eatery     { return &FactoryB{} } //返回实例化的工厂A

//FactoryB 工厂B的车间和食堂 (可以把车间和食堂分开)
type FactoryB struct{}

func (o *FactoryB) Work()   { fmt.Println("FactoryB::Work") }
func (o *FactoryB) Eating() { fmt.Println("FactoryB::Eating") }

//通过提供的参数创建不同的工厂
func FindWork(Money int) Factory {
	if Money > 10000 {
		return &FactoryAInfo{}
	}else {
		return &FactoryBInfo{}
	}
}

func main() {
	w20k := FindWork(20000)
	w20k.Workshop().Work()
	w20k.Eatery().Eating()
	
	w2k := FindWork(2000)
	w2k.Workshop().Work()
	w2k.Eatery().Eating()
}