package main

import (
	"time"

	"github.com/debeando/go-common/env"
	"github.com/debeando/go-common/log"
	"github.com/debeando/go-common/mysql"
)

func main() {
	log.Info("Start DeBeAndo Agent for ProxySQL")

	if getDebug() {
		log.SetLevel(log.DebugLevel)
	}

	log.DebugWithFields("Environment Variables", log.Fields{
		"DEBUG":             getDebug(),
		"INFLUXDB_BUCKET":   influxDB.Bucket,
		"INFLUXDB_HOST":     influxDB.Host,
		"INFLUXDB_PORT":     influxDB.Port,
		"INFLUXDB_TOKEN":    influxDB.Token,
		"INTERVAL":          getInterval(),
		"PROXYSQL_HOST":     MySQL.Host,
		"PROXYSQL_PASSWORD": MySQL.Password,
		"PROXYSQL_PORT":     MySQL.Port,
		"PROXYSQL_TIMEOUT":  MySQL.Timeout,
		"PROXYSQL_USER":     MySQL.Username,
		"SERVER":            getServer(),
	})

	influxDB.New()
	defer influxDB.Close()
	MySQL.Connection = mysql.New(MySQL.Host, MySQL.DSN())
	defer MySQL.Connection.Close()

	for {
		metrics := Metrics{}

		MySQL.Connection.Connect()

		for _, query := range Queries {
			metric := Metric{}

			if !query.IsTime(query.Interval) {
				continue
			}

			MySQL.Connection.FetchAll(query.Beautifier(), func(row map[string]string) {
				metric.Measurement = query.Name
				metric.AddTag(Tag{
					Name:  "server",
					Value: getServer(),
				})

				if query.UnPivot {
					for column, value := range row {
						if valueParsed, ok := mysql.ParseNumberValue(value); ok {
							metric.AddField(Field{
								Name:  column,
								Value: valueParsed,
							})
						} else {
							metric.AddTag(Tag{
								Name:  column,
								Value: value,
							})
						}
					}
				} else if valueParsed, ok := mysql.ParseNumberValue(row[query.Value]); ok {
					metric.AddField(Field{
						Name:  row[query.Key],
						Value: valueParsed,
					})
				}
			})

			metrics.Add(metric)
		}

		if metrics.Count() > 0 {
			influxDB.Write(metrics)
		}

		metrics.Reset()
		log.Debug("Wait until next collect metrics.")
		time.Sleep(getInterval())
	}
}

func getDebug() bool {
	return env.GetBool("DEBUG", true)
}

func getInterval() time.Duration {
	return time.Duration(env.GetInt("INTERVAL", 3)) * time.Second
}

func getServer() string {
	return env.Get("SERVER", MySQL.Host)
}
