package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

var engin *gorose.Engin

func init() {
	// log.Println("init db")
	connStr := Mc.User + ":" +Mc.Pwd + "@tcp(" +Mc.Host + ")/" +Mc.Db + "?charset=utf8"
	// log.Println(connStr)
	engin, _ = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: connStr})
}
func DB() gorose.IOrm {
	return engin.NewOrm()
}
