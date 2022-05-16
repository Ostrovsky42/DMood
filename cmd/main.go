package main

import (
	"DMood/config"
	"DMood/infrastructure/bot_handle"
	"DMood/infrastructure/storage"
	logInit "DMood/log"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)


func main() {

	logInit.Init()
	config:=config.NewConfig()
	

	db := ConnectDB(config.DBConfig)
	defer db.Close()

	storage:= storage.NewStorage(*db)
	fn:=bot_handle.New
	dsp:=bot_handle.NewDispatcher(config.TgToken,storage,fn)
	log.Info(dsp.Poll())
for{
	log.Info(time.Now().Clock())
	time.Sleep(time.Second*10)
}
}


func ConnectDB(conf config.DBConfig) (db *sql.DB){
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
//	runMigrate(conf)
	log.Info("db connected")
	return db
}


func runMigrate(conf config.DBConfig){
	// psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)
	m, err := migrate.New(
		"file://DMood/migrate",psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

}
