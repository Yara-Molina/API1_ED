package repositories

import "api/src/domain"

type LoanRepository interface {
	CreateLoan(loan *domain.Loan) error
}
