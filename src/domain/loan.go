package domain

import "time"

type Loan struct {
	ID         int32
	Title      string
	Borrower   string
	LoanDate   time.Time
	DueDate    time.Time
	ReturnDate *time.Time
	Status     string
}
