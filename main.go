package main

import (
	usecase "api/src/application/useCase"
	"api/src/infraestructure"
	"api/src/infraestructure/controllers"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	eventService, err := infraestructure.Setup()
	if err != nil {
		log.Fatal("Failed to set up infrastructure: ", err)
	}

	loanUseCase := &usecase.LoanUseCase{EventService: *eventService}

	loanController := controllers.NewLoanController(loanUseCase)

	router := gin.Default()

	// Configuración de CORS usando gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},         // Origen permitido
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},        // Métodos permitidos
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Encabezados permitidos
		AllowCredentials: true,
	}))

	// Definir las rutas
	router.POST("/loans", loanController.CreateLoan)

	log.Println("Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
