package repository

import (
	"context"
	"database/sql"
	"golang-restfullapi/model/domain"
)

type CategoryRepository interface {
	//Untuk parameternya biasakan kalau menggunakan golang itu selalu diawali dengan konteks 
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category //dikasih return
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category//(domain.Category)-> mengembalikan data category nya
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}