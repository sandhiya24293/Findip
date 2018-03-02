package Mysql

import (
	InitDb "Findip/Common/DB/Mysql/InitializeDB"
	"database/sql"
)

var OpenConnection = make(map[string]*sql.DB)

func init() {
	OpenConnection = InitDb.Ret()
}
