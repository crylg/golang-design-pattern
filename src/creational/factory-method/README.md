工厂方法模式
=====

# 简介
**工厂方法模式**是Golang中最常用的设计模式之一，属于创建型模式，它提供了一种创建对象的最佳方式。

# 意图
定义一个用于创建对象的工厂，根据不同的接收条件具体实例化相应的工作类。

# 工厂方法模式有三个要素
+ **产品接口(Arithmetic)** 定义产品的规范，所有的产品实现都必须遵循产品接口定义的规范。产品接口是调用者最为关心的，产品接口定义的优劣直接决定了调用者代码的稳定性
```
   type Arithmetic interface {
   	Calculate(int, int) int
   }
```
+ **产品实现(Add/Sub/Mul/Div)** 产品接口的具体实现类，不同的产品也需要不同的产品实现类，产品实现类与功能创建类相对应
```
  type Add struct{}

  func (a *Add) Calculate(left int, right int) int {
  	return left + right
  }
```
+ **工厂实现(Factory)** 决定如何实例化产品，是实现扩展的途径，与调用者直接交互用来提供产品的创建。
```
   func Factory(operation string) Arithmetic {
   	switch  operation {
   	case add:
   		return &Add{}
   	case sub:
   		return &Sub{}
   	case mul:
   		return new(Mul)
   	case div:
   		return new(Div)
   	default:
   		panic("invalid operation")
   	}
   }
```
# 优点
1. 一个调用者想创建一个对象，只要知道其名称就可以了。
2. 扩展性高，如果想增加一个产品，只要扩展工厂类就可以。
3. 屏蔽产品的具体实现，调用者只关心产品的接口。

# 缺点
1. 每次增加一个产品时，都需要增加一个具体类和对象实现工厂。
2. 使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。
