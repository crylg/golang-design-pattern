package prototype

import (
	"bytes"
	"strconv"
)

type Prototype struct {
	name string
	sex  string
	age  int
}

func (p *Prototype) SetName(name string) {
	p.name = name
}

func (p *Prototype) SetSex(sex string) {
	p.sex = sex
}

func (p *Prototype) SetAge(age int) {
	p.age = age
}

func (p *Prototype) Clone() *Prototype {
	clone := *p
	return &clone
}

func (p *Prototype) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Prototype:{ name = ")
	buf.WriteString(p.name)
	buf.WriteString(", sex = ")
	buf.WriteString(p.sex)
	buf.WriteString(", age = ")
	buf.WriteString(strconv.Itoa(p.age))
	buf.WriteString("}")
	return buf.String()
}
