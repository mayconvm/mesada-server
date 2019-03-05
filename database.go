package main


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


func connDatabase() {
	host :="172.18.0.2"
	pass :="mesada"
	user :="mesada"
	database :="MESADA"

	orm.RegisterDataBase("default", "mysql", user + ":" + pass + "@tcp(" + host + ":3306)/" + database + "?charset=utf8", 30)

	// env dev
	orm.Debug = true

	orm.RunSyncdb("default", false, true)
}