package main

import (
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
)

const mytableSQL2 = "CREATE TABLE IF NOT EXISTS `mytables` (" +
	"    `ID` int(11) NOT NULL," +
	"    `username` varchar(16) DEFAULT NULL," +
	"    `age` int(8) NOT NULL," +
	"    `phone` varchar(11) NOT NULL," +
	"    INDEX `idx_username` (`username`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"

const usertableSQL = "create table IF NOT EXISTS user (" +
	" id bigint unsigned auto_increment primary key," +
	"  username varchar(255) null," +
	"  password varchar(255) null," +
	"  phone varchar(255) null," +
	"  create_time timestamp default CURRENT_TIMESTAMP not null," +
	"  update_time timestamp default CURRENT_TIMESTAMP not null," +
	"  is_deleted int null);"

const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/week2?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnectDB(MySQLDSN).Debug()

	//prepare(dal.DB) // prepare table for generate ./generate.go:12:2: undefined: prepare1
	dal.DB.Exec(usertableSQL)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		//OutPath: "../../dal/query",
		OutPath: "/Users/zhanlifeng/Documents/workspace/go/week2/query",
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
