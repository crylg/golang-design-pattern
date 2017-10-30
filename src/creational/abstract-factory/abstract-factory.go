package abstract_factory

import (
	"log"
)

/*
  首先定义了interface Worker，其中包含一个方法Work(task *string)
  然后定义创建Worker的interface WorkerCreator，包含一个接口Create() ，返回值为Worker
*/

type Worker interface {
	Work(task string)
}

type WorkerCreator interface {
	Create() Worker
}

/*
 假设有一个工种程序员，实现了Work接口，同时定义其构造类，实现Create()接口
*/
type Programmer struct {
}

func (p *Programmer) Work(task string) {
	log.Println("Programmer process", task)
}

type ProgrammerCreator struct {
}

func (c *ProgrammerCreator) Create() Worker {
	return new(Programmer)
}

/*
  有一个工种农场主，实现了Work接口，同时定义其构造类，实现Create()接口
*/
type Farmer struct {
}

func (f *Farmer) Work(task string) {
	log.Println("Farmer process", task)
}

type FarmerCreator struct {
}

func (c *FarmerCreator) Create() Worker {
	return new(Farmer)
}
