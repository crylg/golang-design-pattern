package factory_method

/*
 首先定义一个运算接口
*/
type Operate interface {
	Calculate(int, int) int
}

/*
 定义 加法运算
*/
type Add struct {
}

func (this *Add) Calculate(rhs int, lhs int) int {
	return rhs + lhs
}

/*
 定义减法运算
*/
type Sub struct {
}

func (this *Sub) Calculate(rhs int, lhs int) int {
	return rhs - lhs
}

/*
 定义工厂
*/
type Factory interface{
	CreateOperate() Operate
}

type AddFactory struct {
}

func NewAddFactory() *AddFactory {
	return new(AddFactory)
}

func (this *AddFactory) CreateOperate() Operate {
	return &Add{}
}

type SubFactory struct {
}

func NewSubFactory() *SubFactory {
	return new(SubFactory)
}

func (this *SubFactory) CreateOperate() Operate {
	return &Sub{}
}

