package InitializeDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var OpenConnection = make(map[string]*sql.DB)

func init() {
	Tracker, err := sql.Open("mysql", "motoskia_rmlist:RENTMATICS2017@tcp(143.95.248.120:3306)/motoskia_IPMONITOR?charset=utf8")

	//Tracker, err := sql.Open("mysql", "root:mypassword@tcp(172.17.0.6:3306)/GTIRC?charset=utf8")
	//Tracker, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/IPMONITOR?charset=utf8")

	if err != nil {
		fmt.Println("error", err)
	}

	OpenConnection["GTI"] = Tracker

}
func Ret() map[string]*sql.DB {
	return OpenConnection
}
