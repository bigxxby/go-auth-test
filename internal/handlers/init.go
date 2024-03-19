package handlers

import (
	"log"
	"project/internal/database"
)

type MainStruct struct {
	Database *database.Database
}

func InitMainStruct() (*MainStruct, error) {
	db, err := database.InitDb()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	mainStruct := MainStruct{
		Database: &database.Database{DB: db},
	}

	return &mainStruct, nil
}
