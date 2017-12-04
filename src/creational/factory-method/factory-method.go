package factory_method

const (
	add = "+"
	sub = "-"
	mul = "*"
	div = "/"
)

type Arithmetic interface {
	Calculate(int, int) int
}

type Add struct{}

func (a *Add) Calculate(left int, right int) int {
	return left + right
}

type Sub struct{}

func (s *Sub) Calculate(left int, right int) int {
	return left - right
}

type Mul struct{}

func (m *Mul) Calculate(left int, right int) int {
	return left * right
}

type Div struct{}

func (d *Div) Calculate(left int, right int) int {
	return left / right
}

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
