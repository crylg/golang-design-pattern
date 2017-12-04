package builder

import (
	"bytes"
	"strconv"
)

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

func newPerson() *Person {
	return new(Person)
}

func (p *Person) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Person{")
	buf.WriteString("name='")
	buf.WriteString(p.Name)
	buf.WriteString(", age=")
	buf.WriteString(strconv.Itoa(p.Age))
	buf.WriteString("}")
	return buf.String()
}

/*
  定义Builder方法
*/
type Builder interface {
	Name(name string) Builder
	Age(age int) Builder
	Build() *Person
}

type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) Name(name string) Builder {
	if p.person == nil {
		p.person = newPerson()
	}
	p.person.SetName(name)
	return p
}

func (p *PersonBuilder) Age(age int) Builder {
	if p.person == nil {
		p.person = newPerson()
	}
	p.person.SetAge(age)
	return p
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}
