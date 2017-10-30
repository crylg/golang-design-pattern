package prototype

import (
	"bytes"
)

type Prototype struct {
	name string
	sex  string
	age  int
}

func (this *Prototype) SetName(name string) {
	this.name = name
}

func (this *Prototype) SetSex(sex string) {
	this.sex = sex
}

func (this *Prototype) SetAge(age int) {
	this.age = age
}

func (this *Prototype) Clone() *Prototype {
	clone := *this
	return &clone
}

func (this *Prototype) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Prototype:{ name = ")
	buf.WriteString(this.name)
	buf.WriteString(", sex = ")
	buf.WriteString(this.sex)
	buf.WriteString(", age = ")
	buf.WriteString(string(this.age))
	buf.WriteString("}")
	return buf.String()

}
