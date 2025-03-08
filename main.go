package main

import (
	usecase "api/src/application/useCase"
	"api/src/infraestructure"
	"api/src/infraestructure/controllers"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar infraestructura (conexión a RabbitMQ, base de datos, etc.)
	eventService, err := infraestructure.Setup()
	if err != nil {
		log.Fatal("Failed to set up infrastructure: ", err)
	}

	// Crear instancia del caso de uso
	loanUseCase := &usecase.LoanUseCase{EventService: *eventService}

	// Crear controlador
	loanController := controllers.NewLoanController(loanUseCase)

	// Configurar Gin
	router := gin.Default()

	// Definir rutas
	router.POST("/loans", loanController.CreateLoan)

	// Iniciar servidor en el puerto 8080
	log.Println("Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
