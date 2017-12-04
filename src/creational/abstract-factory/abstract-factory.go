package abstract_factory

import (
	"log"
)

const (
	mysql  = "mysql"
	oracle = "oracle"
	sqlite = "sqlite"
)

/*
  define interface: Driver
  define abstract create factory interface: AbstractFactory
*/
type Driver interface {
	Registry(name string)
}

type AbstractFactory interface {
	Create() Driver
}

/*
  Mysql Class implement Driver
  MysqlFactory implement AbstractFactory
*/
type Mysql struct{}

func (m *Mysql) Registry(name string) {
	log.Printf("registry %s driver with: %s", mysql, name)
}

type MysqlFactory struct{}

func (m *MysqlFactory) Create() Driver {
	return new(Mysql)
}

/*
  Oracle implements Driver
  OracleFactory implement AbstractFactory
*/
type Oracle struct{}

func (o *Oracle) Registry(name string) {
	log.Printf("registry %s driver with: %s", oracle, name)
}

type OracleFactory struct{}

func (c *OracleFactory) Create() Driver {
	return new(Oracle)
}

/*
 Sqlite implements Driver
 SqliteFactory implements AbstractFactory
*/
type Sqlite struct{}

func (s *Sqlite) Registry(name string) {
	log.Printf("registry %s driver with: %s", sqlite, name)
}

type SqliteFactory struct{}

func (s *SqliteFactory) Create() Driver {
	return new(Sqlite)
}
