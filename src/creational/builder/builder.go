package builder

import "bytes"

/*
  定义一个对象
*/

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Person{")
	buf.WriteString("name='")
	buf.WriteString(p.Name)
	buf.WriteString(", age=")
	buf.WriteString(string(p.Age))
	buf.WriteString("}")
	return buf.String()
}

/*
  定义Builder方法
*/
type Builder interface {
	SetName(name string) Builder
	SetAge(age int) Builder
	Build() *Person
}

type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) SetName(name string) Builder {
	if p.person == nil {
		p.person = &Person{}
	}
	p.person.SetName(name)
	return p
}

func (p *PersonBuilder) SetAge(age int) Builder {
	if p.person == nil {
		p.person = &Person{}
	}
	p.person.SetAge(age)
	return p
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}
