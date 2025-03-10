package db

import (
	"api/src/domain"
	"database/sql"
)

type MySQLRepository struct {
	DB *sql.DB
}

func (repo *MySQLRepository) CreateLoan(loan *domain.Loan) error {
	query := "INSERT INTO loans (title, borrower, loan_date, due_date, Status) VALUES (?, ?, ?, ?, ?)"
	_, err := repo.DB.Exec(query, loan.Title, loan.Borrower, loan.LoanDate, loan.DueDate, loan.Status)
	return err
}
