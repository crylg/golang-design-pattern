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
 定义工厂接口
*/
type Factory interface{
	CreateOperate() Operate
}
/*
  实现加法工厂
*/
type AddFactory struct {
}

func NewAddFactory() *AddFactory {
	return new(AddFactory)
}

func (this *AddFactory) CreateOperate() Operate {
	return &Add{}
}
/*
  实现减法工厂
*/
type SubFactory struct {
}

func NewSubFactory() *SubFactory {
	return new(SubFactory)
}

func (this *SubFactory) CreateOperate() Operate {
	return &Sub{}
}
