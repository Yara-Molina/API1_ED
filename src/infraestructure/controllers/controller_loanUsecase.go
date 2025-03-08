package controllers

import (
	usecase "api/src/application/useCase"
	"api/src/domain"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var idCounter int32
var mu sync.Mutex

type LoanController struct {
	LoanUseCase *usecase.LoanUseCase
}

func NewLoanController(loanUseCase *usecase.LoanUseCase) *LoanController {
	return &LoanController{LoanUseCase: loanUseCase}
}

func (lc *LoanController) CreateLoan(c *gin.Context) {
	var loan domain.Loan

	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Usar el contador para asignar un ID incremental
	mu.Lock()
	loan.ID = idCounter
	idCounter++ // Incrementar el contador después de asignar el ID
	mu.Unlock()

	// Generar fechas de préstamo y vencimiento
	loan.LoanDate = time.Now()
	loan.DueDate = time.Now().Add(30 * 24 * time.Hour)
	loan.Status = "Pending"

	// Llamar al caso de uso para crear el préstamo
	err := lc.LoanUseCase.CreateLoan(&loan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el préstamo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Préstamo creado y evento publicado", "loan": loan})
}
