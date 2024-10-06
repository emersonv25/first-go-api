package main

import (
	"first-go-api/inits"
	"first-go-api/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
}
