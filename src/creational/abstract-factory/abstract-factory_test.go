package abstract_factory

import "testing"

const (
	mysql_driver  = "com.mysql.Driver"
	sqlite_driver = "com.sqlite.Driver"
	oracle_driver = "com.oracle.Driver"
)

func TestAbstractFactory(t *testing.T) {
	mysqlFactory := new(MysqlFactory)
	mysql := mysqlFactory.Create()
	if mysql != nil {
		mysql.Registry(mysql_driver)
	} else {
		t.Error("create mysql driver error")
	}

	oracleFactory := new(OracleFactory)
	oracle := oracleFactory.Create()
	if oracle != nil {
		oracle.Registry(oracle_driver)
	} else {
		t.Error("create oracle driver error")
	}

	sqliteFactory := new(SqliteFactory)
	sqlite := sqliteFactory.Create()
	if sqlite != nil {
		sqlite.Registry(sqlite_driver)
	} else {
		t.Error("create sqlite driver error")
	}
}
