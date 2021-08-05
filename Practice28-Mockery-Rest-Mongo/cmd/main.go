package main

import (
	"awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/model"
	db "awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/repository"
	. "awesomeProject/Practice28-Mockery-Rest-Mongo/router"
)

func main() {
	dbHost := "localhost:27017"
	db.Init(&model.Database{
		Driver:   "mongodb",
		Endpoint: dbHost})
	defer db.Exit()

	Router()

}
