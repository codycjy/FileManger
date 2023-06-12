main: cmd/main.go
	go run cmd/main.go

testMysql: test/MysqlTest/MysqlTestConnection.go
	go run test/MysqlTest/MysqlTestConnection.go

test: test/ModelTest/FileModelTest.go cmd/main.go
	go run test/ModelTest/FileModelTest.go
	go run cmd/main.go
