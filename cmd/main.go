package main

import (
	controller "serviceX/src/api"
	"serviceX/src/config"
	"serviceX/src/handler/broker/nats"
	"serviceX/src/handler/database/nosql"
	"serviceX/src/handler/log"
	"serviceX/src/handler/validator"
	"serviceX/src/model/other"
)

// @title           Test app API
// @version         1.0
// @description     API of test app.
// @BasePath  /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	err := config.CreateInstance()
	if err != nil {
		panic(err)
	}
	//Initializate log
	log.CreateInstance(log.ErrorLevel)

	//Initializate validation
	validator.New()
	err = validator.Get().Validate(config.Get())
	if err != nil {
		panic(err)
	}

	//Initializate broker
	err = nats.CreateInstance(config.Get().Name)
	if err != nil {
		panic(err)
	}
	nats.Get().SetMessageEvent(func(msg []byte) error {
		return nil
	})
	log.Get().SetLocal(true) //only print in terminal remove after debug
	log.Get().SetChannels(nats.Get().GetChannel())

	/* 	//Config database SQL
	   	stuff := new(stuff.Stuff)
	   	db := sql.CreatePostgres(stuff)
	   	defer func() {
	   		db.Close()
	   	}()

	   	//SQL controller
	   	controller.New(db, stuff) */

	//Config database NoSQL

/* 	dbMongo := nosql.CreateMongo(config.Get().Name)
	defer func() {
		dbMongo.Close()
	}()
	//NoSQL controler
	controller.New(dbMongo, &other.Other{}) */

}
