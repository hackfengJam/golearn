package sqlhelper

import "github.com/go-sql-driver/mysql"

func IsDup(err error) bool {
	if er, ok := err.(*mysql.MySQLError); ok {
		return er.Number == 1062
	}
	return false
}
