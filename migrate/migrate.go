package main

import (
	"github.com/itismrx/go-crud/initalizers"
	"github.com/itismrx/go-crud/models"
)

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}
