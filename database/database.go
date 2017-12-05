package database

import "database/sql"

//DbClient CLient of the database
var DbClient *sql.DB

//InitDB is here to declare one time the database
func InitDB() {
	DbClient, err := sql.Open("mysql", "root:root@/memes")
	defer DbClient.Close()
	if err != nil {
		panic(err.Error())
	}
}
