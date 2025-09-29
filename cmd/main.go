package main

import (
	"go-backend-api/cmd/api"
	"go-backend-api/config"
	"go-backend-api/db"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := db.NewMySQLStorage(mysql.Config{
		User: 					config.Envs.DBUser,
		Passwd: 				config.Envs.DBPassword,
		Addr: 					config.Envs.DBAddress,
		DBName: 				config.Envs.DBName,
		Net: 					"tcp",
		AllowNativePasswords: 	true,
		ParseTime: 				true,
	})
	initStorage(db)

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	} 
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("DB: Successfully connected!")
}