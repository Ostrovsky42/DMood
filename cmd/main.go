package main

import (
	"DMood/config"
	"DMood/infrastructure/bot_handle"
	"DMood/infrastructure/storage"
	//logInit "DMood/log"
	"database/sql"
	"fmt"
	"time"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	//l:=logInit.Init()

	config := config.NewConfig()

	db := ConnectDB(config.DBConfig)
	defer db.Close()

	storage := storage.NewStorage(*db)

	dsp := bot_handle.NewDispatcher(config.BotConfig, storage)
	log.Info(dsp.Poll())
	for {
		log.Info(time.Now().Clock())
		time.Sleep(time.Minute)
	}
}

func ConnectDB(conf config.DBConfig) (db *sql.DB) {
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
