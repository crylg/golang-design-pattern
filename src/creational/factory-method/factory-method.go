package factory_method

/*
 首先定义一个运算接口
*/
type Operator interface {
	Operator(int, int) int
}

/*
 定义 加法运算
*/
type Add struct {
}

func (this *Add) Operator(rhs int, lhs int) int {
	return rhs + lhs
}

/*
 定义减法运算
*/
type Sub struct {
}

func (this *Sub) Operator(rhs int, lhs int) int {
	return rhs - lhs
}

/*
 定义工厂
*/
type OperatorFactory struct {
}

func NewOperatorFactory() *OperatorFactory {
	return new(OperatorFactory)
}

func (this *OperatorFactory) CreateOperate(operator string) Operator {
	switch operator {
	case "+":
		return &Add{}
	case "-":
		return &Sub{}
	default:
		panic("invalid operator")
		return nil
	}
}
