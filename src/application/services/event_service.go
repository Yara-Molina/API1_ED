package services

import (
	"api/src/application/repositories"
	"api/src/domain"
)

type EventService struct {
	LoanRepo repositories.LoanRepository
}

func (es *EventService) PublishLoanEvent(loan *domain.Loan) error {
	return es.LoanRepo.CreateLoan(loan)
}
