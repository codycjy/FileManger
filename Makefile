main: cmd/main.go
	go run cmd/main.go

testMysql: test/MysqlTest/MysqlTestConnection.go
	go run test/MysqlTest/MysqlTestConnection.go
