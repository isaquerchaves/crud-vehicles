package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaquerchaves/GO-CRUD/config"
	"github.com/isaquerchaves/GO-CRUD/controllers"
	"github.com/isaquerchaves/GO-CRUD/models"
)

func init() {
	config.ConnectToDb()
}

func main() {
	r := gin.Default()

	config.DB.AutoMigrate(&models.Car{})

	r.POST("/carros", controllers.CreateCar)
	r.PUT("/carros/:id", controllers.UpdateCar)
	r.GET("/carros/", controllers.GetAllCars)
	r.GET("/carros/:id", controllers.GetCar)
	r.DELETE("/carros/:id", controllers.DeleteCar)
	r.Run()
}
