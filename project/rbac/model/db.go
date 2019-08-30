package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/url"
)

var db *sqlx.DB

func init() {
	PrepareDbConnection()
	log.Println("准备数据库连接完毕")
}

func AppendParseTime(dsn string) string {
	u, err := url.Parse(dsn)
	if err != nil {
		return dsn
	}
	query := u.Query()
	query.Set("parseTime", "true")
	query.Set("loc", "Local")
	u.RawQuery = query.Encode()
	return u.String()
}

func PrepareDbConnection() {
	var err error
	mysqlDsn := "root:123345@tcp(127.0.0.1:3307)/abac_test?charset=utf8mb4&collation=utf8mb4_unicode_ci"
	if mysqlDsn == "" {
		log.Panicln("config not found: dsn")
	}
	db, err = sqlx.Open("mysql", AppendParseTime(mysqlDsn))
	if err != nil {
		log.Panic(err)
	}
	//db.LogMode(true)
}
