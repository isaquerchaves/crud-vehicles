package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/isaquerchaves/GO-CRUD/config"
	"github.com/isaquerchaves/GO-CRUD/models"
)

func CreateCar(c *gin.Context) {
	// Get data off req body = Obter dados do corpo da solicitação
	var vehicle struct {
		Id      int
		Placa   string
		Chassi  string
		Renavam int
		Modelo  string
		Marca   string
		Ano     int
	}

	c.Bind(&vehicle)

	// Criar carro
	car := models.Car{Placa: vehicle.Placa, Chassi: vehicle.Chassi, Renavam: vehicle.Renavam, Modelo: vehicle.Modelo, Marca: vehicle.Marca, Ano: vehicle.Ano}

	result := config.DB.Create(&car)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Retornar a criação do carro

	c.JSON(200, gin.H{
		"": car,
	})
}

func GetCar(c *gin.Context) {
	// buscar id da url
	id := c.Param("id")

	// Buscar carro
	var car models.Car
	config.DB.First(&car, id)

	// retorna o carro
	c.JSON(200, gin.H{
		"": car,
	})
}

func GetAllCars(c *gin.Context) {
	// Buscar carro
	var cars []models.Car
	config.DB.Find(&cars)

	// Respond with them
	c.JSON(200, gin.H{
		"": cars,
	})
}

func UpdateCar(c *gin.Context) {
	// Buscar o id
	id := c.Param("id")

	// Buscar os dados da solicitação
	var vehicle struct {
		Id      int
		Placa   string
		Chassi  string
		Renavam int
		Modelo  string
		Marca   string
		Ano     int
	}

	c.Bind(&vehicle)

	// Encontrar o carro que estamos atualizando
	var car models.Car
	config.DB.First(&car, id)

	// Atualizar
	config.DB.Model(&car).Updates(models.Car{
		Placa:   vehicle.Placa,
		Chassi:  vehicle.Chassi,
		Renavam: vehicle.Renavam,
		Modelo:  vehicle.Modelo,
		Marca:   vehicle.Marca,
		Ano:     vehicle.Ano,
	})

	// Retornar a criação do carro
	c.JSON(200, gin.H{
		"": car,
	})
}

func DeleteCar(c *gin.Context) {
	// Buscar o id
	id := c.Param("id")

	// Deletar o carro
	config.DB.Delete(&models.Car{}, id)

	// Resposta sucesso
	c.Status(200)
}
