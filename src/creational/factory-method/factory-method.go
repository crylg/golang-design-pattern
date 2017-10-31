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
 定义工厂方法
*/
type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (this *Factory) CreateOperate(operateName string) Operate {
	switch operateName {
	case "+":
		return &Add{}
	case "-":
		return &Sub{}
	default:
		panic("invalid operate")
	}
}
