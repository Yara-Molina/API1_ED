package UseCase

import (
	"api/src/application/services"
	"api/src/domain"
)

type LoanUseCase struct {
	EventService services.EventService
}

func (uc *LoanUseCase) CreateLoan(loan *domain.Loan) error {
	return uc.EventService.PublishLoanEvent(loan)
}
