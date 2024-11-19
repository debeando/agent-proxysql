package main

import (
	"github.com/debeando/go-common/env"
	"github.com/debeando/go-common/mysql"
)

var MySQL mysql.MySQL

func init() {
	MySQL.Host = env.Get("PROXYSQL_HOST", "127.0.0.1")
	MySQL.Password = env.Get("PROXYSQL_PASSWORD", "monitoring")
	MySQL.Port = env.GetUInt16("PROXYSQL_PORT", 3306)
	MySQL.Timeout = env.GetUInt8("PROXYSQL_TIMEOUT", 10)
	MySQL.Username = env.Get("PROXYSQL_USER", "monitoring")
}
