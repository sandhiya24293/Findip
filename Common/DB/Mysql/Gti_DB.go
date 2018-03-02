package Mysql

import (
	_ "database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Insertip(ip string, mailid string) {
	log.Println("inserting------")
	var Count int
	query := "SELECT COUNT(*) FROM Monitor WHERE Monitor.IP ='" + ip + "'"
	row := OpenConnection["GTI"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {
		rowss, err := OpenConnection["GTI"].Exec("insert into Monitor (IP,Mailid) values (?,?)", ip, mailid)
		if err != nil {
			log.Println("Error -DB: Profile", err, rowss)
		}
	} else {
		log.Println("IP Already Exist")
	}

}
