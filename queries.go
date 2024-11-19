package main

import (
	"fmt"
	"strings"
	"time"
)

type Query struct {
	Counter   int64
	Interval  int
	Key       string
	Name      string
	Statement string
	UnPivot   bool
	Value     string
}

var Queries = []*Query{
	&Query{
		Name:      "proxysql_global",
		Statement: "SELECT Variable_Name, Variable_Value FROM stats.stats_mysql_global;",
		Key:       "Variable_Name",
		Value:     "Variable_Value",
	},
}

func (q *Query) Beautifier() string {
	q.Statement = strings.ReplaceAll(q.Statement, "\r\n", " ")
	q.Statement = strings.ReplaceAll(q.Statement, "\n", " ")
	q.Statement = strings.ReplaceAll(q.Statement, "\t", " ")
	q.Statement = strings.ReplaceAll(q.Statement, "  ", "")
	q.Statement = strings.Trim(q.Statement, " ")

	return q.Statement
}

func (q *Query) IsTime(i int) bool {
	if q.Interval == 0 {
		return true
	}

	if q.Counter == 0 || int(time.Since(time.Unix(q.Counter, 0)).Seconds()) >= i {
		(*q).Counter = int64(time.Now().Unix())

		return true
	}

	return false
}
